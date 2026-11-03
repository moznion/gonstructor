package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func batchFixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	t.Setenv("GOWORK", "off")
	t.Setenv("GOPROXY", "off")
	writeFixture(t, filepath.Join(root, "go.mod"), "module example.com/batch\n\ngo 1.26.0\n")
	writeFixture(t, filepath.Join(root, "first/service.go"), "package first\nimport alias \"bytes\"\ntype Service struct { dependency *alias.Buffer; ignored string `gonstructor:\"-\"` }\n")
	writeFixture(t, filepath.Join(root, "second/service.go"), "package second\nimport alias \"strings\"\ntype Service struct { dependency *alias.Reader }\n")
	return root
}

func writeFixture(t *testing.T, path, source string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(source), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestBatchMatchesSingleCommandsAndRestoresDirectory(t *testing.T) {
	root := batchFixture(t)
	original, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.Chdir(original) }()
	args := []string{"-type", "Service", "-constructorTypes", "allArgs,builder", "-withGetter", "-returnValue", "-output", "constructor_gen.go"}
	expected := map[string]string{}
	for _, directory := range []string{"first", "second"} {
		if err := os.Chdir(filepath.Join(root, directory)); err != nil {
			t.Fatal(err)
		}
		options, err := parseOptions(args)
		if err != nil {
			t.Fatal(err)
		}
		if err := generateSingle(options); err != nil {
			t.Fatal(err)
		}
		data, err := os.ReadFile("constructor_gen.go")
		if err != nil {
			t.Fatal(err)
		}
		expected[directory] = string(data)
		if err := os.Remove("constructor_gen.go"); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.Chdir(original); err != nil {
		t.Fatal(err)
	}
	manifest := batchManifest{Jobs: []batchJob{{Directory: "first", Args: args}, {Directory: "second", Args: args}}}
	data, err := json.Marshal(manifest)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(root, "constructors.json")
	writeFixture(t, path, string(data))
	if err := runBatch(path); err != nil {
		t.Fatal(err)
	}
	gotDirectory, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if gotDirectory != original {
		t.Fatalf("batch changed directory: %q", gotDirectory)
	}
	for directory, want := range expected {
		got, err := os.ReadFile(filepath.Join(root, directory, "constructor_gen.go"))
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != want {
			t.Fatalf("%s: batch differs from single command\ngot:\n%s\nwant:\n%s", directory, got, want)
		}
	}
}

func TestBatchValidatesBeforeReplacingOutputs(t *testing.T) {
	root := batchFixture(t)
	path := filepath.Join(root, "first/constructor_gen.go")
	original := "// Keep this file on failure.\npackage first\n"
	writeFixture(t, path, original)
	manifest := batchManifest{Jobs: []batchJob{
		{Directory: "first", Args: []string{"-type", "Service", "-output", "constructor_gen.go"}},
		{Directory: "second", Args: []string{"-type", "Missing", "-output", "constructor_gen.go"}},
	}}
	data, err := json.Marshal(manifest)
	if err != nil {
		t.Fatal(err)
	}
	input := filepath.Join(root, "constructors.json")
	writeFixture(t, input, string(data))
	if err := runBatch(input); err == nil {
		t.Fatal("missing type accepted")
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != original {
		t.Fatal("failed batch replaced an existing output")
	}
}

func TestBatchRejectsInvalidManifests(t *testing.T) {
	root := batchFixture(t)
	tests := map[string]string{
		"empty":             `{"jobs":[]}`,
		"unknown field":     `{"jobs":[],"unknown":true}`,
		"trailing JSON":     `{"jobs":[]} {}`,
		"missing directory": `{"jobs":[{"directory":"missing","args":["-type","Service"]}]}`,
		"nested batch":      `{"jobs":[{"directory":"first","args":["-batch","other.json"]}]}`,
		"version":           `{"jobs":[{"directory":"first","args":["-version"]}]}`,
		"multiple packages": `{"jobs":[{"directory":"first","args":["-type","Service",".","../second"]}]}`,
		"duplicate output":  `{"jobs":[{"directory":"first","args":["-type","Service","-output","same.go"]},{"directory":"first","args":["-type","Service","-output","same.go"]}]}`,
	}
	for name, source := range tests {
		t.Run(name, func(t *testing.T) {
			path := filepath.Join(root, "constructors.json")
			writeFixture(t, path, source)
			if err := runBatch(path); err == nil {
				t.Fatal("invalid manifest accepted")
			}
		})
	}
}

func TestBatchSupportsMultipleOutputsFromOnePackage(t *testing.T) {
	root := batchFixture(t)
	writeFixture(t, filepath.Join(root, "first/other.go"), "package first\ntype Other struct { count int }\n")
	source := `{"jobs":[{"directory":"first","args":["-type","Service","-output","service_gen.go"]},{"directory":"first","args":["-type","Other","-output","other_gen.go"]}]}`
	path := filepath.Join(root, "constructors.json")
	writeFixture(t, path, source)
	if err := runBatch(path); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"service_gen.go", "other_gen.go"} {
		if _, err := os.Stat(filepath.Join(root, "first", name)); err != nil {
			t.Fatal(err)
		}
	}
}

func TestBatchFlagRejectsSingleCommandOptions(t *testing.T) {
	_, err := parseOptions([]string{"-batch", "constructors.json", "-type", "Service"})
	if err == nil {
		t.Fatal("mixed batch and single-command flags accepted")
	}
}

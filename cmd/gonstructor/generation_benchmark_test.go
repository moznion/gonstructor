package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// BenchmarkGeneration includes package discovery, parsing, formatting, and real
// output writes. Each iteration regenerates every output; no result is cached.
func BenchmarkGeneration(b *testing.B) {
	for _, count := range []int{1, 10, 50} {
		b.Run(fmt.Sprintf("packages=%d", count), func(b *testing.B) {
			root := b.TempDir()
			b.Setenv("GOWORK", "off")
			b.Setenv("GOPROXY", "off")
			b.Chdir(root)
			writeFixture(b, filepath.Join(root, "go.mod"), "module example.com/benchmark\n\ngo 1.26.0\n")
			args := []string{"-type", "Service", "-constructorTypes", "allArgs,builder", "-withGetter", "-output", "constructor_gen.go"}
			options, err := parseOptions(args)
			if err != nil {
				b.Fatal(err)
			}
			manifest := batchManifest{}
			for i := 0; i < count; i++ {
				directory := fmt.Sprintf("service%d", i)
				writeFixture(b, filepath.Join(root, directory, "service.go"), fmt.Sprintf("package service%d\nimport alias \"bytes\"\ntype Service struct { dependency *alias.Buffer; name string; ignored int `gonstructor:\"-\"` }\n", i))
				manifest.Jobs = append(manifest.Jobs, batchJob{Directory: directory, Args: args})
			}
			data, err := json.Marshal(manifest)
			if err != nil {
				b.Fatal(err)
			}
			path := filepath.Join(root, "constructors.json")
			writeFixture(b, path, string(data))
			single := func(b *testing.B) {
				for _, job := range manifest.Jobs {
					if err := os.Chdir(filepath.Join(root, job.Directory)); err != nil {
						b.Fatal(err)
					}
					if err := generateSingle(options); err != nil {
						b.Fatal(err)
					}
				}
				if err := os.Chdir(root); err != nil {
					b.Fatal(err)
				}
			}
			batch := func(b *testing.B) {
				if err := runBatch(path); err != nil {
					b.Fatal(err)
				}
			}

			// Warm both paths and check output equivalence outside the timed sections.
			single(b)
			expected := make([]string, count)
			for i, job := range manifest.Jobs {
				output := filepath.Join(root, job.Directory, "constructor_gen.go")
				content, err := os.ReadFile(output)
				if err != nil {
					b.Fatal(err)
				}
				expected[i] = string(content)
				if err := os.Remove(output); err != nil {
					b.Fatal(err)
				}
			}
			batch(b)
			for i, job := range manifest.Jobs {
				content, err := os.ReadFile(filepath.Join(root, job.Directory, "constructor_gen.go"))
				if err != nil {
					b.Fatal(err)
				}
				if string(content) != expected[i] {
					b.Fatalf("batch output differs for %s", job.Directory)
				}
			}
			for _, method := range []struct {
				name string
				run  func(*testing.B)
			}{
				{"SingleCommands", single}, {"Batch", batch},
			} {
				b.Run(method.name, func(b *testing.B) {
					b.ReportAllocs()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						method.run(b)
					}
					b.ReportMetric(float64(count), "outputs/op")
				})
			}
		})
	}
}

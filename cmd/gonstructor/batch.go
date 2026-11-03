package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"io"
	"os"
	"path/filepath"

	"github.com/moznion/gonstructor/internal"
	"golang.org/x/tools/go/packages"
)

type batchManifest struct {
	Jobs []batchJob `json:"jobs"`
}
type batchJob struct {
	Directory string   `json:"directory"`
	Args      []string `json:"args"`
}
type plannedJob struct {
	directory, pattern, output string
	options                    options
}
type parsedPackage struct {
	metadata *packages.Package
	files    []*ast.File
}
type generatedOutput struct{ path, code string }

func runBatch(path string) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open batch: %w", err)
	}
	defer func() { _ = file.Close() }()
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	var manifest batchManifest
	if err := decoder.Decode(&manifest); err != nil {
		return fmt.Errorf("decode batch: %w", err)
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return fmt.Errorf("batch must contain one JSON document")
	}
	if len(manifest.Jobs) == 0 {
		return fmt.Errorf("batch contains no jobs")
	}
	base := filepath.Dir(path)
	jobs, err := planBatch(base, manifest.Jobs)
	if err != nil {
		return err
	}
	patterns := make([]string, 0, len(jobs))
	seen := map[string]bool{}
	for _, job := range jobs {
		if !seen[job.pattern] {
			patterns = append(patterns, job.pattern)
			seen[job.pattern] = true
		}
	}
	metadata, err := internal.ParsePackages(base, patterns)
	if err != nil {
		return fmt.Errorf("discover batch packages: %w", err)
	}
	parsed := map[string]parsedPackage{}
	for _, pkg := range metadata {
		if len(pkg.GoFiles) == 0 {
			return fmt.Errorf("package %s has no Go files: %v", pkg.ID, pkg.Errors)
		}
		directory, err := filepath.EvalSymlinks(filepath.Dir(pkg.GoFiles[0]))
		if err != nil {
			return err
		}
		files, err := internal.ParseFiles(pkg.GoFiles)
		if err != nil {
			return err
		}
		parsed[directory] = parsedPackage{metadata: pkg, files: files}
	}
	original, err := os.Getwd()
	if err != nil {
		return err
	}
	defer func() { _ = os.Chdir(original) }()
	outputs := make([]generatedOutput, 0, len(jobs))
	for index, job := range jobs {
		pkg, ok := parsed[job.pattern]
		if !ok {
			return fmt.Errorf("job %d: no package found for %s", index+1, job.pattern)
		}
		// goimports resolves aliases using the owning package's source files.
		if err := os.Chdir(job.directory); err != nil {
			return err
		}
		code, err := render(job.options, pkg.metadata, pkg.files)
		if err != nil {
			return fmt.Errorf("job %d: %w", index+1, err)
		}
		outputs = append(outputs, generatedOutput{path: job.output, code: code})
	}
	for _, output := range outputs {
		if err := writeOutput(output.path, output.code); err != nil {
			return err
		}
	}
	return nil
}

func planBatch(base string, jobs []batchJob) ([]plannedJob, error) {
	result := make([]plannedJob, 0, len(jobs))
	outputs := map[string]bool{}
	for index, job := range jobs {
		directory := job.Directory
		if !filepath.IsAbs(directory) {
			directory = filepath.Join(base, directory)
		}
		directory, err := filepath.EvalSymlinks(directory)
		if err != nil {
			return nil, fmt.Errorf("job %d directory: %w", index+1, err)
		}
		options, err := parseOptions(job.Args)
		if err != nil {
			return nil, fmt.Errorf("job %d: %w", index+1, err)
		}
		if options.batch != "" || options.showVersion {
			return nil, fmt.Errorf("job %d: nested batch and version commands are not generation jobs", index+1)
		}
		if len(options.patterns) != 1 {
			return nil, fmt.Errorf("job %d: batch jobs require one package directory", index+1)
		}
		pattern := options.patterns[0]
		if !filepath.IsAbs(pattern) {
			pattern = filepath.Join(directory, pattern)
		}
		pattern, err = filepath.EvalSymlinks(pattern)
		if err != nil {
			return nil, fmt.Errorf("job %d package directory: %w", index+1, err)
		}
		info, err := os.Stat(pattern)
		if err != nil {
			return nil, err
		}
		if !info.IsDir() {
			return nil, fmt.Errorf("job %d: batch jobs require a package directory", index+1)
		}
		output := options.output
		if output == "" {
			output = getFilenameToGenerate(options.typeNames[0], "", []string{pattern})
		}
		if !filepath.IsAbs(output) {
			output = filepath.Join(directory, output)
		}
		output = filepath.Clean(output)
		if outputs[output] {
			return nil, fmt.Errorf("job %d: duplicate output %s", index+1, output)
		}
		outputs[output] = true
		result = append(result, plannedJob{directory: directory, pattern: pattern, output: output, options: options})
	}
	return result, nil
}

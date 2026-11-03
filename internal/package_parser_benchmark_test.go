package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/packages"
)

// BenchmarkPackageDiscovery compares the current discovery requirements with
// the former mode, which also loaded syntax and type information.
func BenchmarkPackageDiscovery(b *testing.B) {
	root := b.TempDir()
	b.Setenv("GOWORK", "off")
	b.Setenv("GOPROXY", "off")
	if err := os.WriteFile(filepath.Join(root, "go.mod"), []byte("module example.com/discovery\n\ngo 1.26.0\n"), 0644); err != nil {
		b.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "service.go"), []byte("package discovery\nimport (\"net/http\"; \"database/sql\"; \"encoding/json\")\ntype Service struct { client *http.Client; db *sql.DB; decoder *json.Decoder }\n"), 0644); err != nil {
		b.Fatal(err)
	}
	for _, mode := range []struct {
		name string
		load func() ([]*packages.Package, error)
	}{
		{"MetadataOnly", func() ([]*packages.Package, error) { return ParsePackages(root, []string{"."}) }},
		{"LegacyTypes", func() ([]*packages.Package, error) {
			return packages.Load(&packages.Config{Dir: root, Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedImports | packages.NeedTypes | packages.NeedTypesSizes | packages.NeedSyntax | packages.NeedTypesInfo}, ".")
		}},
	} {
		b.Run(mode.name, func(b *testing.B) {
			load := func() {
				pkgs, err := mode.load()
				if err != nil {
					b.Fatal(err)
				}
				if len(pkgs) != 1 {
					b.Fatalf("expected one package, got %d", len(pkgs))
				}
				if len(pkgs[0].Errors) != 0 {
					b.Fatal(fmt.Sprint(pkgs[0].Errors))
				}
				if pkgs[0].Name != "discovery" || len(pkgs[0].GoFiles) != 1 {
					b.Fatal("missing package metadata")
				}
			}
			load() // Warm Go's normal build cache before timing either mode.
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				load()
			}
			b.ReportMetric(1, "packages/op")
		})
	}
}

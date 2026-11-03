package internal

import "testing"

func TestParsePackageDoesNotRequestUnusedTypeInformation(t *testing.T) {
	pkg, err := ParsePackage([]string{"."})
	if err != nil {
		t.Fatal(err)
	}
	if pkg.Name != "internal" || len(pkg.GoFiles) == 0 {
		t.Fatalf("missing package metadata: %#v", pkg)
	}
	if pkg.Types != nil || pkg.TypesInfo != nil || pkg.TypesSizes != nil || len(pkg.Syntax) != 0 {
		t.Fatal("constructor discovery loaded unused type information or syntax")
	}
}

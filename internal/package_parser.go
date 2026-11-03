package internal

import (
	"fmt"

	"golang.org/x/tools/go/packages"
)

// ParsePackage parses files according to given directory pattern to get the package information.
func ParsePackage(dirPattern []string) (*packages.Package, error) {
	pkgs, err := ParsePackages("", dirPattern)
	if err != nil {
		return nil, fmt.Errorf("failed to load package: %w", err)
	}
	if len(pkgs) != 1 {
		return nil, fmt.Errorf("ambiguous error; %d packages found", len(pkgs))
	}
	return pkgs[0], nil
}

// ParsePackages discovers package names and source files without loading types.
func ParsePackages(directory string, patterns []string) ([]*packages.Package, error) {
	return packages.Load(&packages.Config{Dir: directory, Mode: packages.NeedName | packages.NeedFiles, Tests: false}, patterns...)
}

package internal

import (
	"fmt"

	"golang.org/x/tools/go/packages"
)

// ParsePackage parses files according to given directory pattern to get the package information.
func ParsePackage(dirPattern []string) (*packages.Package, error) {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedTypes |
			packages.NeedTypesSizes |
			packages.NeedSyntax |
			packages.NeedTypesInfo,
		Tests: false,
	}, dirPattern...)
	if err != nil {
		return nil, fmt.Errorf("failed to load package: %w", err)
	}
	if len(pkgs) != 1 {
		return nil, fmt.Errorf("ambiguous error; %d packages found", len(pkgs))
	}
	return pkgs[0], nil
}

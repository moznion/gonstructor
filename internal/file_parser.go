package internal

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// ParseFiles parses the files to get AST.
func ParseFiles(files []string) ([]*ast.File, error) {
	fset := token.NewFileSet()

	astFiles := make([]*ast.File, len(files))
	for i, file := range files {
		parsed, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
		if err != nil {
			return nil, fmt.Errorf("failed to parse file: %w", err)
		}
		astFiles[i] = parsed
	}
	return astFiles, nil
}

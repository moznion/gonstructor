package internal

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// FileParser is a parser of go files to get AST.
type FileParser struct {
	astCache map[string]*ast.File
}

// NewFileParser is a constructor of FileParser.
func NewFileParser() *FileParser {
	return &FileParser{
		astCache: make(map[string]*ast.File),
	}
}

// Parse parses a file to get AST.
func (p *FileParser) Parse(files []string) ([]*ast.File, error) {
	fset := token.NewFileSet()

	astFiles := make([]*ast.File, len(files))
	for i, file := range files {
		if parsed := p.astCache[file]; parsed != nil {
			astFiles[i] = parsed
			continue
		}

		parsed, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
		if err != nil {
			return nil, fmt.Errorf("failed to parse file: %w", err)
		}
		astFiles[i] = parsed
		p.astCache[file] = parsed
	}
	return astFiles, nil
}

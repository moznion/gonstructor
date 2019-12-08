package internal

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type FileParser struct {
	astCache map[string]*ast.File
}

func NewFileParser() *FileParser {
	return &FileParser{
		astCache: make(map[string]*ast.File),
	}
}

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

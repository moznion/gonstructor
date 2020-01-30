package constructor

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	g "github.com/moznion/gowrtr/generator"
)

// AllArgsConstructorGenerator is a struct type that has the responsibility to generate a statement of a constructor with all of arguments.
type AllArgsConstructorGenerator struct {
	TypeName string
	Fields   []*Field
	InitFunc string
}

// Generate generates a constructor statement with all of arguments.
func (cg *AllArgsConstructorGenerator) Generate() g.Statement {
	funcSignature := g.NewFuncSignature(fmt.Sprintf("New%s", strcase.ToCamel(cg.TypeName)))

	retStructureKeyValues := make([]string, 0)
	for _, field := range cg.Fields {
		if field.ShouldIgnore {
			continue
		}
		funcSignature = funcSignature.AddParameters(g.NewFuncParameter(strcase.ToLowerCamel(field.FieldName), field.FieldType))
		retStructureKeyValues = append(retStructureKeyValues, fmt.Sprintf("%s: %s", field.FieldName, strcase.ToLowerCamel(field.FieldName)))
	}

	funcSignature = funcSignature.AddReturnTypes("*" + cg.TypeName)

	retStructure := fmt.Sprintf("&%s{%s}", cg.TypeName, strings.Join(retStructureKeyValues, ","))

	var stmts []g.Statement
	if cg.InitFunc != "" {
		stmts = []g.Statement{
			g.NewRawStatementf("r := %s", retStructure),
			g.NewRawStatementf("r.%s()", cg.InitFunc),
			g.NewReturnStatement("r"),
		}
	} else {
		stmts = []g.Statement{
			g.NewReturnStatement(retStructure),
		}
	}

	fn := g.NewFunc(
		nil,
		funcSignature,
		stmts...,
	)
	return fn
}

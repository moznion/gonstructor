package constructor

import (
	"fmt"

	"github.com/iancoleman/strcase"
	g "github.com/moznion/gowrtr/generator"
)

// AllArgsConstructorGenerator is a struct type that has the responsibility to
// generate a statement of a constructor with all of arguments.
type AllArgsConstructorGenerator struct {
	TypeName string
	Fields   []*Field
	InitFunc string
}

// Generate generates a constructor statement with all of arguments.
func (cg *AllArgsConstructorGenerator) Generate(indentLevel int) g.Statement {
	funcSignature := g.NewFuncSignature(fmt.Sprintf("New%s", strcase.ToCamel(cg.TypeName)))

	retStructureKeyValues := make([]string, 0)
	for _, field := range cg.Fields {
		if field.ShouldIgnore {
			continue
		}
		funcSignature = funcSignature.AddParameters(
			g.NewFuncParameter(toLowerCamel(field.FieldName), field.FieldType),
		)
		retStructureKeyValues = append(
			retStructureKeyValues,
			fmt.Sprintf("%s: %s", field.FieldName, toLowerCamel(field.FieldName)),
		)
	}

	funcSignature = funcSignature.AddReturnTypes("*" + cg.TypeName)

	retStructure := generateStructure(cg.TypeName, retStructureKeyValues, indentLevel+1)

	var stmts []g.Statement
	if cg.InitFunc != "" {
		stmts = []g.Statement{
			g.NewRawStatementf("r := %s", retStructure),
			g.NewNewline(),
			g.NewRawStatementf("r.%s()", cg.InitFunc),
			g.NewNewline(),
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

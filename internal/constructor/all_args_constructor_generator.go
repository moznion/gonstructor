package constructor

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	g "github.com/moznion/gowrtr/generator"
)

// AllArgsConstructorGenerator is a struct type that has the responsibility to
// generate a statement of a constructor with all of arguments.
type AllArgsConstructorGenerator struct {
	TypeName                 string
	Fields                   []*Field
	InitFunc                 string
	InitFuncReturnTypes      []string
	PropagateInitFuncReturns bool
	ReturnValue              bool
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

	funcSignature = funcSignature.
		AddReturnTypes(withPrefix(cg.TypeName, "*", !cg.ReturnValue)).
		AddReturnTypes(func() []string {
			if len(cg.InitFuncReturnTypes) <= 0 {
				return []string{}
			}
			return cg.InitFuncReturnTypes
		}()...)

	retStructure := generateStructure(cg.TypeName, retStructureKeyValues, indentLevel+1, cg.ReturnValue)

	var stmts []g.Statement

	if cg.InitFunc != "" {
		stmts = []g.Statement{
			g.NewRawStatementf("r := %s", retStructure),
			g.NewNewline(),
		}

		if cg.PropagateInitFuncReturns && len(cg.InitFuncReturnTypes) > 0 {
			initFuncReturnValues := make([]string, len(cg.InitFuncReturnTypes))
			for i := 0; i < len(cg.InitFuncReturnTypes); i++ {
				initFuncReturnValues[i] = fmt.Sprintf("ret_%s%d", cg.InitFunc, i)
			}

			stmts = append(stmts,
				g.NewRawStatementf(
					"%s := r.%s()",
					strings.Join(initFuncReturnValues, ", "),
					cg.InitFunc,
				),
				g.NewNewline(),
				g.NewReturnStatement("r").AddReturnItems(initFuncReturnValues...),
			)
		} else {
			stmts = append(stmts,
				g.NewRawStatementf("r.%s()", cg.InitFunc),
				g.NewNewline(),
				g.NewReturnStatement("r"),
			)
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

package constructor

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	g "github.com/moznion/gowrtr/generator"
)

type AllArgsConstructorGenerator struct {
	TypeName string
	Fields   []*Field
}

func (cg *AllArgsConstructorGenerator) Generate() g.Statement {
	funcSignature := g.NewFuncSignature(fmt.Sprintf("New%s", strcase.ToCamel(cg.TypeName)))

	retStructureKeyValues := make([]string, 0)
	for _, field := range cg.Fields {
		if field.ShouldIgnore {
			continue
		}
		funcSignature = funcSignature.AddFuncParameters(g.NewFuncParameter(strcase.ToLowerCamel(field.FieldName), field.FieldType))
		retStructureKeyValues = append(retStructureKeyValues, fmt.Sprintf("%s: %s", field.FieldName, strcase.ToLowerCamel(field.FieldName)))
	}

	funcSignature = funcSignature.AddReturnTypes("*" + cg.TypeName)

	return g.NewFunc(
		nil,
		funcSignature,
		g.NewReturnStatement(fmt.Sprintf("&%s{%s}", cg.TypeName, strings.Join(retStructureKeyValues, ","))),
	)
}

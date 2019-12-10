package constructor

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	g "github.com/moznion/gowrtr/generator"
)

// BuilderGenerator is a struct type that has the responsibility to generate a statement of a builder.
type BuilderGenerator struct {
	TypeName string
	Fields   []*Field
}

// Generate generates a builder statement.
func (cg *BuilderGenerator) Generate() g.Statement {
	builderConstructorName := fmt.Sprintf("New%sBuilder", strcase.ToCamel(cg.TypeName))
	builderType := fmt.Sprintf("%sBuilder", strcase.ToCamel(cg.TypeName))

	builderConstructorFunc :=
		g.NewFunc(
			nil,
			g.NewFuncSignature(builderConstructorName).AddReturnTypes(fmt.Sprintf("*%s", builderType)),
			g.NewReturnStatement(fmt.Sprintf("&%s{}", builderType)),
		)

	builderStruct := g.NewStruct(builderType)
	fieldRegistererFunctions := make([]*g.Func, 0)
	retStructureKeyValues := make([]string, 0)
	for _, field := range cg.Fields {
		if field.ShouldIgnore {
			continue
		}

		builderStruct = builderStruct.AddField(
			toLowerCamel(field.FieldName),
			field.FieldType,
		)

		fieldRegistererFunctions = append(fieldRegistererFunctions, g.NewFunc(
			g.NewFuncReceiver("b", "*"+builderType),
			g.NewFuncSignature(strcase.ToCamel(field.FieldName)).
				AddFuncParameters(g.NewFuncParameter(toLowerCamel(field.FieldName), field.FieldType)).
				AddReturnTypes("*"+builderType),
			g.NewRawStatement(fmt.Sprintf("b.%s = %s", toLowerCamel(field.FieldName), strcase.ToLowerCamel(field.FieldName))),
			g.NewReturnStatement("b"),
		))

		retStructureKeyValues = append(retStructureKeyValues, fmt.Sprintf("%s: b.%s", field.FieldName, toLowerCamel(field.FieldName)))
	}

	buildFunc := g.NewFunc(
		g.NewFuncReceiver("b", "*"+builderType),
		g.NewFuncSignature("Build").
			AddReturnTypes("*"+cg.TypeName),
		g.NewReturnStatement(fmt.Sprintf("&%s{%s}", cg.TypeName, strings.Join(retStructureKeyValues, ","))),
	)

	stmt := g.NewRoot(builderStruct, builderConstructorFunc)
	for _, f := range fieldRegistererFunctions {
		stmt = stmt.AddStatements(f)
	}
	return stmt.AddStatements(buildFunc)
}

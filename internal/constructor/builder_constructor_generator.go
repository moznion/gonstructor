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
	InitFunc string
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
			strcase.ToLowerCamel(field.FieldName),
			field.FieldType,
		)

		fieldRegistererFunctions = append(fieldRegistererFunctions, g.NewFunc(
			g.NewFuncReceiver("b", "*"+builderType),
			g.NewFuncSignature(strcase.ToCamel(field.FieldName)).
				AddParameters(g.NewFuncParameter(strcase.ToLowerCamel(field.FieldName), field.FieldType)).
				AddReturnTypes("*"+builderType),
			g.NewRawStatement(fmt.Sprintf("b.%s = %s", strcase.ToLowerCamel(field.FieldName), strcase.ToLowerCamel(field.FieldName))),
			g.NewReturnStatement("b"),
		))

		retStructureKeyValues = append(retStructureKeyValues, fmt.Sprintf("%s: b.%s", field.FieldName, strcase.ToLowerCamel(field.FieldName)))
	}

	buildResult := fmt.Sprintf("&%s{%s}", cg.TypeName, strings.Join(retStructureKeyValues, ","))

	var buildStmts []g.Statement
	if cg.InitFunc != "" {
		buildStmts = append(buildStmts, g.NewRawStatementf("r := %s", buildResult))
		buildStmts = append(buildStmts, g.NewRawStatementf("r.%s()", cg.InitFunc))
		buildStmts = append(buildStmts, g.NewReturnStatement("r"))
	} else {
		buildStmts = append(
			buildStmts,
			g.NewReturnStatement(buildResult),
		)
	}

	buildFunc := g.NewFunc(
		g.NewFuncReceiver("b", "*"+builderType),
		g.NewFuncSignature("Build").
			AddReturnTypes("*"+cg.TypeName),
		buildStmts...,
	)

	stmt := g.NewRoot(builderStruct, builderConstructorFunc)
	for _, f := range fieldRegistererFunctions {
		stmt = stmt.AddStatements(f)
	}
	return stmt.AddStatements(buildFunc)
}

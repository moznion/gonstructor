package constructor

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	g "github.com/moznion/gowrtr/generator"
)

// BuilderGenerator is a struct type that has the responsibility to generate a statement of a builder.
type BuilderGenerator struct {
	TypeName                 string
	Fields                   []*Field
	InitFunc                 string
	InitFuncReturnTypes      []string
	PropagateInitFuncReturns bool
	ReturnValue              bool
	SetterPrefix             string
}

// Generate generates a builder statement.
func (cg *BuilderGenerator) Generate(indentLevel int) g.Statement {
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
			g.NewFuncSignature(withConditionalPrefix(strcase.ToCamel(field.FieldName), cg.SetterPrefix, cg.SetterPrefix != "")).
				AddParameters(g.NewFuncParameter(toLowerCamel(field.FieldName), field.FieldType)).
				AddReturnTypes("*"+builderType),
			g.NewRawStatement(fmt.Sprintf("b.%s = %s", toLowerCamel(field.FieldName), toLowerCamel(field.FieldName))),
			g.NewReturnStatement("b"),
		))

		retStructureKeyValues = append(retStructureKeyValues, fmt.Sprintf("%s: b.%s", field.FieldName, toLowerCamel(field.FieldName)))
	}

	buildResult := generateStructure(cg.TypeName, retStructureKeyValues, indentLevel+1, cg.ReturnValue)

	var buildStmts []g.Statement
	if cg.InitFunc != "" {
		buildStmts = append(buildStmts, g.NewRawStatementf("r := %s", buildResult))

		if cg.PropagateInitFuncReturns && len(cg.InitFuncReturnTypes) > 0 {
			initFuncReturnValues := make([]string, len(cg.InitFuncReturnTypes))
			for i := 0; i < len(cg.InitFuncReturnTypes); i++ {
				initFuncReturnValues[i] = fmt.Sprintf("ret_%s%d", cg.InitFunc, i)
			}

			buildStmts = append(buildStmts, g.NewRawStatementf(
				"%s := r.%s()",
				strings.Join(initFuncReturnValues, ", "),
				cg.InitFunc,
			))
			buildStmts = append(buildStmts, g.NewReturnStatement("r").AddReturnItems(initFuncReturnValues...))
		} else {
			buildStmts = append(buildStmts, g.NewRawStatementf("r.%s()", cg.InitFunc))
			buildStmts = append(buildStmts, g.NewReturnStatement("r"))
		}
	} else {
		buildStmts = append(
			buildStmts,
			g.NewReturnStatement(buildResult),
		)
	}

	buildFunc := g.NewFunc(
		g.NewFuncReceiver("b", "*"+builderType),
		g.NewFuncSignature("Build").AddReturnTypes(withConditionalPrefix(cg.TypeName, "*", !cg.ReturnValue)).AddReturnTypes(func() []string {
			if len(cg.InitFuncReturnTypes) <= 0 {
				return []string{}
			}
			return cg.InitFuncReturnTypes
		}()...),
		buildStmts...,
	)

	stmt := g.NewRoot(builderStruct, builderConstructorFunc)
	for _, f := range fieldRegistererFunctions {
		stmt = stmt.AddStatements(g.NewNewline(), f)
	}
	return stmt.AddStatements(g.NewNewline(), buildFunc)
}

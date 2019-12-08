package internal

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/moznion/gonstructor/internal/constructor"
	g "github.com/moznion/gowrtr/generator"
)

const receiverName = "s"

// GenerateGetters generates getters for each field.
func GenerateGetters(typeName string, fields []*constructor.Field) g.Statement {
	stmt := g.NewRoot()
	for _, field := range fields {
		stmt = stmt.AddStatements(
			g.NewFunc(
				g.NewFuncReceiver(receiverName, "*"+typeName),
				g.NewFuncSignature(fmt.Sprintf("Get%s", strcase.ToCamel(field.FieldName))).
					AddReturnTypes(field.FieldType),
				g.NewReturnStatement(fmt.Sprintf("%s.%s", receiverName, field.FieldName)),
			),
		)
	}
	return stmt
}

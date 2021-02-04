package constructor

import (
	"fmt"
	"go/ast"
	"go/types"
	"reflect"
	"strings"
)

const gonstructorTag = "gonstructor"

// CollectConstructorFieldsFromAST collects fields to include in a constructor from the AST.
func CollectConstructorFieldsFromAST(typeName string, astFiles []*ast.File) ([]*Field, error) {
	for _, astFile := range astFiles {
		for _, decl := range astFile.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				structName := typeSpec.Name.Name
				if typeName != structName {
					continue
				}

				structType, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}

				return convertStructFieldsToConstructorOnes(structType.Fields.List), nil
			}
		}
	}

	return nil, fmt.Errorf("there is no suitable struct that matches given typeName [given=%s]", typeName)
}

func convertStructFieldsToConstructorOnes(fields []*ast.Field) []*Field {
	fs := make([]*Field, len(fields))
	for i, field := range fields {
		shouldIgnore := false
		if field.Tag != nil && len(field.Tag.Value) >= 1 {
			customTag := reflect.StructTag(field.Tag.Value[1 : len(field.Tag.Value)-1])
			shouldIgnore = customTag.Get(gonstructorTag) == "-"
		}

		fieldType := types.ExprString(field.Type)

		var fieldName string
		if len(field.Names) > 0 {
			fieldName = field.Names[0].Name
		} else {
			// split 'mypackage.MyType'
			chunks := strings.Split(fieldType, ".")

			// it could be a pointer: '*mypackage.MyType' or '*MyType'
			fieldName = strings.TrimPrefix(chunks[len(chunks)-1], "*")
		}

		fs[i] = &Field{
			FieldName:    fieldName,
			FieldType:    fieldType,
			ShouldIgnore: shouldIgnore,
		}
	}
	return fs
}

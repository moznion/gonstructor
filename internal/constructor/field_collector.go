package constructor

import (
	"fmt"
	"go/ast"
	"go/types"
	"reflect"
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

		fs[i] = &Field{
			FieldName:    field.Names[0].Name,
			FieldType:    types.ExprString(field.Type),
			ShouldIgnore: shouldIgnore,
		}
	}
	return fs
}

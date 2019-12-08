package constructor

// ConstructorField represents a field of the structure for a constructor to be generated.
type Field struct {
	// FieldName is a name of the field.
	FieldName string
	// FieldType is a type of the field.
	FieldType string
	// ShouldIgnore marks whether the field should be ignored or not in a constructor.
	ShouldIgnore bool
}

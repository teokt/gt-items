package filter

import (
	"reflect"
	"strings"

	"github.com/teokt/gt-items/internal/utils"
)

type FieldType string

const (
	FieldTypeString = "string"
	FieldTypeEnum   = "enum"
	FieldTypeFlags  = "flags"
	FieldTypeInt    = "int"
)

type Field[T any] struct {
	Name      string
	Type      FieldType
	Accessor  func(T) any
	Condition Condition
}

func createFieldMap[T any]() map[string]*Field[T] {
	fields := make(map[string]*Field[T])

	typ := reflect.TypeFor[T]()
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	for i := range typ.NumField() {
		fieldTyp := typ.Field(i)
		fieldKind := fieldTyp.Type.Kind()

		if fieldTyp.PkgPath != "" {
			continue
		}

		var field Field[T]

		switch {
		case fieldKind == reflect.String:
			field.Type = FieldTypeString

		case fieldTyp.Type.Implements(reflect.TypeFor[utils.Enum]()):
			field.Type = FieldTypeEnum

		case fieldTyp.Type.Implements(reflect.TypeFor[utils.Flags]()):
			field.Type = FieldTypeFlags

		case utils.IsInt(fieldKind):
			field.Type = FieldTypeInt

		default:
			continue
		}

		field.Name = fieldTyp.Name

		field.Accessor = func(v T) any {
			val := reflect.Indirect(reflect.ValueOf(v))
			return val.Field(i).Interface()
		}

		fieldName := strings.ToLower(field.Name)
		fields[fieldName] = &field
	}

	return fields
}

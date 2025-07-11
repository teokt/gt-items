package filter

import (
	"fmt"
	"reflect"
	"strings"
)

type Matcher[T any] struct {
	conditions map[string]Condition
	accessors  map[string]func(T) any
}

func NewMatcher[T any]() *Matcher[T] {
	return &Matcher[T]{
		conditions: make(map[string]Condition),
		accessors:  createAccessors[T](),
	}
}

func createAccessors[T any]() map[string]func(T) any {
	accessors := make(map[string]func(T) any)

	typ := reflect.TypeFor[T]()
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	for i := range typ.NumField() {
		fieldTyp := typ.Field(i)
		if fieldTyp.PkgPath != "" {
			continue
		}

		// no struct support yet
		fieldKind := fieldTyp.Type.Kind()
		if fieldKind == reflect.Struct {
			continue
		}

		fieldName := strings.ToLower(fieldTyp.Name)

		accessors[fieldName] = func(v T) any {
			val := reflect.Indirect(reflect.ValueOf(v))
			return val.Field(i).Interface()
		}
	}

	return accessors
}

func (m *Matcher[T]) ClearFilters() {
	m.conditions = make(map[string]Condition)
}

func (m *Matcher[T]) AddFilter(filter string) error {
	if !strings.Contains(filter, "=") {
		return fmt.Errorf("invalid filter format: %s", filter)
	}

	parts := strings.SplitN(strings.ToLower(filter), "=", 2)
	fieldName := strings.TrimPrefix(parts[0], "--")
	expr := parts[1]

	if m.accessors[fieldName] == nil {
		return fmt.Errorf("invalid field: %s", fieldName)
	}

	cond, err := ParseExpression(expr)
	if err != nil {
		return fmt.Errorf("parse error in %s: %w", filter, err)
	}

	// fmt.Printf("condition: %#v\n", cond)
	m.conditions[fieldName] = cond
	return nil
}

func (m *Matcher[T]) Matches(v T) bool {
	for fieldName, cond := range m.conditions {
		accessor, ok := m.accessors[fieldName]
		if !ok {
			return false
		}

		value := accessor(v)
		if value == nil {
			return false
		}

		if !cond.Match(value) {
			return false
		}
	}
	return true
}

package filter

import (
	"fmt"
	"strings"
)

type Matcher[T any] struct {
	Fields map[string]*Field[T]
}

func NewMatcher[T any]() *Matcher[T] {
	return &Matcher[T]{
		Fields: createFieldMap[T](),
	}
}

func (m *Matcher[T]) ClearFilters() {
	for _, field := range m.Fields {
		field.Condition = nil
	}
}

func (m *Matcher[T]) AddFilter(filter string) error {
	if !strings.Contains(filter, "=") {
		return fmt.Errorf("invalid filter format: %s", filter)
	}

	parts := strings.SplitN(strings.ToLower(filter), "=", 2)
	fieldName := strings.TrimPrefix(parts[0], "--")
	expr := strings.TrimRight(parts[1], " ")

	field, ok := m.Fields[fieldName]
	if !ok {
		return fmt.Errorf("invalid field: %s", fieldName)
	}

	cond, err := ParseExpression(field.Type, expr)
	if err != nil {
		return fmt.Errorf("parse error in %s: %w", filter, err)
	}

	// fmt.Printf("condition: %#v\n", cond)
	field.Condition = cond
	return nil
}

func (m *Matcher[T]) Matches(v T) bool {
	for _, field := range m.Fields {
		cond := field.Condition
		if cond == nil {
			continue
		}

		value := field.Accessor(v)
		if value == nil {
			continue
		}

		if !cond.Match(value) {
			return false
		}
	}
	return true
}

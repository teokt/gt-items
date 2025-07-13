package filter

import (
	"strings"

	"github.com/teokt/gt-items/internal/utils"
)

type Condition interface {
	Match(value any) bool
}

type Or []Condition

func (o Or) Match(v any) bool {
	for _, cond := range o {
		if cond.Match(v) {
			return true
		}
	}
	return false
}

type And []Condition

func (a And) Match(v any) bool {
	for _, cond := range a {
		if !cond.Match(v) {
			return false
		}
	}
	return true
}

type Not struct{ Condition }

func (n Not) Match(v any) bool {
	return !n.Condition.Match(v)
}

type IntRange struct {
	Min *int
	Max *int
}

func (r IntRange) Match(v any) bool {
	n := utils.ToInt(v)
	if r.Min != nil && n < *r.Min {
		return false
	}
	if r.Max != nil && n > *r.Max {
		return false
	}
	return true
}

type Int struct{ int }

func (i Int) Match(v any) bool {
	return utils.ToInt(v) == i.int
}

type String struct{ string }

func (s String) Match(v any) bool {
	str, ok := v.(string)
	if !ok {
		return false
	}
	return strings.Contains(strings.ToLower(str), s.string)
}

type Enum struct{ string }

func (e Enum) Match(v any) bool {
	enum, ok := v.(utils.Enum)
	if !ok {
		return false
	}
	return strings.ToLower(enum.String()) == e.string
}

type Flags struct{ string }

func (f Flags) Match(v any) bool {
	flags, ok := v.(utils.Flags)
	if !ok {
		return false
	}

	flagsStr := strings.TrimSpace(strings.ToLower(flags.String()))
	for flag := range strings.SplitSeq(flagsStr, ",") {
		if flag == f.string {
			return true
		}
	}
	return false
}

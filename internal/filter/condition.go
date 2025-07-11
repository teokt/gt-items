package filter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/teokt/gt-items/internal/utils"
)

type Condition interface {
	Match(value any) bool
}

type Or []Condition

func (o Or) Match(val any) bool {
	for _, cond := range o {
		if cond.Match(val) {
			return true
		}
	}
	return false
}

type And []Condition

func (a And) Match(val any) bool {
	for _, cond := range a {
		if !cond.Match(val) {
			return false
		}
	}
	return true
}

type Not struct{ Condition }

func (n Not) Match(val any) bool {
	return !n.Condition.Match(val)
}

type NumericRange struct {
	Min *int
	Max *int
}

func (r NumericRange) Match(v any) bool {
	n := utils.ToInt(v)
	if r.Min != nil && n < *r.Min {
		return false
	}
	if r.Max != nil && n > *r.Max {
		return false
	}
	return true
}

type StringValue string

func (sv StringValue) Match(v any) bool {
	str := strings.ToLower(string(sv))

	if utils.IsFlag(v) {
		s, ok := v.(fmt.Stringer)
		if !ok {
			return false
		}
		flags := strings.SplitSeq(strings.ToLower(s.String()), ",")
		for flag := range flags {
			if strings.TrimSpace(flag) == str {
				return true
			}
		}
		return false
	}

	if utils.IsEnum(v) {
		s, ok := v.(fmt.Stringer)
		if !ok {
			return false
		}
		return strings.ToLower(s.String()) == str
	}

	if s, ok := v.(string); ok {
		return strings.Contains(strings.ToLower(s), str)
	}

	n, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return utils.ToInt(v) == n
}

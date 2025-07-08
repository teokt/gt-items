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

type NumericValue int

func (v NumericValue) Match(val any) bool {
	return int(v) == utils.ToInt(val)
}

type NumericRange struct {
	Min *int
	Max *int
}

func (r NumericRange) Match(val any) bool {
	n := utils.ToInt(val)
	if r.Min != nil && n < *r.Min {
		return false
	}
	if r.Max != nil && n > *r.Max {
		return false
	}
	return true
}

type StringValue string

func (sv StringValue) Match(val any) bool {
	str := strings.ToLower(string(sv))

	if utils.IsFlag(val) {
		s, ok := val.(fmt.Stringer)
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

	if utils.IsEnum(val) {
		s, ok := val.(fmt.Stringer)
		if !ok {
			return false
		}
		return strings.ToLower(s.String()) == str
	}

	if s, ok := val.(string); ok {
		return strings.Contains(strings.ToLower(s), str)
	}

	n, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return utils.ToInt(val) == n
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

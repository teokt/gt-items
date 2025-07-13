package filter

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/teokt/gt-items/internal/utils"
)

func ParseExpression(typ FieldType, expr string) (Condition, error) {
	if expr == "" {
		return nil, fmt.Errorf("empty expression")
	}

	if expr[0] == '(' && expr[len(expr)-1] == ')' {
		depth := 1
		for i := 1; i < len(expr)-1; i++ {
			if expr[i] == '(' {
				depth++
			} else if expr[i] == ')' {
				depth--
				if depth == 0 {
					break
				}
			}
		}
		if depth == 1 {
			return ParseExpression(typ, expr[1:len(expr)-1])
		}
	}

	if strings.HasPrefix(expr, "!") {
		inner, err := ParseExpression(typ, expr[1:])
		if err != nil {
			return nil, err
		}
		return Not{inner}, nil
	}

	orIndex := -1
	andIndex := -1
	depth := 0

	for i, ch := range expr {
		switch ch {
		case '(':
			depth++
		case ')':
			depth--
		case '|':
			if depth == 0 && orIndex == -1 {
				orIndex = i
			}
		case '&':
			if depth == 0 && andIndex == -1 {
				andIndex = i
			}
		}
	}

	if orIndex != -1 {
		left, err := ParseExpression(typ, expr[:orIndex])
		if err != nil {
			return nil, err
		}
		right, err := ParseExpression(typ, expr[orIndex+1:])
		if err != nil {
			return nil, err
		}
		return Or{left, right}, nil
	}

	if andIndex != -1 {
		left, err := ParseExpression(typ, expr[:andIndex])
		if err != nil {
			return nil, err
		}
		right, err := ParseExpression(typ, expr[andIndex+1:])
		if err != nil {
			return nil, err
		}
		return And{left, right}, nil
	}

	expr = strings.ToLower(expr)

	switch typ {
	case FieldTypeString:
		return String{expr}, nil

	case FieldTypeEnum:
		return Enum{expr}, nil

	case FieldTypeFlags:
		return Flags{expr}, nil

	case FieldTypeInt:
		expr = strings.TrimSpace(expr)
		if strings.Contains(expr, ":") {
			parts := strings.Split(expr, ":")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid range format")
			}

			r := IntRange{}
			if parts[0] != "" {
				min, err := strconv.Atoi(parts[0])
				if err != nil {
					return nil, fmt.Errorf("invalid min value: %s", parts[0])
				}
				r.Min = &min
			}
			if parts[1] != "" {
				max, err := strconv.Atoi(parts[1])
				if err != nil {
					return nil, fmt.Errorf("invalid max value: %s", parts[1])
				}
				r.Max = &max
			}
			return r, nil
		}

		return Int{utils.ToInt(expr)}, nil

	default:
		return nil, errors.New("unknown field type")
	}
}

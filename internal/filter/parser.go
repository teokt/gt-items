package filter

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseExpression(expr string) (Condition, error) {
	expr = strings.TrimSpace(expr)
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
			return ParseExpression(expr[1 : len(expr)-1])
		}
	}

	if strings.HasPrefix(expr, "!") {
		inner, err := ParseExpression(expr[1:])
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
		left, err := ParseExpression(expr[:orIndex])
		if err != nil {
			return nil, err
		}
		right, err := ParseExpression(expr[orIndex+1:])
		if err != nil {
			return nil, err
		}
		return Or{left, right}, nil
	}

	if andIndex != -1 {
		left, err := ParseExpression(expr[:andIndex])
		if err != nil {
			return nil, err
		}
		right, err := ParseExpression(expr[andIndex+1:])
		if err != nil {
			return nil, err
		}
		return And{left, right}, nil
	}

	if strings.Contains(expr, ":") {
		parts := strings.Split(expr, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid range format")
		}

		r := NumericRange{}
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

	return StringValue(expr), nil
}

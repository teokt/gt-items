package utils

import (
	"reflect"
	"strconv"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Enum interface {
	IsEnum() bool
	String() string
}

type Flags interface {
	IsFlags() bool
	String() string
}

func IsInt(t reflect.Kind) bool {
	switch t {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	default:
		return false
	}
}

func ToInt[T any](v T) int {
	switch val := any(v).(type) {
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(val).Int())
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(val).Uint())
	case float32, float64:
		return int(reflect.ValueOf(val).Float())
	case string:
		n, err := strconv.Atoi(val)
		if err == nil {
			return n
		}
	}
	return 0
}

package utils

import (
	"reflect"
)

func ToInt[T any](v T) int {
	switch val := any(v).(type) {
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(val).Int())
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(val).Uint())
	case float32, float64:
		return int(reflect.ValueOf(val).Float())
	default:
		return 0
	}
}

func HasMethod(v any, method string) bool {
	val := reflect.Indirect(reflect.ValueOf(v))
	return val.MethodByName(method).IsValid()
}

func IsEnum(v any) bool {
	return HasMethod(v, "IsEnum")
}

func IsFlag(v any) bool {
	return HasMethod(v, "IsFlag")
}

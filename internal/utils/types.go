package utils

import (
	"reflect"
)

func ToInt[T any](val T) int {
	switch v := any(val).(type) {
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(v).Int())
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(v).Uint())
	case float32, float64:
		return int(reflect.ValueOf(v).Float())
	default:
		return 0
	}
}

func HasMethod(val any, method string) bool {
	v := reflect.ValueOf(val)
	return v.MethodByName(method).IsValid()
}

func IsEnum(val any) bool {
	return HasMethod(val, "IsEnum")
}

func IsFlag(val any) bool {
	return HasMethod(val, "IsFlag")
}

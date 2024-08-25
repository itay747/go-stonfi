package utils

import (
	"net/url"
	"reflect"
	"strings"
	"unicode"
)

// CamelCase converts a string from snake_case to CamelCase.
func CamelCase(s string) string {
	var result strings.Builder
	capitalize := false
	for _, r := range s {
		switch {
		case r == '_':
			capitalize = true
		case capitalize:
			result.WriteRune(unicode.ToUpper(r))
			capitalize = false
		default:
			result.WriteRune(unicode.ToLower(r))
		}
	}
	return result.String()
}

// CamelCaseKeys converts all keys in a map or slice from snake_case to camelCase recursively.
func CamelCaseKeys(data interface{}) interface{} {
	switch data := data.(type) {
	case map[string]interface{}:
		camelCased := make(map[string]interface{}, len(data))
		for k, v := range data {
			camelCased[CamelCase(k)] = CamelCaseKeys(v)
		}
		return camelCased
	case []interface{}:
		for i, v := range data {
			data[i] = CamelCaseKeys(v)
		}
	}
	return data
}

// SnakeCaseKeys converts all keys in a map or slice from camelCase to snake_case recursively.
func SnakeCaseKeys(data map[string]string) map[string]string {
	snakeCased := make(map[string]string)
	for k, v := range data {
		snakeCased[SnakeCase(k)] = url.QueryEscape(v)
	}
	return snakeCased
}

// DenullifyValues recursively replaces nil pointers and nil interfaces in any complex data structure (maps, slices, etc.).
func DenullifyValues(data interface{}) interface{} {
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Map:
		for _, key := range v.MapKeys() {
			val := v.MapIndex(key)
			if val.IsNil() {
				v.SetMapIndex(key, reflect.Zero(val.Type()))
			} else {
				newVal := DenullifyValues(val.Interface())
				v.SetMapIndex(key, reflect.ValueOf(newVal))
			}
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			elem := v.Index(i)
			if !elem.IsNil() {
				newElem := DenullifyValues(elem.Interface())
				elem.Set(reflect.ValueOf(newElem))
			}
		}
	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			return nil
		}
		newV := DenullifyValues(v.Elem().Interface())
		return reflect.New(reflect.TypeOf(newV)).Elem().Interface()
	}

	return data
}

// SnakeCase converts a string from camelCase to snake_case.
func SnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteByte('_')
			}
			result.WriteByte(byte(unicode.ToLower(r)))
		} else {
			result.WriteByte(byte(r))
		}
	}
	return result.String()
}
func DecamelizeKeys(data map[string][]string) map[string][]string {
	result := make(map[string][]string)
	for k, v := range data {
		result[SnakeCase(k)] = v
	}
	return result
}
func ToUrlSafe(s string) string {
	return url.QueryEscape(s)
}

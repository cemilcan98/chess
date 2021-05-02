package builder

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

func appendIfHasVal(m bson.M, key string, val interface{}) {
	if !IsNil(val) {
		m[key] = val
	}
}
func appendIfHasValByElement(e bson.E, val interface{}) {
	if !IsNil(val) {
		e.Value = val
	}
}

func IsNil(val interface{}) (result bool) {

	if val == nil {
		return true
	}

	switch v := reflect.ValueOf(val); v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer,
		reflect.Interface, reflect.Slice:
		return v.IsNil()
	}

	return
}

func AnyNil(values ...interface{}) bool {
	for _, val := range values {

		if IsNil(val) {
			return true
		}

	}

	return false
}

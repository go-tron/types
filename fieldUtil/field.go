package fieldUtil

import "reflect"

func IsEmpty(s interface{}) bool {
	if s == nil {
		return true
	}
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return true
		}
		return v.Elem().IsZero()
	} else if v.Kind() == reflect.Slice {
		return v.Len() == 0
	}
	return v.IsZero()
}

package structUtil

import (
	"github.com/go-tron/base-error"
	"reflect"
)

var ErrorFieldInvalid = baseError.SystemFactoryStack(3, "2101", "field:{} invalid")
var ErrorFieldCantSet = baseError.SystemFactoryStack(3, "2102", "field:{} can't set")

func SetValue(obj interface{}, field string, value interface{}) error {
	valueV := reflect.ValueOf(obj)
	if valueV.Kind() == reflect.Ptr {
		valueV = valueV.Elem()
	}

	fieldV := valueV.FieldByName(field)
	if !fieldV.IsValid() {
		return ErrorFieldInvalid(field)
	}

	if !fieldV.CanSet() {
		return ErrorFieldCantSet(field)
	}

	fieldV.Set(reflect.ValueOf(value))
	return nil
}

func GetValue(obj interface{}, field string) (interface{}, error) {
	valueV := reflect.ValueOf(obj)
	if valueV.Kind() == reflect.Ptr {
		valueV = valueV.Elem()
	}

	fieldV := valueV.FieldByName(field)
	if !fieldV.IsValid() {
		return nil, ErrorFieldInvalid(field)
	}
	return fieldV.Interface(), nil
}

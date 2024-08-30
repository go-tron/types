package structUtil

import (
	"encoding/json"
	"fmt"
	"github.com/go-tron/types/stringUtil"
	"github.com/thoas/go-funk"
	"reflect"
	"sort"
)

type ToSortStringOption func(*ToSortStringConfig)

type ToSortStringConfig struct {
	Exclude        []string
	IgnoreEmptyStr bool
	AnonymousField bool
}

func ToSortStringWithExclude(val ...string) ToSortStringOption {
	return func(opts *ToSortStringConfig) {
		opts.Exclude = append(opts.Exclude, val...)
	}
}
func ToSortStringWithIgnoreEmptyStr() ToSortStringOption {
	return func(opts *ToSortStringConfig) {
		opts.IgnoreEmptyStr = true
	}
}

func ToSortStringWithAnonymousField() ToSortStringOption {
	return func(opts *ToSortStringConfig) {
		opts.AnonymousField = true
	}
}

func ToSortString(obj interface{}, opts ...ToSortStringOption) string {
	if obj == nil {
		return ""
	}
	c := &ToSortStringConfig{}
	for _, apply := range opts {
		if apply != nil {
			apply(c)
		}
	}

	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}
		v = v.Elem()
	}
	t := v.Type()

	keys := SortKeys(t, c)

	str := ""
	for _, key := range keys {
		name := stringUtil.FirstCharToLower(key)
		field := v.FieldByName(key)
		//fmt.Printf("%s:%s:%v\n", name, field.Type(), field.Interface())
		var value string
		if field.IsZero() {
			value = ""
		} else {
			if field.Kind() == reflect.Ptr {
				field = field.Elem()
			}
			if field.Kind() == reflect.Map || field.Kind() == reflect.Slice || field.Kind() == reflect.Array {
				byteVal, err := json.Marshal(field.Interface())
				if err == nil {
					value = string(byteVal)
				} else {
					value = fmt.Sprint(field.Interface())
				}
			} else {
				value = fmt.Sprint(field.Interface())
			}
		}
		if c.IgnoreEmptyStr && value == "" {
			continue
		}
		if str != "" {
			str += "&"
		}
		str += name + "=" + value
	}
	return str
}

func SortKeys(t reflect.Type, c *ToSortStringConfig) []string {
	var keys []string
	for k := 0; k < t.NumField(); k++ {
		if len(c.Exclude) > 0 && funk.ContainsString(c.Exclude, t.Field(k).Name) {
			continue
		}
		if t.Field(k).Anonymous {
			if c.AnonymousField {
				res := SortKeys(t.Field(k).Type, c)
				keys = append(keys, res...)
			}
		} else {
			keys = append(keys, t.Field(k).Name)
		}
	}
	sort.Strings(keys)
	return keys
}

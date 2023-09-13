package structUtil

import (
	"encoding/json"
	"fmt"
	"github.com/go-tron/types-util/stringUtil"
	"github.com/thoas/go-funk"
	"reflect"
	"sort"
)

type ToSortStringOption func(*ToSortStringConfig)

type ToSortStringConfig struct {
	Exclude        []string
	IgnoreEmptyStr bool
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

	var keys []string
	for k := 0; k < t.NumField(); k++ {
		if len(c.Exclude) > 0 && funk.ContainsString(c.Exclude, t.Field(k).Name) {
			continue
		}
		keys = append(keys, t.Field(k).Name)
	}
	sort.Strings(keys)

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

// ToSortStringIncludeAnonymous 处理匿名属性的排序方法
// type A struct{Age int}  type B struct {A Name string} ,var aa = B{Name:'xx',{Age:18}}
// 预期得到字符串 age=18&name=xx, ToSortString 得到 A={18}&name=xx
func ToSortStringIncludeAnonymous(obj interface{}, exclude ...string) string {
	if obj == nil {
		return ""
	}
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}
		v = v.Elem()
	}
	t := v.Type()

	keys := SortKeys(t, exclude...)

	str := ""
	for _, key := range keys {
		if str != "" {
			str += "&"
		}
		field := v.FieldByName(key)
		name := stringUtil.FirstCharToLower(key)

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
		str += name + "=" + value
	}
	return str
}

func SortKeys(t reflect.Type, exclude ...string) []string {
	var keys []string
	for k := 0; k < t.NumField(); k++ {
		if len(exclude) > 0 && funk.ContainsString(exclude, t.Field(k).Name) {
			continue
		}
		if t.Field(k).Anonymous {
			res := SortKeys(t.Field(k).Type, exclude...)
			keys = append(keys, res...)
		} else {
			keys = append(keys, t.Field(k).Name)
		}
	}

	sort.Strings(keys)

	return keys
}

package mapUtil

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
)

func ToSortString(obj map[string]interface{}) string {
	var keys []string
	for k := range obj {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var str = ""
	for _, key := range keys {
		if str != "" {
			str += "&"
		}
		var value string
		if obj[key] == nil {
			value = ""
		} else {
			kind := reflect.TypeOf(obj[key]).Kind()
			if kind == reflect.Map || kind == reflect.Slice || kind == reflect.Array {
				byteVal, err := json.Marshal(obj[key])
				if err == nil {
					value = string(byteVal)
				} else {
					value = fmt.Sprint(obj[key])
				}
			} else {
				value = fmt.Sprint(obj[key])
			}
		}
		str += key + "=" + value
	}
	return str
}

func FromStruct(from interface{}) (map[string]interface{}, error) {
	var m = make(map[string]interface{})
	tmp, err := json.Marshal(from)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(tmp, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func ToStruct(from interface{}, to interface{}) error {
	//return mapstructure.Decode(from, to)
	tmp, err := json.Marshal(from)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(tmp, to); err != nil {
		return err
	}
	return nil
}

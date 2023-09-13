package jsonUtil

import (
	"encoding/json"
	"strconv"
	"strings"
)

func NewIntArray(arr []int) StringArray {
	temp := ""
	for _, val := range arr {
		if temp != "" {
			temp += ","
		}

		temp = temp + strconv.Itoa(val)
	}
	return StringArray(temp)
}

type StringArray string

func (sa StringArray) MarshalJSON() ([]byte, error) {
	return []byte(`[` + sa + `]`), nil
}

func (sa *StringArray) UnmarshalJSON(b []byte) error {
	tmp := string(b)
	tmp = strings.TrimPrefix(tmp, "[")
	tmp = strings.TrimSuffix(tmp, "]")
	*sa = StringArray(tmp)
	return nil
}

func (sa StringArray) ToString() string {
	return string(sa)
}

func (sa StringArray) ToStringArray() []string {
	return strings.Split(string(sa), ",")
}

func (sa StringArray) ToIntArray() ([]int, error) {
	var arr = make([]int, 0)
	if string(sa) == "" {
		return arr, nil
	}
	for _, val := range strings.Split(string(sa), ",") {
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		arr = append(arr, i)
	}
	return arr, nil
}

func (sa StringArray) ToStringArrayArray() ([][]string, error) {
	var arr = make([][]string, 0)
	if string(sa) == "" || string(sa) == `""` {
		return arr, nil
	}
	if err := json.Unmarshal([]byte(`[`+sa+`]`), &arr); err != nil {
		return nil, err
	}
	return arr, nil
}

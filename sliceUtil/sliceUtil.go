package sliceUtil

import "github.com/thoas/go-funk"

func ContainsInt(a []int, b []int) bool {
	for _, v := range a {
		if funk.ContainsInt(b, v) {
			return true
		}
	}
	return false
}

func ContainsString(a []string, b []string) bool {
	for _, v := range a {
		if funk.ContainsString(b, v) {
			return true
		}
	}
	return false
}

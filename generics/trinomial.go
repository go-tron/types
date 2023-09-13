package generics

func Ternary[T any](condition bool, s1 T, s2 T) T {
	if condition {
		return s1
	} else {
		return s2
	}
}

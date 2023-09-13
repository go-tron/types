package stringUtil

import (
	"fmt"
	"unicode"
)

func FirstCharToLower(name string) string {
	if len(name) < 1 {
		return ""
	}
	strRune := []rune(name)
	strRune[0] = unicode.ToLower(strRune[0])
	r := string(strRune)
	return r
}

func FirstCharToUpper(name string) string {
	if len(name) < 1 {
		return ""
	}
	strRune := []rune(name)
	strRune[0] = unicode.ToUpper(strRune[0])
	r := string(strRune)
	return r
}

func AsteriskPhone(number string) string {
	if len(number) < 11 {
		return number
	}
	return fmt.Sprintf("%s****%s", number[0:3], number[len(number)-4:])
}

func AsteriskBankCard(number string) string {
	if len(number) <= 16 {
		return number
	}
	return fmt.Sprintf("%s **** **** %s", number[0:4], number[len(number)-4:])
}

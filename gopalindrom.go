package gopalindrom

import "strings"

func IsPalindrom(text string) bool {
	var arr []string
	arr = strings.Split(text, "")
	for i, val := range arr {
		if val != arr[(len(arr)-1)-i] {
			return false
		}
	}

	return true
}

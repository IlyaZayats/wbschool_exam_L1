package main

import (
	"fmt"
	"unicode"
)

func isUniqueRunes(str string) bool {
	set := map[rune]struct{}{}
	runeStr := []rune(str)
	for i := 0; i < len(runeStr); i++ {
		key := unicode.ToLower(runeStr[i])
		if _, ok := set[key]; ok == false {
			set[key] = struct{}{}
		} else {
			return false
		}
	}
	return true
}

func main() {
	strSlice := []string{"abcd", "abCdefAaf", "aabcd", "Aabcd", "абвг", "абвА"}
	for _, item := range strSlice {
		fmt.Printf("string: %s\nresult: %v\n", item, isUniqueRunes(item))
	}
}

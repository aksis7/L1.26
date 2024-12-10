package main

import (
	"fmt"
	"strings"
)

func allUnique(input string) bool {
	// Приводим строку к нижнему регистру
	input = strings.ToLower(input)

	// Используем карту для проверки уникальности символов
	charMap := make(map[rune]bool)

	for _, char := range input {
		if charMap[char] {
			return false // Если символ уже был, возвращаем false
		}
		charMap[char] = true
	}
	return true // Все символы уникальны
}

func main() {
	testCases := []string{"abcd", "aabCdefAaf", "aabcd"}

	for _, testCase := range testCases {
		fmt.Printf("Input: %s, Unique: %t\n", testCase, allUnique(testCase))
	}
}

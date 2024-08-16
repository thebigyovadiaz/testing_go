package iteration

import (
	"fmt"
	"strings"
)

func Repeat(char string, count int) string {
	var newString string

	if count <= 0 {
		return char
	}

	otherStr := strings.Repeat(char, count)
	fmt.Println(otherStr)

	for i := 0; i < count; i++ {
		newString += char
	}

	return newString
}

func IterArr(arr []int) int {
	acc := 1

	if len(arr) == 0 {
		return 0
	}

	for _, v := range arr {
		acc *= v
	}

	return acc
}

package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 5, 6, 8}

	fmt.Println(chop(8, arr))
}

func chop(needle int, haystack []int) int {
	if needle > haystack[len(haystack)-1] {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		if needle == haystack[i] {
			return i
		}
	}

	return -1
}

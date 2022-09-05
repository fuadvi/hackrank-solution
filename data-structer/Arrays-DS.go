package main

import "fmt"

func reverseArray(a []int32) []int32 {
	// Write your code here
	var result []int32
	for i := len(a) - 1; i >= 0; i-- {
		result = append(result, a[i])
	}
	return result
}

func main() {
	data := []int32{1, 2, 4}
	result := reverseArray(data)
	fmt.Println(result)
}

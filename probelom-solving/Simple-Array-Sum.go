package main

import "fmt"

func simpleArraySum(ar []int32) int32 {
	// Write your code here
	var result int32

	for _, value := range ar {
		result += value
	}

	return result
}

func main() {
	fmt.Println(simpleArraySum([]int32{1, 2, 3, 4}))
}

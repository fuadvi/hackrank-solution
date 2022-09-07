package main

import "fmt"

func staircase(n int32) {
	// Write your code here
	var i, j, k int32

	for i = 0; i < n; i++ {

		for j = 0; j < n-i-1; j++ {
			fmt.Print("*")
		}
		for k = 0; k < i+1; k++ {
			fmt.Print("#")
		}
		fmt.Println("")

	}
}

func main() {
	staircase(3)
}

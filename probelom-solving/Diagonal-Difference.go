package main

import "fmt"

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here

	var left = arr[0][0]
	var right = arr[0][len(arr)-1]
	fmt.Println("result right", right)
	fmt.Println("result left", left)

	for i := 1; i < len(arr); i++ {
		// fmt.Println(arr[i])

		for j := i; j < len(arr); j++ {
			// fmt.Println(arr[i][j])
			left += arr[i][j]
			break
		}

		for j := len(arr) - i; j > 0; j-- {
			fmt.Println(arr[i][j])
			right += arr[i][j-1]
			break
		}
	}

	fmt.Println("result right", right)
	fmt.Println("result left", left)
	result := right - left

	if result < 0 {
		return result * -1
	}

	return right - left
}

func main() {

	data := [][]int32{
		{6, 6, 7, -10, 9, -3, 8, 9, -1},
		{9, 7, -10, 6, 4, 1, 6, 1, 1},
		{-1, -2, 4, -6, 1, -4, -6, 3, 9},
		{-8, 7, 6, -1, -6, -6, 6, -7, 2},
		{-10, -4, 9, 1, -7, 8, -5, 3, -5},
		{-8, -3, -4, 2, -3, 7, -5, 1, -5},
		{-2, -7, -4, 8, 3, -1, 8, 2, 3},
		{-3, 4, 6, -7, -7, -8, -3, 9, -6},
		{-2, 0, 5, 4, 4, 4, -3, 3, 0},
	}

	fmt.Println(diagonalDifference(data))

}

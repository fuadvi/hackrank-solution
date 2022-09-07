package main

import "fmt"

func miniMaxSum(arr []int32) {
	// Write your code here
	var max, min int64

	exc := arr[0]

	for _, v := range arr {
		if v < exc {
			exc = v
		}
	}

	for i := 0; i < len(arr)-1; i++ {

		min += int64(arr[i])
	}

	for j := 0; j < len(arr); j++ {
		if int(exc) == int(arr[j]) {
			fmt.Println("ss", arr[j])
			continue
		}
		max += int64(arr[j])
	}

	fmt.Println(min, max, exc)

}

func main() {
	arr := []int32{7, 69, 2, 221, 8974}
	miniMaxSum(arr)
}

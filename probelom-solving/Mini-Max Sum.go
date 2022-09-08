package main

import "fmt"

// [1,2,3,4,5]
// apabila hitung jumalah kecuali nilai min yang di dalam array
// [2,3,4,5]

func miniMaxSum(arr []int32) {
	// Write your code here
	var max, min int64

	for i := 0; i < len(arr); i++ {
		// fmt.Print(arr[0])
		for j := 0; j < len(arr)-1; j++ {
			temp := arr[j]
			if arr[j] > arr[j+1] {
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
			// fmt.Println(" ")
		}
	}

	for i := 1; i < len(arr); i++ {

		min += int64(arr[i])
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			temp := arr[j]
			if arr[j] < arr[j+1] {
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

	for j := 1; j < len(arr); j++ {
		max += int64(arr[j])
	}

	fmt.Println(max, min)

}

func main() {
	arr := []int32{1, 2, 3, 4, 5}
	miniMaxSum(arr)
}

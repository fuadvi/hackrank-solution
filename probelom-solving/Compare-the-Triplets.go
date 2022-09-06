package main

import "fmt"

func compareTriplets(a []int32, b []int32) []int32 {
	result := make([]int32, 2)
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			result[0] += 1
		} else if a[i] < b[i] {
			result[1] += 1
		} else {
			continue
		}

		fmt.Println(i)
	}
	return result
}

func main() {
	fmt.Println(compareTriplets([]int32{5, 6, 7}, []int32{3, 6, 10}))
}

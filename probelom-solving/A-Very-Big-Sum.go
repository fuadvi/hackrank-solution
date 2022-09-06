package main

import "fmt"

func aVeryBigSum(ar []int64) int64 {
	// Write your code here
	var result int64
	for _, v := range ar {
		result += v
	}
	return result

}

func main() {
	fmt.Println(aVeryBigSum([]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}))
}

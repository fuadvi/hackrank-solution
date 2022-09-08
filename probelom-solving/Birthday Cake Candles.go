package main

import "fmt"

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var Max int32
	var result int32
	for _, v := range candles {

		if v > Max {
			Max = v
		}
	}

	for _, v := range candles {
		if Max == v {
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(birthdayCakeCandles([]int32{4, 4, 3, 1}))
}

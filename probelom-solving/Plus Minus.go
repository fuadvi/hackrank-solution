package main

import "fmt"

func plusMinus(arr []int32) {
	// Write your code here
	var positif float32
	var negatif float32
	var zero float32

	for _, v := range arr {

		if v == 0 {
			zero++
		} else if v < 0 {
			negatif++
		} else {
			positif++
		}
	}

	fmt.Println(positif / float32(len(arr)))
	fmt.Println(negatif / float32(len(arr)))
	fmt.Println(zero / float32(len(arr)))
	fmt.Println(positif)
	fmt.Println(negatif)
	fmt.Println(zero)
}

func main() {
	arr := []int32{-4, 3, -9, 0, 4, 1}
	plusMinus(arr)
}

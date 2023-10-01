package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{10, 8, 9, 2, 3, 9, 19}

	maxValue := -math.MaxInt32
	secondMaxValue := -math.MaxInt32

	for i := 0; i < len(arr); i++ {
		if arr[i] > maxValue {
			secondMaxValue = maxValue
			maxValue = arr[i]
		} else if arr[i] > secondMaxValue {
			secondMaxValue = arr[i]
		}
	}

	fmt.Println("Max element of array: ", maxValue)
	fmt.Println("Second largest element of array: ", secondMaxValue)
}

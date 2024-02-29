package main

import (
	"fmt"
	"math"
)

func kadaneAlgorithm(arr []int) int {
	maxEndingHere := arr[0]
	maxSoFar := arr[0]

	for i := 1; i < len(arr); i++ {
		maxEndingHere = int(math.Max(float64(arr[i]), float64(maxEndingHere+arr[i])))
		maxSoFar = int(math.Max(float64(maxSoFar), float64(maxEndingHere)))
	}

	return maxSoFar
}

func main() {
	// Example usage
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := kadaneAlgorithm(arr)
	fmt.Printf("Maximum subarray sum: %d\n", result)
}

package main

import (
	"fmt"
	"math"
)

// Complete the diagonalDifference function below.
func diagonalDifference(arr [][]int32) int32 {
	s := 4
	size := int(s)

	var ld, rd []int32

	for i, item := range arr {
		n := int(i)
		ld = append(ld, item[i])
		rd = append(rd, item[size-1-n])
	}

	var ldsum, rdsum int

	for index := 0; index < size; index++ {

		lnum := int(ld[index])
		rnum := int(rd[index])

		ldsum += lnum
		rdsum += rnum

	}
	abs := int32(math.Abs(float64(ldsum) - float64(rdsum)))
	fmt.Println(abs)
	return abs
}

// Hacker rank simulator.
func main() {

	matrix := [][]int32{{11, 2, 4, 5}, {2, 4, 5, 6}, {5, 10, 8, -12}, {15, 1, 4, -2}}
	diagonalDifference(matrix)
	// fmt.Println(res)
}

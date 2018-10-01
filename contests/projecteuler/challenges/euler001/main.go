package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		// Convert string to int before doing math.
		n, _ := strconv.Atoi(line)
		// use math to find if any of the natural number below n are multiples of 3 or 5.
		sum := findMultsof35(n)
		// If the sum is positive, print the sum.
		if sum != 0 {
			fmt.Println(sum)
		}

	}

}

// find natural numbers below n that are multiples of 3 or 5.
func findMultsof35(n int) (sum int) {
	var slice []int
	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			slice = append(slice, i)
		}
	}
	return sumSlice(slice)
}

// Sum all the numbers ina  slice.
func sumSlice(nums []int) (r int) {
	for i := range nums {
		r += nums[i]
	}
	return r
}

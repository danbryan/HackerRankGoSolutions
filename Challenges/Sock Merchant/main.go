package main

import (
	"fmt"
)

// sockMerchant takes an inventory list, and returns number of matching socks.
func sockMerchant(n int32, ar []int32) (pairsofsocks int32) {

	// Get length of the array to validate input.
	len := int32(len(ar))
	if n == len {

		// m is a map of duplcates
		m := dupeMap(ar)

		// for each value in map, divide by 2 to get pairs.
		for _, v := range m {
			pair := int(v / 2)

			// for each number of pair found increment pairs of socks.
			for index := 0; index < pair; index++ {
				pairsofsocks++
			}
		}
	} else {
		fmt.Println("Invalid Input")
	}
	return pairsofsocks
}

// dupeMap is a map of sock id's to number in inventory.
func dupeMap(ar []int32) map[int32]int32 {

	dupes := make(map[int32]int32)

	// count how many of each type of sock are in inventory.
	for _, item := range ar {
		_, exist := dupes[item]
		// 1 the item is in the inventory already, increase it
		if exist {
			dupes[item]++
			// if not start it at 1
		} else {
			dupes[item] = 1
		}
	}
	return dupes
}

// Hacker rank simulator.
func main() {
	nums := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	len := int32(len(nums))
	sockMerchant(len, nums)
}

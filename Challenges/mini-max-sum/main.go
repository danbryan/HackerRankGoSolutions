package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {

	// Sort the Slice from smallest to largest.
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	// min adds all but last item.
	min := arraySum(arr[:len(arr)-1])

	// max adds all but first item.
	max := arraySum(arr[1:])

	// Print min and max.
	fmt.Printf("%v %v\n", min, max)
}

func arraySum(ar []int32) int {
	var i int
	for _, item := range ar {
		i += int(item)
	}
	return i
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

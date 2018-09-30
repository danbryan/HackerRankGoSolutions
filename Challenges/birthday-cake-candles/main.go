package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the birthdayCakeCandles function below.
func birthdayCakeCandles(ar []int32) int32 {
	var matches int32
	ln := findLargestInt(ar)
	for _, item := range ar {
		if item == ln {
			matches++
		}
	}
	fmt.Println(matches)
	return matches
}

func findLargestInt(slice []int32) int32 {
	var largestNumber int32
	for _, item := range slice {
		if item >= largestNumber {
			largestNumber = item
		}
	}
	return largestNumber
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// file, err := os.Open("output") // For read access.
	stdout, err := os.Create("output.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	arCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int32

	for i := 0; i < int(arCount); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := birthdayCakeCandles(ar)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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

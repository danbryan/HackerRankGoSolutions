package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {

	timeslice := strings.Split(s, ":")

	hh := timeslice[0]
	mm := timeslice[1]
	ss := timeslice[2]

	if strings.Contains(ss, "PM") {
		ssslice := strings.Split(ss, "PM")
		ss = ssslice[0]
		if hh != "12" {
			mhh, _ := strconv.Atoi(hh)
			mhh += 12
			hh = strconv.Itoa(mhh)
		}
	} else if strings.Contains(ss, "AM") {
		ssslice := strings.Split(ss, "AM")
		ss = ssslice[0]
		if hh == "12" {
			hh = "00"
		}
	}

	miltime := hh + ":" + mm + ":" + ss
	fmt.Println("MT: ", miltime)
	return miltime
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create("OUTPUT_PATH")
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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

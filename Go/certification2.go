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

func SortOddEven(arr []int32) ([]int32, []int32) {
	odds := []int32{}
	evens := []int32{}

	for _, num := range arr {
		if num%2 == 0 {
			evens = append(evens, num)
		} else {
			odds = append(odds, num)
		}
	}

	sort.Slice(odds, func(i, j int) bool {
		return odds[i] < odds[j]
	})

	sort.Slice(evens, func(i, j int) bool {
		return evens[i] < evens[j]
	})

	return odds, evens
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrCount, _ := strconv.Atoi(readLine(reader))

	arr := make([]int32, arrCount)

	for i := 0; i < arrCount; i++ {
		item, _ := strconv.ParseInt(readLine(reader), 10, 64)
		arr[i] = int32(item)
	}

	odds, evens := SortOddEven(arr)

	for _, odd := range odds {
		fmt.Println(odd)
	}

	for _, even := range evens {
		fmt.Println(even)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}


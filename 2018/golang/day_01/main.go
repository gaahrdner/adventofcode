package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	frequencyFile, err := os.Open("input")
	check(err)

	var frequencies []int

	scanner := bufio.NewScanner(frequencyFile)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		frequencies = append(frequencies, num)
	}

	check(scanner.Err())

	sum := 0
	for _, num := range frequencies {
		sum += num
	}

	fmt.Println("Frequency:", sum)

	duplicateFrequency := make(map[int]int)

	freq := 0
	for i := 0; ; i = (i + 1) % len(frequencies) {
		freq += frequencies[i]
		duplicateFrequency[freq]++

		if duplicateFrequency[freq] == 2 {
			fmt.Println("First duplicate frequency:", freq)
			break
		}
	}
}

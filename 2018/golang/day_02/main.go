package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	boxIDsFile, err := os.Open("input")
	check(err)

	// Part One
	scanner := bufio.NewScanner(boxIDsFile)
	twos, threes := 0, 0
	var boxIDs []string

	for scanner.Scan() {
		ID := scanner.Text()
		boxIDs = append(boxIDs, ID)
		characters := strings.Split(ID, "")

		characterMap := make(map[string]int)

		for _, character := range characters {
			characterMap[character]++
		}

		has2, has3 := false, false
		for _, count := range characterMap {
			switch count {
			case 3:
				has3 = true
			case 2:
				has2 = true
			}
		}

		if has2 {
			twos++
		}

		if has3 {
			threes++
		}
	}

	fmt.Println("Checksum:", twos*threes)

	// Part Two

	var firstCorrectBox, secondCorrectBox string

	for i := range boxIDs {
		for j := i + 1; j < len(boxIDs); j++ {
			if offByOne(boxIDs[i], boxIDs[j]) {
				firstCorrectBox, secondCorrectBox = boxIDs[i], boxIDs[j]
			}
		}
	}

	var commonLetters strings.Builder
	for index := 0; index < len(firstCorrectBox); index++ {
		if firstCorrectBox[index] == secondCorrectBox[index] {
			commonLetters.WriteByte(firstCorrectBox[index])
		}
	}
	fmt.Println("Common IDs between boxes:", commonLetters.String())

}

func offByOne(firstWord, secondWord string) bool {
	// should refactor this to return the differences
	differenceCount := 0
	for index := 0; index < len(firstWord); index++ {
		if firstWord[index] != secondWord[index] {
			differenceCount++
		}

		if differenceCount > 1 {
			return false
		}
	}
	return true
}

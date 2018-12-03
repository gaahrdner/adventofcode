package main

import (
	"bufio"
	"fmt"
	"os"
)

// Claim represents an elf's claim for the ideal area of fabric for Santa's suit
type Claim struct {
	ID     int
	Left   int
	Top    int
	Width  int
	Height int
}

func loadInput(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func convertClaimString(s string) Claim {
	// #3 @ 5,5: 2x2
	var claim Claim
	fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &claim.ID, &claim.Left, &claim.Top, &claim.Width, &claim.Height)
	return claim
}

func main() {

	input := loadInput("input")

	var claims []Claim

	for _, claim := range input {
		claims = append(claims, convertClaimString(claim))
	}

	fabric := [1000][1000]int{}
	squareInchesWithinTwoOrMoreClaims := 0

	// Part 1
	for _, claim := range claims {
		for y := claim.Left; y < claim.Left+claim.Width; y++ {
			for x := claim.Top; x < claim.Top+claim.Height; x++ {
				fabric[y][x]++
				if fabric[y][x] > 1 && fabric[y][x] <= 2 {
					squareInchesWithinTwoOrMoreClaims++
				}
			}
		}
	}

	// Part 2
	for _, claim := range claims {
		isOverlap := false
		for y := claim.Left; y < claim.Left+claim.Width; y++ {
			for x := claim.Top; x < claim.Top+claim.Height; x++ {
				if fabric[y][x] > 1 {
					isOverlap = true
				}
			}
		}

		if !isOverlap {
			fmt.Println("Non-overlapping claim ID:", claim.ID)
		}
	}

	fmt.Println("Square inches with two or more claims:", squareInchesWithinTwoOrMoreClaims)

}

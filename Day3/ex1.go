package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const symbols = "*/+=@=$#&^%-"

type Number struct {
	Value  int
	Length int
	X      int
	Y      int
	Found  bool
}

func NumInBetween(num, min, max int) bool {
	return num >= min && num <= max
}

func Ex1(input string) {
	file, _ := os.Open(input)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var result int
	var numbers []Number
	number := Number{Y: -1}
	// Find all numbers
	for y, line := range lines {
		for x, char := range line {
			if unicode.IsNumber(char) {
				if number.Y == -1 {
					number.Y = y
					number.X = x
				}
				number.Length++
				number.Value = number.Value*10 + int(char-'0')
				if x == len(line)-1 {
					numbers = append(numbers, number)
					number = Number{Y: -1}
				}
			} else {
				if number.Length > 0 {
					numbers = append(numbers, number)
					number = Number{Y: -1}
				}
			}
		}
	}

	// Find all symbols
	for y, line := range lines {
		for x, char := range line {
			if strings.ContainsRune(symbols, char) {
				fmt.Println(x, y, string(char))
				for i := range numbers {
					if NumInBetween(y, numbers[i].Y-1, numbers[i].Y+1) && NumInBetween(x, numbers[i].X-1, numbers[i].X+numbers[i].Length) && !numbers[i].Found {
						fmt.Println(numbers[i].Value)
						result += numbers[i].Value
						numbers[i].Found = true
					}
				}
			}
		}
	}

	fmt.Printf("Result: %d", result)
}

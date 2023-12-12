package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func Ex2(input string) {
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
			var foundNums []Number
			if char == '*' {
				for i := range numbers {
					if NumInBetween(y, numbers[i].Y-1, numbers[i].Y+1) && NumInBetween(x, numbers[i].X-1, numbers[i].X+numbers[i].Length) && !numbers[i].Found {
						foundNums = append(foundNums, numbers[i])
						numbers[i].Found = true
					}
				}
				if len(foundNums) == 2 {
					result += foundNums[0].Value * foundNums[1].Value
				}
			}
		}
	}

	fmt.Printf("Result: %d", result)
}

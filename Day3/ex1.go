package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const symbols = "*/+=@=$#&^%-"

func CheckIfPartNumber(numStart, numEnd, i int, lines []string, linesCount int) bool {
	for k := numStart; k <= numEnd; k++ {
		if strings.Contains(symbols, string(lines[i][k])) {
			return true
		}
		if i > 0 {
			if strings.Contains(symbols, string(lines[i-1][k])) {
				return true
			}
		}
		if i < linesCount-1 {
			if strings.Contains(symbols, string(lines[i+1][k])) {
				return true
			}
		}
	}
	return false
}

func Ex1(input string) {
	file, _ := os.Open(input)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	linesCount := len(lines)
	lineLength := len(lines[0])
	var numString string
	var numStart = -1
	var result int
	for i := 0; i < linesCount; i++ {
		for j := 0; j < lineLength; j++ {
			char := lines[i][j]
			if unicode.IsNumber(rune(char)) {
				numString += string(char)
				if numStart == -1 {
					if j > 0 {
						numStart = j - 1
					} else {
						numStart = j
					}
				}
				if j == lineLength-1 {
					numEnd := j
					num, _ := strconv.Atoi(numString)
					isValid := CheckIfPartNumber(numStart, numEnd, i, lines, linesCount)
					if isValid {
						result += num
					}
					numStart = -1
					numString = ""
				}
			} else {
				if len(numString) > 0 {
					numEnd := j
					num, _ := strconv.Atoi(numString)
					isValid := CheckIfPartNumber(numStart, numEnd, i, lines, linesCount)
					if isValid {
						result += num
					}
					numStart = -1
					numString = ""
				}
			}
		}
	}
	fmt.Printf("Result: %d", result)
}

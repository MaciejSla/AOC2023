package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: Day1 <input file>")
		os.Exit(1)
	}
	input := os.Args[1]
	file, _ := os.Open(input)
	scanner := bufio.NewScanner(file)
	var result int
	for scanner.Scan() {
		line := scanner.Text()
		var left, right = "", ""
		for _, i := range line {
			if unicode.IsNumber(i) {
				if left == "" {
					left = string(i)
				}
				right = string(i)
			}
		}
		final := left + right
		num, _ := strconv.Atoi(final)
		result += num
	}
	fmt.Printf("Final result: %d\n", result)
}

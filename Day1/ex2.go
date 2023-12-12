package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Number struct {
	Value  int
	Number string
}

var numbers = []Number{
	{1, "one"},
	{2, "two"},
	{3, "three"},
	{4, "four"},
	{5, "five"},
	{6, "six"},
	{7, "seven"},
	{8, "eight"},
	{9, "nine"},
}

func findNumber(line string) (num int) {
	var left, right = -1, -1
	var str string
	for i := 0; i < len(line); i++ {
		if unicode.IsLetter(rune(line[i])) {
			str += string(line[i])
			for _, number := range numbers {
				if strings.Contains(str, number.Number) {
					if left == -1 {
						left = number.Value
					}
					right = number.Value
					str = ""
					i -= len(number.Number) - 1
					break
				}
			}
		} else {
			if left == -1 {
				left, _ = strconv.Atoi(string(line[i]))
			}
			right, _ = strconv.Atoi(string(line[i]))
		}
	}
	num = left*10 + right
	return
}

func Ex2(input string) {
	file, _ := os.Open(input)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var result int
	for scanner.Scan() {
		line := scanner.Text()
		result += findNumber(line)
	}
	fmt.Printf("Final result: %d\n", result)
}

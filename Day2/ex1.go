package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var colorMap = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func CheckCubes(cubes []string) bool {
	for _, cube := range cubes {
		color := strings.Split(cube, " ")[2]
		value, _ := strconv.Atoi(strings.Split(cube, " ")[1])
		if colorMap[color] < value {
			return false
		}
	}
	return true
}

func Ex1(input string) {
	file, _ := os.Open(input)
	scanner := bufio.NewScanner(file)
	var result int
	for scanner.Scan() {
		line := scanner.Text()
		game := strings.Split(line, ":")
		gameNum, _ := strconv.Atoi(strings.Split(game[0], " ")[1])
		rounds := strings.Split(game[1], ";")
		isValid := true
		for _, round := range rounds {
			cubes := strings.Split(round, ",")
			if !CheckCubes(cubes) {
				isValid = false
				break
			}
		}
		if isValid {
			result += gameNum
		}
	}
	fmt.Println(result)

}

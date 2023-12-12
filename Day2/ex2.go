package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countMaxCubesInRound(rounds []string) int {
	maxCubeMap := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, round := range rounds {
		cubes := strings.Split(round, ",")
		for _, cube := range cubes {
			color := strings.Split(cube, " ")[2]
			value, _ := strconv.Atoi(strings.Split(cube, " ")[1])
			if maxCubeMap[color] < value {
				maxCubeMap[color] = value
			}
		}
	}
	return maxCubeMap["red"] * maxCubeMap["green"] * maxCubeMap["blue"]
}

func Ex2(input string) {
	file, _ := os.Open(input)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var result int
	for scanner.Scan() {
		line := scanner.Text()
		game := strings.Split(line, ":")
		rounds := strings.Split(game[1], ";")
		result += countMaxCubesInRound(rounds)
	}
	fmt.Println(result)

}

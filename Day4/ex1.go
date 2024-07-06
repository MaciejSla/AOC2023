package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Ex1(input string) {
	var globalScore int
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		var winningNums []string
		var yourNums []string
		currentScore := 0
		if err != nil {
			fmt.Println(err)
			break
		}
		card := strings.Split(line, ":")[1]
		lists := strings.Split(card, "|")
		winningNums = strings.Fields(lists[0])
		yourNums = strings.Fields(lists[1])
		for _, w := range winningNums {
			for _, y := range yourNums {
				if w == y {
					if currentScore == 0 {
						currentScore++
					} else {
						currentScore *= 2
					}
					break
				}
			}
		}
		globalScore += currentScore
	}
	fmt.Printf("Score: %d\n", globalScore)
}

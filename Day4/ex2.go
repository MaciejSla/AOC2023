package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func getCardCount(lines []string, id int) (count int) {
	var won int
	if id >= len(lines) {
		return 0
	}
	lists := strings.Split(lines[id], "|")
	winning := strings.Fields(lists[0])
	yours := strings.Fields(lists[1])
	for _, w := range winning {
		for _, y := range yours {
			if w == y {
				won++
				break
			}
		}
	}
	if won == 0 {
		return 1
	}
	for x := 1; x <= won; x++ {
		count += getCardCount(lines, id+x)
	}
	return count + 1
}

func Ex2(input string) {
	file, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(file), "\n")
	var count int
	for x := range lines {
		count += getCardCount(lines, x)
	}
	fmt.Printf("Count: %d cards\n", count)
}

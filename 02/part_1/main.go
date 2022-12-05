package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func sum(array []int) int {
	result := 0

	for _, num := range array {
		result += num
	}

	return result
}

func calculateRound(shape1 string, shape2 string) int {
	if shape1 == "A" && shape2 == "X" {
		return 1 + 3
	}

	log.Fatalf("invalid input %v and %v", shape1, shape2)
	return 0
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	m := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	scanner := bufio.NewScanner(file)

	result := 0

	for scanner.Scan() {
		line := scanner.Text()

		s := strings.Split(line, " ")

		theirPick := m[s[0]]
		myPick := m[s[1]]

		if theirPick == myPick {
			result += 3 + myPick
			continue
		}

		if myPick == 1 && theirPick == 3 {
			result += 6 + myPick
			continue
		}

		if myPick == 2 && theirPick == 1 {
			result += 6 + myPick
			continue
		}

		if myPick == 3 && theirPick == 2 {
			result += 6 + myPick
			continue
		}

		result += 0 + myPick
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

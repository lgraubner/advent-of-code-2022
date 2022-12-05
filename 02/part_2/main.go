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
	}

	scanner := bufio.NewScanner(file)

	result := 0

	for scanner.Scan() {
		line := scanner.Text()

		s := strings.Split(line, " ")

		theirPick := s[0]
		myPick := s[1]

		if theirPick == "A" && myPick == "Z" {
			result += 6 + m["B"]
			continue;
		}

		if theirPick == "B" && myPick == "Z" {
			result += 6 + m["C"]
			continue;
		}

		if theirPick == "C" && myPick == "Z" {
			result += 6 + m["A"]
			continue;
		}

		if myPick == "Y" {
			result += 3 + m[theirPick]
			continue;
		}

		if theirPick == "A" && myPick == "X" {
			result += 0 + m["C"]
			continue;
		}

		if theirPick == "B" && myPick == "X" {
			result += 0 + m["A"]
			continue;
		}

		if theirPick == "C" && myPick == "X" {
			result += 0 + m["B"]
			continue;
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

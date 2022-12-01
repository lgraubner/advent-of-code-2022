package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func sum(array []int) int {
	result := 0

	for _, num := range array {
		result += num
	}

	return result
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	elfExpedition := make([]int, 0)
	elfCalories := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			elfExpedition = append(elfExpedition, elfCalories)

			elfCalories = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("can not convert string to int: %v", err)
		}

		elfCalories += calories
	}

	elfExpedition = append(elfExpedition, elfCalories)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(elfExpedition)

	fmt.Println(sum(elfExpedition[len(elfExpedition)-3:]))
}

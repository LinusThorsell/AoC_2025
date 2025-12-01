package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func first() int {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// State
	var position = 50
	var count = 0

	// File reading loop
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var currentLine = scanner.Text()

		// Parse line into prefix (L/R) and rotation amount as int
		var prefix = currentLine[:1]
		var rotation, err = strconv.Atoi(currentLine[1:])
		if err != nil {
			log.Fatal(err)
		}

		if prefix == "L" {
			// Rotate left
			position -= rotation
		} else {
			// Rotate right
			position += rotation
		}

		position = position % 100
		if position == 0 {
			count += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return count
}

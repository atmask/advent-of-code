package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/atmask/advent-of-code/internal/dial"
)

func main() {
	fmt.Println("Hello North Pole!")

	inputFile := flag.String("input", "input.txt", "path to input data")
	flag.Parse()

	// Givens
	startingPosition := 50
	numPositions := 100
	keyValue := 0

	dial := dial.NewDial(startingPosition, numPositions, keyValue)

	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		if line == "" {
			break
		}

		direction := line[0] // 'L' or 'R' byte
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Rotating %d steps %c\n", steps, direction)

		switch direction {
		case 'L':
			dial.RotateLeft(steps)
		case 'R':
			dial.RotateRight(steps)
		}
	}
	fmt.Println("Finished file")
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Final Answer
	fmt.Printf("The key values %d appeared %d times", keyValue, dial.KeyValueSeen)

}

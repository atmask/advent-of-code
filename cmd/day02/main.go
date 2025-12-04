package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"strings"
)

func main() {
	/*
				so...
				for two digit numbers
		      find all permutations (x_1, x_2) within a range
				for three digit numbers
					not possible

				for fours digit numbers where (x_1,y_1) = (x_2,y_2)

				Conditions:
		    - number must have even number of digits
		    - split at halfway point
		    - x|y : x = y
	*/
	inputFile := flag.String("input", "input.txt", "path to input data")
	flag.Parse()

	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	for {
		range, err := reader.ReadString(',')

		if range != "" {
			fmt.Printf("Current Range: %s\n", range)
		}
		
		// Clean up string
		range = strings.TrimSpace(strings.TrimSuffix(range, ","))

		strlen := len(range)
		if (strlen % 2) != 0 {
			fmt.Println("This cannot be a mirrored range")
		}

		half := strlen / 2
		first := range[:half]
		last := range[half:]
		fmt.Printf("First: %s, Last: %s\n", first, last)

		// Smarter way would be to generate the numbers in the range rather than scanning
		
		// EOF
		if err != nil {
			log.Fatal(err)
		}
	}
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"strconv"
	"strings"
)

func rangeSum(min int, max int) int {
	// Return the sum of values between max and min
	return ((max - min + 1) * (max + min)) / 2
}

func getIterableMin(rangeMin string) int {
	/*
		if AB < CD:
		MIN is (AB+1)(AB+1)
		elif AB >= CD:
		MIN is ABAB
	*/

	// Handle odd-digit numbers: jump to next even digit count
	if len(rangeMin)%2 == 1 {
		halfDigits := (len(rangeMin) + 1) / 2
		if halfDigits == 1 {
			return 1 // smallest half for 2-digit doubled (11)
		}
		// Return 10^(halfDigits-1), e.g., for 3 digits → return 10 (for 1010)
		min := 1
		for i := 0; i < halfDigits-1; i++ {
			min *= 10
		}
		return min
	}

	firstHalf, _ := strconv.Atoi(rangeMin[:len(rangeMin)/2])
	secondHalf, _ := strconv.Atoi(rangeMin[len(rangeMin)/2:])

	if firstHalf < secondHalf {
		return firstHalf + 1
	}
	return firstHalf
}

func getIterableMax(rangeMax string) int {
	// if EF =< GH:
	//   MAX is EFEF
	// elif EF > GH:
	//   MAX is (EF-1)(EF-1)

	// Handle odd-digit numbers: use previous even digit count's max
	if len(rangeMax)%2 == 1 {
		halfDigits := (len(rangeMax) - 1) / 2
		if halfDigits == 0 {
			return 0 // No valid doubled numbers below single-digit numbers
		}
		// Return 10^halfDigits - 1, e.g., for 3 digits → return 9 (for 99)
		max := 1
		for i := 0; i < halfDigits; i++ {
			max *= 10
		}
		return max - 1
	}

	firstHalf, _ := strconv.Atoi(rangeMax[:len(rangeMax)/2])
	secondHalf, _ := strconv.Atoi(rangeMax[len(rangeMax)/2:])

	if firstHalf <= secondHalf {
		return firstHalf
	}
	return firstHalf - 1
}

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
	nums := make([]int, 0)
	sum := 0
	for {
		idRange, err := reader.ReadString(',')

		if idRange != "" {
			fmt.Printf("Current Range: %s\n", idRange)
		}

		// Clean up string
		idRange = strings.TrimSpace(strings.TrimSuffix(idRange, ","))

		// Get range min and max
		parts := strings.Split(idRange, "-")
		rangeMin := parts[0]
		rangeMax := parts[1]

		iterMin := getIterableMin(rangeMin)
		iterMax := getIterableMax(rangeMax)

		for i := iterMin; i <= iterMax; i++ {
			doubled := fmt.Sprintf("%d%d", i, i)
			val, _ := strconv.Atoi(doubled)
			nums = append(nums, val)
			sum += val
		}

		// EOF
		if err != nil {
			break
		}
	}
	fmt.Printf("Sum of nums: %d\nThere were %d bad ids", sum, len(nums))

}

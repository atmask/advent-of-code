package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"strings"
)

func rangeSum(min int, max int){
	// Return the sum of values between max and min 
	return ((max - min +1) * (max + min)) / 2
}

func getIterableMin(rangeMin int) int {
	/*
  if AB < CD:
    MIN is (AB-1)(AB-1)
  elif AB >= CD:
    MIN is ABAB  
  */ 
	return -1
}

func getIterableMax(rangeMax int) int {
	return -1
}

func getSolvableRanges(min int, max int) {
  return
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
	for {
		range, err := reader.ReadString(',')

		if range != "" {
			fmt.Printf("Current Range: %s\n", range)
		}
		
		// Clean up string
		range = strings.TrimSpace(strings.TrimSuffix(range, ","))
		
		// Get range min and max
		parts := strings.Split(s, "-")
		rangeMin := parts[0]
		rangeMax := parts[1]
		solvableRanges = []int
		if len(rangeMin) != len(rangeMax) {
			getSolvableRanges(rangeMin, rangeMax)
		} else {
			solvableRanges := [rangeMin, rangeMax]
	  }
		
		

		// 



		strlen := len(range)
		if (strlen % 2) != 0 {
			fmt.Println("This cannot be a mirrored range")
		}

		half := strlen / 2
		first := range[:half]
		last := range[half:]
		fmt.Printf("First: %s, Last: %s\n", first, last)

		// Smarter way would be to generate the numbers in the range rather than scanning
		/*
			double range

			For 1234 - 6789

			you generate all number x_1,y_1 that are greater than 1234 and lees than 6789.
			if x_1, y_1 == 1 2 (i.e. range min) then start at min 34
				12
				12
			
			if x_1, y_1 == 67 then stop at 89

		*/

		// For a range with fixed length 12345678. This would be all 1234 that are greater than min and less than max
		
		// EOF
		if err != nil {
			log.Fatal(err)
		}
	}
}


/*
ex. 1234 - 5545

two digits


## Start
3412
1234 determines the min by it's first two digits. They must be larger than the second two though to be in range

if AB < CD:
  MIN is (AB-1)(AB-1)
elif AB >= CD:
  MIN is ABAB  


12 12 (Bad)
13 (good)

## End 
5545 determines the max it's first two but they must be less than 45

This means that 
if EF =< GH:
  MAX is EFEF
elif EF > GH:
  MAX is (EF-1)(EF-1)
  

55


ex. 3456 - 5412



ex. 3412 - 5000

34
35 
36
37
..
49


ex. 3415 - 12345

34


ex 3415 - 123456


*/

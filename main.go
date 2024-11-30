package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"math-skills/funcs"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Input command correctly!")
		return
	}

	dataFile := os.Args[1]

	data, err := os.Open(dataFile)
	if err != nil {
		fmt.Println("Data File missing!")
		return
	}

	ints := []int{}

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		intg, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error in data file, contains elements other than numbers!")
			return
		}
		ints = append(ints, intg)
	}

	if len(ints) == 0 {
		fmt.Println("Missing data!")
		return
	}

	fmt.Printf("Average: %.0f\n", math.Round(funcs.Average(ints)))
	fmt.Printf("Median: %.0f\n", math.Round(funcs.Median(ints)))
	fmt.Printf("Variance: %.0f\n", math.Round(funcs.Variance(ints)))
	fmt.Printf("Standard Deviation: %.0f\n", math.Round(funcs.StandardDeviation(ints)))
}

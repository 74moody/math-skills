package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sort"
	"math"
)

func main() {
	var validInt []int
	log.SetFlags(0) // to remove date and time in error message

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run your-program.go data.txt")
		os.Exit(1)
	}

	dataFile := os.Args[1]
	data, err := os.ReadFile(dataFile)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	linesWithIntegers := 0

	for _, line := range lines {
		// Trim leading and trailing whitespace
		line = strings.TrimSpace(line)

		// Skip empty lines
		if line == "" {
			continue
		}

		// Convert the line to an integer
		validatedInt, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("[-] Sorry, invalid data. Usage example:\n123\n123\n345")
		}

		// Add the integer to the slice
		validInt = append(validInt, validatedInt)

		// Increment the count of lines with integers
		linesWithIntegers++
	}

	// Check if the minimum number of lines with integers is met
	if linesWithIntegers < 3 {
		log.Fatalf("[-] Error: Expected a minimum of three lines with integers")
	}

	// Calling function
	average := calculate_Mean(validInt)
	fmt.Println("Average:", format_decimals(average))
	median := calculate_median(validInt)
	fmt.Println("Median:", format_decimals(median))
	variance := calculate_variance(validInt)
	fmt.Println("Variance:", format_decimals(variance))
	standard_deviation := calculate_standard_deviation(validInt)
	fmt.Println("Standard Deviation:", format_decimals(standard_deviation))
}

// Calculate the mean
func calculate_Mean(validInt []int) float64 {
	sum := 0
	for i := 0; i < len(validInt); i++ {
		sum = sum + validInt[i]
	}

	average := float64(sum) / float64(len(validInt))
	return average
}

func calculate_median(validInt []int) float64 {
	sort.Ints(validInt)

	integer_value := len(validInt)
	if integer_value%2==1 {
		get_odd_int := integer_value/2
		return float64(validInt[get_odd_int])
	} else {
		// where the median is even we will take both evens
		get_first_even := integer_value/2-1
		get_second_even := integer_value/2
		return float64(validInt[get_first_even]+validInt[get_second_even]) / 2.0
	}
}

func calculate_variance(validInt []int) float64 {
	average := calculate_Mean(validInt)
	squared_sum_integers := 0.0
	for _, integers := range validInt {
		differece := float64(integers) - average
		square_difference := differece*differece
		squared_sum_integers += square_difference
	}
	variance_result := squared_sum_integers/ float64(len(validInt))
	return variance_result
}

func calculate_standard_deviation(validInt []int) float64 {
	variance := calculate_variance(validInt)
	standard_deviation := math.Sqrt(variance)
	return standard_deviation
}

func format_decimals(number float64) string {
	rounded_integer := int(math.Round(number))
	return fmt.Sprintf("%d", rounded_integer)
}


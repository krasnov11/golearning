package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var numbers []float64
	arguments := os.Args[1:]
	for _, argument := range arguments {
		value, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, value)
	}

	fmt.Printf("Average: %0.2f\n", average(numbers...))
}

func average(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0.0
	}

	sum := nums[0]

	for _, value := range nums[1:] {
		sum += value
	}

	return sum / float64(len(nums))
}

package main

import (
	"fmt"
	"golearning/datafile"
	"log"
)

func main() {
	path := "data.txt"
	nums, err := datafile.GetFloats(path)
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, number := range nums {
		//fmt.Printf("[%d] %f\n", i, v)
		sum += number
	}
	sampleCount := float64(len(nums))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}

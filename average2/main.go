package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var sum float64
	arguments := os.Args[1:]
	for _, argument := range arguments {
		value, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}

		sum += value
	}

	fmt.Printf("Average: %0.2f\n", sum/float64(len(arguments)))
}

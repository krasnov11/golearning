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

	for i, v := range nums {
		fmt.Printf("[%d] %f\n", i, v)
	}
}

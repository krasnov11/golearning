package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	path := "data.txt"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"fmt"
	"golearning/datafile"
	"log"
	"sort"
)

func main() {
	main1()

	fmt.Println("-------------------------------")

	main2()
}

func main2() {
	path := "votes.txt"
	lines, err := datafile.GetStrings(path)
	if err != nil {
		log.Fatal(err)
	}

	counters := map[string]int{}

	for _, line := range lines {
		counters[line]++
	}

	var keys []string
	for key := range counters {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%s: %d\n", key, counters[key])
	}

	//fmt.Println(counters)
}

func main1() {
	path := "votes.txt"
	lines, err := datafile.GetStrings(path)
	if err != nil {
		log.Fatal(err)
	}

	var names []string
	var counts []int

	for _, line := range lines {
		notFound := true

		for i, name := range names {
			if line == name {
				counts[i]++
				notFound = false
				break
			}
		}

		if notFound {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}

	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}

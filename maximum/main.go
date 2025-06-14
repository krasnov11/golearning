package main

import (
	"fmt"
)

func main() {

	var min, max float64

	/*arguments := os.Args[1:]
	if len(arguments) == 0 {
		log.Fatal("Please give one or more floats.")
	}


	for i, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			continue
		}
		if i == 0 {
			min = number
			max = number
			continue
		}
		if number < min {
			min = number
		}
		if number > max {
			max = number
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)*/

	max, min = maxmin(1, 3, 4, -29.11, 4.22)
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

	max, min = maxmin(1)
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

	max, min = maxmin()
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

	fmt.Println("Max:", maximum(-1, 2, 3, -4, 5, 6, 7, -8, 9, 10))
	fmt.Println("Max:", minimum(1, -2, 3, 4, 5, -6, 7, 8, -9, 10))

	fmt.Println("inRange", inRange(1, 100, -12.5, 3.2, 0, 50, 103.5))
	fmt.Println("inRange", inRange(-10, 10, 4.1, 12, -12, -5.2))

	fmt.Println("average", average(-10, 10, 4.1, 12, -12, -5.2))
}

func maximum(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0.0
	}

	max := nums[0]

	for _, value := range nums[1:] {
		if value > max {
			max = value
		}
	}

	return max
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

func minimum(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0.0
	}

	min := nums[0]

	for _, value := range nums[1:] {
		if value < min {
			min = value
		}
	}

	return min
}

func maxmin(nums ...float64) (float64, float64) {
	if len(nums) == 0 {
		return 0.0, 0.0
	}

	max := nums[0]
	min := max

	for _, value := range nums[1:] {
		if value > max {
			max = value
		}

		if value < min {
			min = value
		}
	}

	return max, min
}

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

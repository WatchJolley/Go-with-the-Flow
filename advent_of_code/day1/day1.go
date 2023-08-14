package main

import (
	"bufio"
	"fmt"
	_ "io"
	"log"
	"os"
	"strconv"
)

// TODO: Cant say I love the use of this function..
func addElfCaloriesTotal(caloriesSlice *[]int, caloriesSum *int, largestCalories *int) {
	// TODO: Surely, there is a better way than to deference everything?
	*caloriesSlice = append(*caloriesSlice, *caloriesSum)
	if *caloriesSum > *largestCalories {
		*largestCalories = *caloriesSum
	}
	*caloriesSum = 0
}

func main() {
	dat, err := os.Open("./day_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer dat.Close()

	var caloriesSlice []int
	caloriesSum := 0
	largestCalories := 0

	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		value_as_string := scanner.Text()
		value_as_int, err := strconv.Atoi(value_as_string)

		caloriesSum += value_as_int
		// The assumption here is when you hit an error you have hit a space
		// it may not be correct for all file formats but seems to fit with
		// the data that has been given.
		if err == nil {
			continue
		}
		addElfCaloriesTotal(&caloriesSlice, &caloriesSum, &largestCalories)
	}

	addElfCaloriesTotal(&caloriesSlice, &caloriesSum, &largestCalories)
	fmt.Printf("Largest total sum of calories found was %v", largestCalories)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mostCalories(filePath string) int {

	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	maximumCalories := 0
	secondMostCalories := 0
	thirdMostCalories := 0
	currentCallorieSum := 0

	for _, line := range fileLines {
		if line == "" {
			if currentCallorieSum > maximumCalories {
				thirdMostCalories = secondMostCalories
				secondMostCalories = maximumCalories
				maximumCalories = currentCallorieSum
			} else if currentCallorieSum > secondMostCalories {
				thirdMostCalories = secondMostCalories
				secondMostCalories = currentCallorieSum
			} else if currentCallorieSum > thirdMostCalories {
				thirdMostCalories = currentCallorieSum
			}

			fmt.Printf("First: %v \n", maximumCalories)
			fmt.Printf("Seconds: %v \n", secondMostCalories)
			fmt.Printf("Third: %v \n", thirdMostCalories)
			currentCallorieSum = 0
		} else {
			var calories int
			calories, err = strconv.Atoi(line)

			if err != nil {
				fmt.Print(err)
			}
			currentCallorieSum += calories
		}
	}

	return maximumCalories + secondMostCalories + thirdMostCalories
}

func main() {
	fmt.Print(mostCalories("input.txt"))
}

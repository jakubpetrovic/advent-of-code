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
	currentCallorieSum := 0

	for _, line := range fileLines {
		if line == "" {
			if currentCallorieSum > maximumCalories {
				maximumCalories = currentCallorieSum
			}
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

	return maximumCalories
}

func main() {
	fmt.Print(mostCalories("input.txt"))
}

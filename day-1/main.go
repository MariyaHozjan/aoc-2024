package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func createColumns(path string) ([]int, []int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var column1 []int
	var column2 []int

	//var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		splt := splitString(scanner.Text())

		string1, err := strconv.Atoi(splt[0])
		if err != nil {
			panic(err)
		}
		column1 = append(column1, string1)

		string2, err := strconv.Atoi(splt[1])
		if err != nil {
			panic(err)
		}
		column2 = append(column2, string2)
	}

	return column1, column2, scanner.Err()
}

func calculateDifference(column1 []int, column2 []int) int {
	sort.Ints(column2)
	sort.Ints(column1)

	Difference := 0

	for i := 0; i < len(column1); i++ {
		if column1[i] < column2[i] {
			Difference += column2[i] - column1[i]
			continue
		}
		Difference += column1[i] - column2[i]
	}

	print(Difference)
	return Difference
}

func splitString(s string) []string {

	words := strings.Fields(s)
	return words
}

func similarityScore(column1 []int, column2 []int) int {
	similarity := 0

	for i := 0; i < len(column1); i++ {
		similarityOfCurrentNumber := 0
		for j := 0; j < len(column2); j++ {
			if column1[i] == column2[j] {
				similarityOfCurrentNumber += column1[i]
			}
		}
		similarity += similarityOfCurrentNumber
	}
	print(similarity)
	return similarity
}

func main() {
	column1, column2, err := createColumns("input.txt")
	if err != nil {
		panic(err)
	}

	print("Difference between columns: ")
	calculateDifference(column1, column2)
	print("\n")
	print("Similarity score: ")
	similarityScore(column1, column2)
}

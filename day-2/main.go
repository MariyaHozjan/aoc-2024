package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func createArrayOfArrays(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var arrayOfArrays [][]int
	for scanner.Scan() {
		var lines []int
		splt := splitString(scanner.Text())

		for i := 0; i < len(splt); i++ {
			string, err := strconv.Atoi(splt[i])
			if err != nil {
				panic(err)
			}
			lines = append(lines, string)
		}
		arrayOfArrays = append(arrayOfArrays, lines)
	}
	return arrayOfArrays, nil
}

func splitString(s string) []string {

	words := strings.Fields(s)
	return words
}

func numberOfSafeReports(arrayOfArrays [][]int) int {
	var numberOfSafes int
	for i := 0; i < len(arrayOfArrays); i++ {
		safe := true
		if arrayOfArrays[i][1] < arrayOfArrays[i][0] {
			for j := 1; j < len(arrayOfArrays[i]); j++ {
				if arrayOfArrays[i][j] > arrayOfArrays[i][j-1] || 1 > arrayOfArrays[i][j-1]-arrayOfArrays[i][j] || 3 < arrayOfArrays[i][j-1]-arrayOfArrays[i][j] {
					safe = false
					break
				}
			}
		}
		if arrayOfArrays[i][1] > arrayOfArrays[i][0] {
			for j := 1; j < len(arrayOfArrays[i]); j++ {
				if arrayOfArrays[i][j] < arrayOfArrays[i][j-1] || 1 > arrayOfArrays[i][j]-arrayOfArrays[i][j-1] || 3 < arrayOfArrays[i][j]-arrayOfArrays[i][j-1] {
					safe = false
					break
				}
			}
		}
		if arrayOfArrays[i][1] == arrayOfArrays[i][0] {
			safe = false
		}
		if safe {
			numberOfSafes++
		}
	}
	return numberOfSafes
}

func main() {
	array, err := createArrayOfArrays("input.txt")
	if err != nil {
	}
	print(numberOfSafeReports(array))
}

/*
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
*/

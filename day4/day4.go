package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFile, err := os.Open("day4/day4.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inputFile.Close()

	var grid []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			grid = append(grid, line)
		}
	}

	count := countXMAS(grid)
	fmt.Println("count is", count)
}

func readInputFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			grid = append(grid, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func countXMAS(grid []string) int {
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	// searching horizontally (forward and backward)
	for _, row := range grid {
		count += strings.Count(row, "XMAS")
		count += strings.Count(row, "SAMX")
	}

	// vertical (down and up)
	for col := 0; col < cols; col++ {
		vertical := ""
		for row := 0; row < rows; row++ {
			vertical += string(grid[row][col])
		}
		count += strings.Count(vertical, "XMAS")
		count += strings.Count(vertical, "SAMX")
	}

	// diagonal (right to left)
	for startRow := 0; startRow < rows; startRow++ {
		diagonal := ""
		for i := 0; startRow+i < rows && i < cols; i++ {
			diagonal += string(grid[startRow+i][i])
		}
		count += strings.Count(diagonal, "XMAS")
		count += strings.Count(diagonal, "SAMX")
	}
	for startCol := 1; startCol < cols; startCol++ {
		diagonal := ""
		for i := 0; i < rows && startCol+i < cols; i++ {
			diagonal += string(grid[i][startCol+i])
		}
		count += strings.Count(diagonal, "XMAS")
		count += strings.Count(diagonal, "SAMX")
	}

	// diagonal (left to right)
	for startRow := 0; startRow < rows; startRow++ {
		diagonal := ""
		for i := 0; startRow+i < rows && cols-1-i >= 0; i++ {
			diagonal += string(grid[startRow+i][cols-1-i])
		}
		count += strings.Count(diagonal, "XMAS")
		count += strings.Count(diagonal, "SAMX")
	}
	for startCol := 0; startCol < cols-1; startCol++ {
		diagonal := ""
		for i := 0; i < rows && startCol-i >= 0; i++ {
			diagonal += string(grid[i][startCol-i])
		}
		count += strings.Count(diagonal, "XMAS")
		count += strings.Count(diagonal, "SAMX")
	}

	return count
}

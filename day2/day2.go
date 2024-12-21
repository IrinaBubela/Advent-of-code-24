package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func intAbs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func isSafe(level []int) bool {
	allDecreasing, allIncreasing := true, true
	prev := level[0]
	for _, el := range level[1:] {
		if el >= prev {
			allDecreasing = false
		}

		if el <= prev {
			allIncreasing = false
		}

		if !allDecreasing && !allIncreasing {
			return false
		}

		if intAbs(prev-el) > 3 {
			return false
		}

		prev = el
	}
	return true
}

func checkLevel(level []int) bool {
	if isSafe(level) {
		return true
	}

	for i := 0; i < len(level); i++ {
		newLevel := []int{}

		//collect elements before i
		newLevel = append(newLevel, level[:i]...)
		//collect elements after i
		newLevel = append(newLevel, level[i+1:]...)
		if isSafe(newLevel) {
			return true
		}
	}
	return false
}

func silver() {
	file, err := os.Open("day2/day2.txt")
	if err != nil {
		//log err
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var levels [][]int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		numbers := make([]int, len(parts))

		//converting str to int
		for i, part := range parts {
			num, _ := strconv.Atoi(part)
			numbers[i] = num
		}
		levels = append(levels, numbers)
	}

	count := 0
	for _, level := range levels {
		if isSafe(level) {
			count++
		}
	}
	fmt.Println(count)
}

func gold() {
	file, err := os.Open("day2/day2.txt")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var levels [][]int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		numbers := make([]int, len(parts))

		//converting str to int
		for i, part := range parts {
			num, _ := strconv.Atoi(part)
			numbers[i] = num
		}
		levels = append(levels, numbers)
	}

	count := 0
	for _, level := range levels {
		if checkLevel(level) {
			count++
		}
	}

	fmt.Println(count)
}

func main() {
	silver()
	gold()
}

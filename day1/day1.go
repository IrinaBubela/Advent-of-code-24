package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("day1/day1-input.txt")

	if err != nil {
		fmt.Println("Error with input")
		return
	}
	var left []int
	var right []int

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	for _, line := range lines {
		numbers := strings.Fields(line)
		if len(numbers) == 2 {
			leftNum, _ := strconv.Atoi(numbers[0])
			rightNum, _ := strconv.Atoi(numbers[1])
			left = append(left, leftNum)
			right = append(right, rightNum)
		}
	}

	counter := map[int]int{}

	for _, el := range right {
		counter[el]++
	}
	fmt.Println(counter)

	sum := 0

	for _, elem := range left {
		n, ok := counter[elem]
		if ok {
			sum += elem * n
		}
	}

	fmt.Println(sum)

}

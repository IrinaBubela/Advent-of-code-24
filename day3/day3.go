package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func sumupMul(match []string) int {
	x, err1 := strconv.Atoi(match[1])
	y, err2 := strconv.Atoi(match[2])
	if err1 != nil || err2 != nil {
		fmt.Println("Error converting values to integers:", err1, err2)
		return 0
	}
	return x * y
}

func first() {
	inputFile, err := os.Open("day3/day3.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var text string
	for scanner.Scan() {
		text += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	regexMul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	regexDo := regexp.MustCompile(`do\(\)`)
	regexDont := regexp.MustCompile(`don't\(\)`)

	mulEnabled := true
	var sum int

	i := 0
	for i < len(text) {
		if regexDo.MatchString(text[i:]) {
			mulEnabled = true
			i += len("do()")
		} else if regexDont.MatchString(text[i:]) {
			mulEnabled = false
			i += len("don't()")
		} else if mulEnabled && regexMul.MatchString(text[i:]) {
			matches := regexMul.FindStringSubmatch(text[i:])
			if len(matches) > 0 {
				sum += sumupMul(matches)
			}
			i += len(matches[0])
		} else {
			i++
		}
	}

	fmt.Println(sum)
}

func main() {
	solutionWithCaputureGroup()

}

func solutionWithCaputureGroup() {

	inputFile, err := os.Open("day3/day3.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var text string
	for scanner.Scan() {
		text += scanner.Text() + "\n"
	}

	regexMul := regexp.MustCompile(`mul\((\d+),(\d+)\)|(do\(\))|(don't\(\))`)

	matches := regexMul.FindAllStringSubmatch(text, -1)
	sum := 0
	enabled := true
	for _, match := range matches {
		if match[1] != "" && match[2] != "" && enabled {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			sum += x * y
		} else if match[3] == "do()" {
			enabled = true
		} else if match[4] == "don't()" {
			enabled = false
		}
	}

	fmt.Println(sum)
}

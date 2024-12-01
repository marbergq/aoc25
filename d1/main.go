package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	// part1()
	part2()
}

func part1() {
	lines := strings.Split(input, "\n")
	leftList := make([]int, 0, len(lines))
	rightList := make([]int, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[3])

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	sum := 0
	for i := 0; i < len(lines); i++ {
		diff := int(math.Abs(float64(leftList[i]) - float64(rightList[i])))
		// fmt.Print(leftList[i], " ", rightList[i], " ")
		// fmt.Print("+", diff)
		sum += diff
	}

	fmt.Println(sum)
}

func part2() {
	lines := strings.Split(input, "\n")
	leftList := make([]int, 0, len(lines))
	rightCount := map[int]int{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[3])
		leftList = append(leftList, left)
		rightCount[right]++
	}

	sum := 0
	for _, left := range leftList {
		rightCount, ok := rightCount[left]
		if !ok {
			rightCount = 0
		}
		sum += left * rightCount
	}

	fmt.Println(sum)
}

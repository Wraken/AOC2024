package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFile() []string {
	r, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	scan := bufio.NewScanner(r)

	scan.Split(bufio.ScanLines)

	lines := []string{}
	for scan.Scan() {
		l := scan.Text()
		lines = append(lines, l)
	}

	return lines
}

func parseLines(lines []string) ([]int, []int) {
	l1, l2 := []int{}, []int{}

	for _, l := range lines {
		values := strings.Split(strings.ReplaceAll(l, "   ", " "), " ")
		i, _ := strconv.Atoi(values[0])
		l1 = append(l1, i)

		i, _ = strconv.Atoi(values[1])
		l2 = append(l2, i)
	}
	return l1, l2
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func part1(lines []string) {
	l1, l2 := parseLines(lines)

	slices.SortFunc(l1, func(i, j int) int {
		return i - j
	})

	slices.SortFunc(l2, func(i, j int) int {
		return i - j
	})

	total := 0

	for i, v := range l1 {
		dist := abs(v - l2[i])
		total += dist
	}
	fmt.Println(total)
}

func containsN(i int, slice []int) int {
	n := 0
	for _, s := range slice {
		if s == i {
			n++
		}
	}
	return n
}

func part2(lines []string) {
	l1, l2 := parseLines(lines)

	total := 0
	for _, v := range l1 {
		n := containsN(v, l2)
		total += v * n
	}

	fmt.Println(total)
}

func main() {
	lines := readFile()
	part1(lines)
	part2(lines)
}

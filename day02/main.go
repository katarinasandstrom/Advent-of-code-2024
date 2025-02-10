package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var slice1 []string
var slice3 [][]int
var totNrOfValid = 0

func main() {
	readFile()
	calculate()

	fmt.Printf("Antalet giltiga rapporter är %d\n", totNrOfValid)
}

func readFile() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Fel vid öppning av fil:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		slice1 = append(slice1, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Fel vid läsning av fil:", err)
	}
}

func calculate() {
	for _, line := range slice1 {
		words := strings.Split(line, " ")
		var slice2 []int

		for _, word := range words {
			num, err := strconv.Atoi(word)
			if err != nil {
				fmt.Println("Fel vid konvertering:", err)
				continue
			}
			slice2 = append(slice2, num)
		}

		if len(slice2) < 2 {
			continue
		}

		if isValidReport(slice2) {
			totNrOfValid++
		} else if extendedCheck(slice2) {
			totNrOfValid++
		} else {
			slice3 = append(slice3, slice2)
		}
	}
}

func isValidReport(report []int) bool {
	incr := false
	decr := false

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		if diff > 0 {
			incr = true
		}
		if diff < 0 {
			decr = true
		}
	}

	return incr != decr
}

func extendedCheck(report []int) bool {
	for i := 0; i < len(report); i++ {

		newReport := append([]int{}, report[:i]...)
		if i+1 < len(report) {
			newReport = append(newReport, report[i+1:]...)
		}

		if isValidReport(newReport) {
			return true
		}
	}
	return false
}

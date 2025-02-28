package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var slice1 []int
var slice2 []int

func main() {
	organiseFile()
	firstTask()
	secondTask()

}

func organiseFile() {

	file, err := os.Open("numbers.txt")
	if err != nil {
		fmt.Println("Fel vid öppning av fil", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)

		if len(numbers) != 2 {
			fmt.Println("Fel på rad. Innehåller inte exakt två nummer", line)
			continue
		}

		num1, err1 := strconv.Atoi(numbers[0])
		num2, err2 := strconv.Atoi(numbers[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Fel vid konvertering av nummer:", line)
			continue
		}

		slice1 = append(slice1, num1)
		slice2 = append(slice2, num2)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Fel vid läsning av fil: ", err)
	}

	sort.Ints(slice1)
	sort.Ints(slice2)
}

func firstTask() {

	fmt.Printf("Slice 1: %d\n ", slice1)
	fmt.Printf("Slice 2: %d\n", slice2)

	totalSum := 0

	for i := range slice1 {
		var sum int
		if slice1[i] > slice2[i] {
			fmt.Printf("nr från lista 1: %d, nr från lista 2: %d\n", slice1[i], slice2[i])

			sum = slice1[i] - slice2[i]
			totalSum += sum
			fmt.Printf("Summa rad %d är %d\n", i, sum)
		} else {

			sum = slice2[i] - slice1[i]
			totalSum += sum
			fmt.Printf("Summa rad %d är %d\n", i, sum)
		}
	}

	fmt.Printf("Total skillnad: %d\n", totalSum)
}

func secondTask() {
	var currentNr int
	var totalSum int

	var partialSum int

	for i := range slice1 {
		nrOfTimes := 0
		currentNr = slice1[i]

		for j := range slice2 {
			if slice2[j] == currentNr {
				nrOfTimes++
			} else {
				continue
			}

		}
		partialSum = nrOfTimes * currentNr
		totalSum += partialSum

	}

	fmt.Printf("Totalsumman är %d\n", totalSum)
}

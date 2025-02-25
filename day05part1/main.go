package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	grid1 := readFile1("data1.txt")
	grid2 := readFile2("data2.txt")

	totalSum := testOrder(grid2, grid1)
	fmt.Println("Total Sum:", totalSum)

}

func readFile1(filename string) [][]int {
	var grid1 [][]int

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Fel vid öppning av fil", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		parts := strings.Split(scanner.Text(), "|")
		var line []int

		for _, part := range parts {
			num, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				fmt.Println("Fel vid konvertering:", err)
				return nil
			}
			line = append(line, num)
		}
		grid1 = append(grid1, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Fel vid läsning av fil:", err)
		return nil
	}

	return grid1

}

func readFile2(filename string) [][]int {
	var grid2 [][]int

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Fel vid öppning av fil", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		parts := strings.Split(scanner.Text(), ",")
		var line []int

		for _, part := range parts {
			num, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				fmt.Println("Fel vid koncertering:", err)
				return nil
			}
			line = append(line, num)
		}
		grid2 = append(grid2, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Fel vid läsning av fil:", err)
		return nil
	}

	return grid2

}

func testOrder(grid2 [][]int, grid1 [][]int) int {

	totalSum := 0

	for _, row := range grid2 {
		if isValidOrder(row, grid1) {
			mid := findMiddle(row)
			totalSum += mid
		}
	}
	return totalSum
}

func isValidOrder(row []int, grid1 [][]int) bool {
	pos := make(map[int]int)
	for i, num := range row {
		pos[num] = i
	}
	for _, rule := range grid1 {
		x, y := rule[0], rule[1]

		// Kolla om både X och Y finns i row
		xPos, xExists := pos[x]
		yPos, yExists := pos[y]

		// Om båda finns men X kommer efter Y → FEL!
		if xExists && yExists && xPos > yPos {
			return false
		}
	}

	return true // Om inga regler bryts, returnera true
}

func findMiddle(row []int) int {
	midIndex := len(row) / 2
	return row[midIndex]
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

var text string
var content []string
var partSum = 0
var totalSum = 0
var enabled = true
var instructions []Instruction

type Instruction struct {
	Type string
	Pos  int
	X, Y int
}

func main() {
	readFile()
	filterText()
	calculate()
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
		text += scanner.Text() + "\n"
	}
}

func filterText() {
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)
	mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	doMatches := doRe.FindAllStringIndex(text, -1)
	dontMatches := dontRe.FindAllStringIndex(text, -1)
	mulMatches := mulRe.FindAllStringIndex(text, -1)

	for _, match := range doMatches {
		instructions = append(instructions, Instruction{"do", match[0], 0, 0})
	}
	for _, match := range dontMatches {
		instructions = append(instructions, Instruction{"don't", match[0], 0, 0})
	}
	for _, match := range mulMatches {
		startPos := match[0]

		submatch := mulRe.FindStringSubmatch(text[startPos:match[1]])

		if len(submatch) < 3 {
			fmt.Println("Fel: förväntade minst 3 element i submatch, fick:", submatch)
			continue
		}

		x, err1 := strconv.Atoi(submatch[1])
		y, err2 := strconv.Atoi(submatch[2])

		if err1 != nil || err2 != nil {
			fmt.Println("Fel vid konvertering av siffror:", err1, err2)
			continue
		}

		instructions = append(instructions, Instruction{"mul", startPos, x, y})
	}
	sort.Slice(instructions, func(i, j int) bool {
		return instructions[i].Pos < instructions[j].Pos
	})

}

func calculate() {

	for _, i := range instructions {
		switch i.Type {
		case "do":
			enabled = true
		case "don't":
			enabled = false
		case "mul":
			if enabled == true {
				totalSum += i.X * i.Y

			}
		}
	}
	fmt.Println("Summan blir ", totalSum)
}

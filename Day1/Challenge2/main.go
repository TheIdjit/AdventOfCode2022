package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := getFile()
	ans := findMostCalTop3(input)
	fmt.Println(ans)
}

func getFile() []string {
	val := []string{}
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	read := bufio.NewReader(file)
	for {
		freqMod, _, err := read.ReadLine()
		if err != nil {
			break
		}
		l := string(freqMod)
		val = append(val, l)
	}
	return val
}

func findMostCalTop3(input []string) int {
	calCounts := calcCalories(input)
	first := 0
	second := 0
	third := 0
	for _, val := range calCounts {
		if val > first {
			second = first
			first = val
		} else if val > second {
			third = second
			second = val
		} else if val > third {
			third = val
		}
	}
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	maxVal := first + second + third
	return maxVal
}

func calcCalories(inp []string) []int {
	output := []int{}
	count := 0
	for i := range inp {
		if inp[i] == "" {
			output = append(output, count)
			count = 0
		}

		val, _ := strconv.Atoi(inp[i])
		count = count + val
	}
	return output
}

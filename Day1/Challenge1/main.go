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
	ans := findMostCal(input)
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

func findMostCal(input []string) int {
	calCounts := calcCalories(input)
	maxVal := 0
	for _, val := range calCounts {
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal
}

func calcCalories(inp []string) []int {
	output := []int{}
	//ElveCount := 0
	count := 0
	for i, _ := range inp {
		if inp[i] == "" {
			output = append(output, count)
			count = 0
			//		ElveCount = 0
		}

		val, _ := strconv.Atoi(inp[i])
		count = count + val
	}
	return output
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type cleanuparea struct {
	min1 int
	max1 int
	min2 int
	max2 int
}

func main() {
	input := getFile()
	out := getAnswer(input)
	fmt.Println(out)
}

func getAnswer(input []string) int {
	answer := 0
	cleanupareas := getCleanUpAreas(input)
	for _, i := range cleanupareas {
		if checkAreas(i) {
			fmt.Println(i)
			answer++
		}
	}
	return answer
}

func checkAreas(input cleanuparea) bool {
	if ((input.min2 >= input.min1) && (input.min2 <= input.max1)) && ((input.max2 >= input.min1) && (input.max2 <= input.max1)) {
		fmt.Println(input)
		return true
	} else if ((input.min1 >= input.min2) && (input.min1 <= input.max2)) && ((input.max1 >= input.min2) && (input.max1 <= input.max2)) {
		fmt.Println(input)
		return true
	} else {
		return false
	}
}

func getCleanUpAreas(input []string) []cleanuparea {
	returnVal := []cleanuparea{}
	for _, i := range input {
		temp := getCleanUpArea(i)
		returnVal = append(returnVal, temp)
	}
	return returnVal
}

func getCleanUpArea(input string) cleanuparea {
	s := strings.Split(input, "-")
	s2 := strings.Split(s[1], ",")
	min1, _ := strconv.Atoi(s[0])
	max1, _ := strconv.Atoi(s2[0])
	min2, _ := strconv.Atoi(s2[1])
	max2, _ := strconv.Atoi(s[2])
	return cleanuparea{
		min1: min1,
		max1: max1,
		min2: min2,
		max2: max2,
	}
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

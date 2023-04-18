package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var Alphabet = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func main() {
	input := getFile()
	outcome := calcOutcome(input)
	fmt.Println(outcome)
}

func calcOutcome(input []string) int {
	returnVal := 0

	for _, i := range input {
		common := getCommonItem(i)
		val := findValue(common)
		returnVal = returnVal + val
	}

	return returnVal
}

func getCommonItem(input string) string {
	var returnval string
	length := len(input) / 2
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if input[i] == input[j+length] {
				returnval = string(input[i])
			}
		}
	}
	return returnval
}

func findValue(input string) int {
	ReturnVal := 0
	for i, a := range Alphabet {
		if a == input {
			ReturnVal = i + 1
		}
	}
	return ReturnVal
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

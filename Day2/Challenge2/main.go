package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := getFile()
	ans := rockPaperScissor(input)
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

func rockPaperScissor(input []string) int {
	totScore := 0

	for _, val := range input {
		switch string(val[2]) {
		case "X":
			totScore = totScore + 0
			if string(val[0]) == "A" {
				totScore = totScore + 3
				break
			}
			if string(val[0]) == "B" {
				totScore = totScore + 1
				break
			}
			if string(val[0]) == "C" {
				totScore = totScore + 2
				break
			}
		case "Y":
			totScore = totScore + 3
			if string(val[0]) == "A" {
				totScore = totScore + 1
				break
			}
			if string(val[0]) == "B" {
				totScore = totScore + 2
				break
			}
			if string(val[0]) == "C" {
				totScore = totScore + 3
				break
			}
		case "Z":
			totScore = totScore + 6
			if string(val[0]) == "A" {
				totScore = totScore + 2
				break
			}
			if string(val[0]) == "B" {
				totScore = totScore + 3
				break
			}
			if string(val[0]) == "C" {
				totScore = totScore + 1
				break
			}
		}

	}
	return totScore
}

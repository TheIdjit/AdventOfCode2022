package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var pallets = [][]string{
	{"F", "H", "B", "V", "R", "Q", "D", "P"},
	{"L", "D", "Z", "Q", "W", "V"},
	{"H", "L", "Z", "Q", "G", "R", "P", "C"},
	{"R", "D", "H", "F", "J", "V", "B"},
	{"Z", "W", "L", "C"},
	{"J", "R", "P", "N", "T", "G", "V", "M"},
	{"J", "R", "L", "V", "M", "B", "S"},
	{"D", "P", "J"},
	{"D", "C", "N", "W", "V"},
}

func main() {
	input := getFile()
	containerMovement(input)
	fmt.Println(pallets[0][len(pallets[0])-1])
	fmt.Println(pallets[1][len(pallets[1])-1])
	fmt.Println(pallets[2][len(pallets[2])-1])
	fmt.Println(pallets[3][len(pallets[3])-1])
	fmt.Println(pallets[4][len(pallets[4])-1])
	fmt.Println(pallets[5][len(pallets[5])-1])
	fmt.Println(pallets[6][len(pallets[6])-1])
	fmt.Println(pallets[7][len(pallets[7])-1])
	fmt.Println(pallets[8][len(pallets[8])-1])
}

func containerMovement(input []string) {
	for _, i := range input {
		action := strings.Split(i, " ")
		numOfPallets, _ := strconv.Atoi(action[1])
		Source, _ := strconv.Atoi(action[3])
		Dest, _ := strconv.Atoi(action[5])
		fmt.Println(len(pallets[Source-1][:]))

		for j := 0; j < numOfPallets; j++ {
			pallets[Dest-1] = append(pallets[Dest-1], pallets[Source-1][len(pallets[Source-1][:])-1])
			fmt.Println(pallets[Source-1][:len(pallets[Source-1][:])-1])
			pallets[Source-1] = pallets[Source-1][:len(pallets[Source-1][:])-1]
			fmt.Println(pallets)
		}
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

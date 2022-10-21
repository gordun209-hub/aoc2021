package main

import (
	"fmt"
	"strings"
)

var data = `
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
var print = fmt.Println

func main() {
	formattedData := strings.TrimSpace(data)
	formattedDataa := strings.Split(formattedData, "\n")
	answer := calculateOxygenRate(formattedDataa)
	print(answer)
}

func calculateOxygenRate(data []string) []string {
	return removeMostCommonBit(data, 0)
}

func removeMostCommonBit(data []string, index int) []string {
	var holdval = []string{}
	mostCommonBit := findMostCommonBitAtGivenIndex(data, index)
	for i, v := range data {
		if v[0] == byte(mostCommonBit) {
			holdval[i] = v
		}
	}
	return holdval

}

func findMostCommonBitAtGivenIndex(data []string, index int) int {
	ones := 0
	zeros := 0
	for _, v := range data {
		// check single element inside array
		if v[index] == 48 {
			zeros++
		} else {
			ones++
		}
	}

	print(zeros, ones)
	if zeros > ones {
		return 48
	} else if ones > zeros {
		return 49
	} else {
		return 31
	}
}

func remove(slice []string, s int) []string {
	slice = append(slice[:s], slice[s+1:]...)
	return slice
}

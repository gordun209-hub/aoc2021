package main

import (
	"fmt"
	"os"
	"strings"
)

type (
	elem string
	Row  struct {
		elements []string
		isMarked bool
	}

	Col struct {
		elements []string
		isMarked bool
	}

	Board struct {
		rows  [5]Row
		cols  [5]Col
		isWon bool
	}
)

var boardData = `
22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
1 12 20 15 19
`
var numbers = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1`

func main() {
	var myboard = Board{}
	sl := strings.TrimSpace(boardData)
	sll := strings.Fields(sl)
	for i := range sll {
		if i < 5 {

			myboard.cols[i].elements = getCol(sll, i)
			myboard.rows[i].elements = getRow(sll, i)
			myboard.cols[i].isMarked = false
			myboard.rows[i].isMarked = false
		}
	}

	numberss := strings.TrimSpace(numbers)
	numbr := strings.Split(numberss, ",")
	iswonn := checkIfWin(myboard, numbr)
	fmt.Println(iswonn)

}
func checkIfWin(myboard Board, numbr []string) bool {
	iswon := 0
	// iterate numbers
	for _, v := range numbr {
		// iterate cols elements
		for i := 0; i < len(myboard.cols); i++ {
			if v == myboard.cols[0].elements[i] {
				iswon++
				if iswon == 5 {
					myboard.isWon = true

				}
			}
		}
	}

	for _, v := range numbr {
		// iterate cols elements
		for i := 0; i < len(myboard.cols); i++ {
			if v == myboard.cols[1].elements[i] {
				iswon++
				if iswon == 5 {
					myboard.isWon = true

				}
			}
		}
	}
	for _, v := range numbr {
		// iterate cols elements
		for i := 0; i < len(myboard.cols); i++ {
			if v == myboard.cols[2].elements[i] {
				iswon++
				if iswon == 5 {
					myboard.isWon = true

				}
			}
		}
	}
	for _, v := range numbr {
		// iterate cols elements
		for i := 0; i < len(myboard.cols); i++ {
			if v == myboard.cols[3].elements[i] {
				iswon++
				if iswon == 5 {
					myboard.isWon = true

				}
			}
		}
	}
	for _, v := range numbr {
		// iterate cols elements
		for i := 0; i < len(myboard.cols); i++ {
			if v == myboard.cols[4].elements[i] {
				iswon++
				if iswon == 5 {
					myboard.isWon = true

				}
			}
		}
	}
	return myboard.isWon
}
func getRow(data []string, index int) []string {
	row := []string{}
	for i := 0; i < 5; i++ {
		row = append(row, data[i+5*index])
	}
	return row
}

func getCol(data []string, index int) []string {
	newdata := []string{}
	for i := 0; i < len(data); i += 5 {
		newdata = append(newdata, data[index+i])
	}
	return newdata
}

func getInput() string {
	content, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("err")
	}

	return string(content)
}

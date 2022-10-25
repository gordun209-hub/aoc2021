package main

// var boardData = `
// 22 13 17 11  0
//  8  2 23  4 24
// 21  9 14 16  7
//  6 10  3 18  5
// 1 12 20 15 19
// `
// var numbers = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1`
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

func main() {
	// var myboard = Board{}
	inputs := getInput()
	// nmbrs := strings.Fields(inputs)
	numbers, tables, _ := strings.Cut(inputs, "\n")
	trimmed := strings.TrimSpace(tables)
	spldd := strings.Fields(trimmed)
	nmbrs := strings.Split(numbers, ",")
	// fmt.Println(checkIfWin(createBoard(spldd[0:25]), nmbrs))
	for i := 0; i < len(spldd); i += 25 {
		boarss := createBoard(spldd[i : i+25])
		brd := checkIfWin(boarss, nmbrs)
		fmt.Println("-------------------------")
		fmt.Println(brd)
		fmt.Println("-------------------------")
		//
	}
	// iswonn := checkIfWin(myboard, numbr)

}
func createBoard(sll []string) Board {
	var myboard = Board{}
	for i := range sll {
		if i < 5 {

			myboard.cols[i].elements = getCol(sll, i)
			myboard.rows[i].elements = getRow(sll, i)
			myboard.cols[i].isMarked = false
			myboard.rows[i].isMarked = false
		}
	}
	return myboard
}
func checkIfWin(myboard Board, numbr []string) Board {
	// iterate numbers
	iswon0 := 0
	for _, v := range numbr {
		// iterate cols elements
		for i := 0; i < len(myboard.cols); i++ {
			if v == myboard.cols[0].elements[i] {
				iswon0++
				if iswon0 == 5 {
					return myboard

				}
			}
		}
	}
	iswon1 := 0
	for _, v := range numbr {
		// iterate cols elements
		for i := 0; i < len(myboard.cols); i++ {
			if v == myboard.cols[1].elements[i] {
				iswon1 = iswon1 + 1
				if iswon1 == 5 {
					return myboard

				}
			}
		}
	}
	iswon2 := 0
	for _, v := range numbr {
		// iterate cols elements
		for i := 0; i < len(myboard.cols); i++ {
			if v == myboard.cols[2].elements[i] {
				iswon2++
				if iswon2 == 5 {
					return myboard

				}
			}
		}
	}
	iswon3 := 0
	for _, v := range numbr {
		// iterate cols elements
		for i := 0; i < len(myboard.cols); i++ {
			if v == myboard.cols[3].elements[i] {
				iswon3++
				if iswon3 == 5 {
					return myboard

				}
			}
		}
	}
	iswon4 := 0
	for _, v := range numbr {
		// iterate cols elements
		for i := 0; i < len(myboard.cols); i++ {
			if v == myboard.cols[4].elements[i] {
				iswon4++
				if iswon4 == 5 {
					return myboard

				}
			}
		}
	}
	iswon5 := 0
	for _, v := range numbr {
		// iterate cols elements
		for i := 0; i < len(myboard.rows); i++ {
			if v == myboard.rows[4].elements[i] {
				iswon5++
				if iswon5 == 5 {
					return myboard

				}
			}
		}
	}
	iswon6 := 0
	for _, v := range numbr {
		// iterate cols elements
		for i := 0; i < len(myboard.rows); i++ {
			if v == myboard.rows[3].elements[i] {
				iswon6++
				if iswon6 == 5 {
					return myboard

				}
			}
		}
	}
	iswon7 := 0
	for _, v := range numbr {
		// iterate cols elements
		for i := 0; i < len(myboard.rows); i++ {
			if v == myboard.rows[2].elements[i] {
				iswon7++
				if iswon7 == 5 {
					return myboard

				}
			}
		}
	}
	iswon8 := 0
	for _, v := range numbr {
		// iterate cols elements
		for i := 0; i < len(myboard.rows); i++ {
			if v == myboard.rows[1].elements[i] {
				iswon8++
				if iswon8 == 5 {
					return myboard

				}
			}
		}
	}
	iswon9 := 0
	for _, v := range numbr {
		// iterate cols elements
		for i := 0; i < len(myboard.rows); i++ {
			if v == myboard.rows[0].elements[i] {
				iswon9 = iswon9 + 1
				if iswon9 == 5 {
					return myboard

				}
			}
		}
	}

	return Board{}
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

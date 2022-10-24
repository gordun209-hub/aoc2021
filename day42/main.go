package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputs := getInput()
	numbers, tables, _ := strings.Cut(inputs, "\n")
	trimmed := strings.TrimSpace(tables)
	spldd := strings.Fields(trimmed)
	fmt.Println(spldd[25:])
	mytabl := makeTable(spldd)
	fmt.Println(mytabl)

}
func nextTable(data []string) []string {
	return data[25:]
}
func makeTable(data []string) [5][5]string {
	var twoDtables = [5][5]string{}
	x, y := 0, 0
	for _, v := range data {
		twoDtables[x][y] = v
		y++
		if y == 5 {
			x++
			y = 0

		}
		if x == 5 {
			break
		}
	}
	return twoDtables
}

// func createBoard(sll []string) Board {
// 	var myboard = Board{}
// 	for i := range sll {
// 		if i < 5 {
//
// 			myboard.cols[i].elements = getCol(sll, i)
// 			myboard.rows[i].elements = getRow(sll, i)
// 			myboard.cols[i].isMarked = false
// 			myboard.rows[i].isMarked = false
// 		}
// 	}
// 	return myboard
// }
// func checkIfWin(myboard Board, numbr []string) bool {
// 	iswon := 0
// 	// iterate numbers
// 	for _, v := range numbr {
// 		// iterate cols elements
// 		for i := 0; i < len(myboard.cols); i++ {
// 			if v == myboard.cols[0].elements[i] {
// 				iswon++
// 				if iswon == 5 {
// 					myboard.isWon = true
//
// 				}
// 			}
// 		}
// 	}

//		for _, v := range numbr {
//			// iterate cols elements
//			for i := 0; i < len(myboard.cols); i++ {
//				if v == myboard.cols[1].elements[i] {
//					iswon++
//					if iswon == 5 {
//						myboard.isWon = true
//
//					}
//				}
//			}
//		}
//		for _, v := range numbr {
//			// iterate cols elements
//			for i := 0; i < len(myboard.cols); i++ {
//				if v == myboard.cols[2].elements[i] {
//					iswon++
//					if iswon == 5 {
//						myboard.isWon = true
//
//					}
//				}
//			}
//		}
//		for _, v := range numbr {
//			// iterate cols elements
//			for i := 0; i < len(myboard.cols); i++ {
//				if v == myboard.cols[3].elements[i] {
//					iswon++
//					if iswon == 5 {
//						myboard.isWon = true
//
//					}
//				}
//			}
//		}
//		for _, v := range numbr {
//			// iterate cols elements
//			for i := 0; i < len(myboard.cols); i++ {
//				if v == myboard.cols[4].elements[i] {
//					iswon++
//					if iswon == 5 {
//						myboard.isWon = true
//
//					}
//				}
//			}
//		}
//		return myboard.isWon
//	}
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

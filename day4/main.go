package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	boardWidth    = 5
	boardHeight   = 5
	boardCapacity = 25
)

type cell [2]string
type row [4]cell
type col [4]cell
type Board struct {
}

var board = `
22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19
`

func main() {
	splitted := strings.Split(board, " ")
	fmt.Println(splitted)
	var lna = cell{"2", "2"}
	var lna1 = cell{"17", "11"}
	var lna2 = cell{" ", "3"}
	var lna3 = cell{" ", "3"}
	var myrow = row{lna, lna1, lna2, lna3}
	fmt.Println(myrow)

}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func getInput() string {

	content, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("err")
	}

	return string(content)
}

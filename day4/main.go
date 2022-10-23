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

	sl := strings.TrimSpace(board)
	// splitted := strings.Split(sl, "\n")
	// fmt.Printf("%v", splitted)
	// rows := getRows(strings.Fields(sl))
	cols := getCols(strings.Fields(sl))
	fmt.Println(cols)

}
func getRows(data []string) []string {
	newdata := []string{}
	for i, v := range data {
		if i < 5 {
			fmt.Println(v)
			newdata = append(newdata, v)
		}
	}
	return newdata

}
func getCols(data []string) []string {
	newdata := []string{}
	for i := 0; i < len(data); i += 5 {
		fmt.Println(data[i])
		newdata = append(newdata, data[i])
	}
	return newdata
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

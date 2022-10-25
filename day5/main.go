package main

import (
	"fmt"
	"strconv"
	"strings"
)

// x1,y1 -> x2,y2
// 0, 9  -> 5 ,9
var data = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

var product = `.......1..
..1....1..
..1....1..
.......1..
.112111211
..........
..........
..........
..........
222111....`

// example
// entry 1,1 -> 1,3 covers 1,1 1,2 and 1,3
// 9,7 7,7 covers 9,7 8,7 and 7,7
func main() {
	data1 := strings.ReplaceAll(data, "->", " ")
	data2 := strings.Fields(data1)
	// coveredPoints := getPair(data2)
	for i := 0; i < len(data2); i += 2 {
		getCoveredPoints(data2[i : i+2])
	}
	// getCoveredPoints(coveredPoints)
}

func getPair(points []string) []string {

	return []string{points[0], points[1]}

}
func getCoveredPoints(data []string) {
	coveredPoints := []int{}
	firstPointPair := data[0]
	secondPointPair := data[1]
	x1 := strings.Split(firstPointPair, ",")[0]
	x2 := strings.Split(firstPointPair, ",")[1]
	y1 := strings.Split(secondPointPair, ",")[0]
	y2 := strings.Split(secondPointPair, ",")[1]
	x1Int, err := strconv.Atoi(x1)
	if err != nil {
		fmt.Println(err)
	}
	x2Int, err := strconv.Atoi(x2)
	y1Int, err := strconv.Atoi(y1)
	y2Int, err := strconv.Atoi(y2)

	if x1Int == x2Int {
		fmt.Println(x1Int, x2Int)
		for i := 0; y1Int < y2Int; i++ {
			coveredPoints = append(coveredPoints, y1Int)
			y1Int++
		}
	}
	if y1Int == y2Int {
		fmt.Println(y1Int, y2Int)
		for i := 0; x1Int < x2Int; i++ {
			coveredPoints = append(coveredPoints, x1Int)
			x1Int++
		}
	}
	fmt.Println(coveredPoints)
}

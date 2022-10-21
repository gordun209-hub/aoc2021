package main

import (
	"fmt"
	"os"
	"strings"
)

// gamma rate and epsilon rate
// power consuption = gamma rate * epsilon rate

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("err")
	}
	trimmed := strings.TrimSpace(string(content))
	trimme := strings.Split(trimmed, "\n")
	data1 := []string{}
	data2 := []string{}
	data3 := []string{}
	data4 := []string{}
	data5 := []string{}
	data6 := []string{}
	data7 := []string{}
	data8 := []string{}
	data9 := []string{}
	data10 := []string{}
	data11 := []string{}
	data12 := []string{}
	for _, val := range trimme {
		data1 = append(data1, string(val[0]))
		data2 = append(data2, string(val[1]))
		data3 = append(data3, string(val[2]))
		data4 = append(data4, string(val[3]))
		data5 = append(data5, string(val[4]))
		data6 = append(data6, string(val[5]))
		data7 = append(data7, string(val[6]))
		data8 = append(data8, string(val[7]))
		data9 = append(data9, string(val[8]))
		data10 = append(data10, string(val[9]))
		data11 = append(data11, string(val[10]))
		data12 = append(data12, string(val[11]))
	}
	fmt.Println(1503 * 2592)
}

func monstCommonBits(data []string) int {
	ones := 0
	zeros := 0
	for _, val := range data {
		if val == "0" {
			zeros++
		} else {
			ones++
		}
	}
	if ones > zeros {
		return 1
	} else {
		return 0
	}
}

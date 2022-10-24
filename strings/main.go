package main

import (
	"fmt"
	"strings"
)

func main() {
	lna := `la 2 mq occ
    mqq lna una `
	splt := strings.Fields(lna)
	fmt.Println(splt[2])
}

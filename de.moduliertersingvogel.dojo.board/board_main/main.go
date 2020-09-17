package main

import (
	"de.moduliertersingvogel.dojo.board/board"
	"fmt"
	"strings"
)

func main() {
	stonefactory := func() board.Stone {
		stone := board.Stone{}
		return stone
	}
	fieldfactory := func(x int, y int) board.Field {
		var color string
		if (x+y)%2 == 0 {
			color = "black"
		} else {
			color = "white"
		}
		field := board.Field{Color: &color}
		return field
	}
	board := board.NewBoard(8, stonefactory, fieldfactory)

	for i := range board.Fields {
		var str1 strings.Builder
		var str2 strings.Builder
		var str3 strings.Builder

		for j := range board.Fields[i] {
			field := board.Fields[i][j]
			if *(field.Color) == "black" {
				str1.WriteString("#####   ")
				str2.WriteString("#####   ")
				str3.WriteString("#####   ")
			} else {
				str1.WriteString("-----   ")
				str2.WriteString("|   |   ")
				str3.WriteString("-----   ")
			}
		}
		fmt.Println(str1.String())
		fmt.Println(str2.String())
		fmt.Println(str3.String())
		fmt.Println()
	}
}

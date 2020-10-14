package main

import (
	"de.moduliertersingvogel.dojo.board/board"
	"de.moduliertersingvogel.dojo.board/engine"
	"fmt"
	"github.com/gookit/color"
	"strconv"
)

var (
	blackfield      = color.S256(0, 242)
	blackfieldblack = color.S256(0, 242)
	blackfieldwhite = color.S256(255, 242)
	whitefield      = color.S256(0, 255)
	whitefieldblack = color.S256(0, 255)
	whitefieldwhite = color.S256(250, 255)
	abc             = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
)

func draw(board board.Board) {
	fmt.Print("\033[2J")
	fmt.Println()

	for i := len(board.Fields) - 1; i >= 0; i-- {
		cellsize := 5

		for j := range board.Fields[i] {
			field := board.Fields[i][j]

			if *(field.Color) == "black" {
				blackfield.Printf("  %s-%d  ", abc[j], i+1)
			} else {
				whitefield.Printf("  %s-%d  ", abc[j], i+1)
			}
		}
		fmt.Println()

		for k := 1; k < cellsize-1; k++ {
			for j := range board.Fields[i] {
				field := board.Fields[i][j]
				if field.Stone != nil {
					if *(field.Color) == "black" {
						blackfield.Printf("  ")
						if *(field.Stone.Color) == "black" {
							blackfieldblack.Printf("%s", field.Stone.Symbol[k-1])
						} else {
							blackfieldwhite.Printf("%s", field.Stone.Symbol[k-1])
						}
						blackfield.Printf("  ")
					} else {
						whitefield.Printf("  ")
						if *(field.Stone.Color) == "black" {
							whitefieldblack.Printf("%s", field.Stone.Symbol[k-1])
						} else {
							whitefieldwhite.Printf("%s", field.Stone.Symbol[k-1])
						}
						whitefield.Printf("  ")
					}
				} else {
					if *(field.Color) == "black" {
						blackfield.Printf("       ")
					} else {
						whitefield.Printf("       ")
					}
				}
			}
			fmt.Println()
		}
	}
}

func initgame() board.Board {
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
	blackcolor := "black"
	whitecolor := "white"
	blacktower := board.Stone{[]string{"\u25A0\u25A0\u25A0", "\u25A0\u25A0\u25A0", "\u25A0\u25A0\u25A0"}, &blackcolor}
	blackbishop := board.Stone{[]string{" \u25A0 ", " \u25A0 ", "\u25A0\u25A0\u25A0"}, &blackcolor}
	blackknight := board.Stone{[]string{"\u25A0\u25A0\u25A0", " \u25A0 ", " \u25A0 "}, &blackcolor}
	blackking := board.Stone{[]string{" \u25A0 ", "\u25A0\u25A0\u25A0", "\u25A0\u25A0\u25A0"}, &blackcolor}
	blackqueen := board.Stone{[]string{"\u25A0\u25A0\u25A0", " \u25A0 ", "\u25A0\u25A0\u25A0"}, &blackcolor}
	blackpawn := board.Stone{[]string{"   ", " \u25A0 ", "\u25A0\u25A0\u25A0"}, &blackcolor}
	whitetower := board.Stone{[]string{"\u25A0\u25A0\u25A0", "\u25A0\u25A0\u25A0", "\u25A0\u25A0\u25A0"}, &whitecolor}
	whitebishop := board.Stone{[]string{" \u25A0 ", " \u25A0 ", "\u25A0\u25A0\u25A0"}, &whitecolor}
	whiteknight := board.Stone{[]string{"\u25A0\u25A0\u25A0", " \u25A0 ", " \u25A0 "}, &whitecolor}
	whiteking := board.Stone{[]string{" \u25A0 ", "\u25A0\u25A0\u25A0", "\u25A0\u25A0\u25A0"}, &whitecolor}
	whitequeen := board.Stone{[]string{"\u25A0\u25A0\u25A0", " \u25A0 ", "\u25A0\u25A0\u25A0"}, &whitecolor}
	whitepawn := board.Stone{[]string{"   ", " \u25A0 ", "\u25A0\u25A0\u25A0"}, &whitecolor}
	size := 8

	board := board.NewBoard(size, fieldfactory)

	board.Fields[7][0].Stone = &blacktower
	board.Fields[7][size-1].Stone = &blacktower
	board.Fields[7][1].Stone = &blackknight
	board.Fields[7][size-2].Stone = &blackknight
	board.Fields[7][2].Stone = &blackbishop
	board.Fields[7][size-3].Stone = &blackbishop
	board.Fields[7][3].Stone = &blackqueen
	board.Fields[7][4].Stone = &blackking
	for i := 0; i < size; i++ {
		board.Fields[6][i].Stone = &blackpawn
	}
	board.Fields[0][0].Stone = &whitetower
	board.Fields[0][7].Stone = &whitetower
	board.Fields[0][1].Stone = &whiteknight
	board.Fields[0][size-2].Stone = &whiteknight
	board.Fields[0][2].Stone = &whitebishop
	board.Fields[0][size-3].Stone = &whitebishop
	board.Fields[0][3].Stone = &whitequeen
	board.Fields[0][4].Stone = &whiteking
	for i := 0; i < size; i++ {
		board.Fields[1][i].Stone = &whitepawn
	}

	return board
}

func eventloop(board board.Board) {
	active := "white"
	for {
		if active == "black" {
			engine.NextMove(board, active)

			active = "white"
			continue
		}

		var text string

		fmt.Scanln(&text)
		if len(text) != 4 {
			continue
		}
		ycoordfrom := -1
		ycoordto := -1
		for i := range abc {
			if abc[i] == string(text[0]) {
				ycoordfrom = i
			}
			if abc[i] == string(text[2]) {
				ycoordto = i
			}
		}
		xcoordfrom, _ := strconv.Atoi(string(text[1]))
		xcoordto, _ := strconv.Atoi(string(text[3]))
		if board.Fields[xcoordfrom-1][ycoordfrom].Stone != nil {
			board.Fields[xcoordto-1][ycoordto].Stone = &(*(board.Fields[xcoordfrom-1][ycoordfrom].Stone))
			board.Fields[xcoordfrom-1][ycoordfrom].Stone = nil

			active = "black"
		} else {
			continue
		}

		draw(board)
	}
}

func main() {
	board := initgame()
	draw(board)
	eventloop(board)
}

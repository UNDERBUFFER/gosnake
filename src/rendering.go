package main

import (
	"fmt"

	"github.com/pborman/ansi"
)

func DrawBorders() {
	ClearDisplay()

	drawX := func() {
		for i := 0; i < (MAP_WIDTH + 2); i++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	drawY := func() {
		for i := 0; i < MAP_HEIGHT; i++ {
			res := ""
			for j := 0; j < MAP_WIDTH; j++ {
				res += " "
			}
			fmt.Printf("#%s#\n", res)
		}
	}

	drawX()
	drawY()
	drawX()
}

func ClearDisplay() {
	fmt.Print(ansi.RIS, "\r")
}

func ResetCursor() {
	fmt.Print(ansi.CSI, "2;2f")
}

func InputCursor() {
	fmt.Print(ansi.CSI, MAP_HEIGHT+7+LOGS_COUNT, ";0f")
}

func CursorLogs() {
	fmt.Print(ansi.CSI, MAP_HEIGHT+4, ";0f")
}

// func FixGreat() {
// 	fmt.Print(ansi.CSI, MAP_HEIGHT+11, ";0f")
// }

func WriteImage(x int, y int, img string) {
	ResetCursor()
	fmt.Print(ansi.CSI, y+2, ";", x+2, "f")
	fmt.Print(img)
}

func ShowFrame(
	player GameObject,
	fruits GameObject,
) {
	DrawBorders()
	ResetCursor()
	for _, val := range fruits {
		WriteImage(val[0], val[1], "🍎")
	}
	for _, val := range player {
		WriteImage(val[0], val[1], "▩")
	}
}

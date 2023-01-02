package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pborman/ansi"
)

const (
	y = 10
	x = 20
)

func DrawBorders(x int, y int) {
	fmt.Print(ansi.RIS, "\r")

	drawX := func() {
		for i := 0; i < (x + 2); i++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	drawY := func() {
		for i := 0; i < y; i++ {
			res := ""
			for j := 0; j < x; j++ {
				res += " "
			}
			fmt.Printf("#%s#\n", res)
		}
	}

	drawX()
	drawY()
	drawX()
}

func ResetCursor() {
	fmt.Print(ansi.CSI, "2;2f")
}

func FixGreat() {
	fmt.Print(ansi.CSI, y + 4, ";0f")
}

func WriteHead(x int, y int) {
	ResetCursor()
	for i := 0; i < y; i++ {
		fmt.Print(ansi.CUD)
	}
	for i := 0; i < x; i++ {
		fmt.Print(ansi.CUF)
	}
	fmt.Print("0")
}

func ReadMove() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your move: ")
	move, _ := reader.ReadString('\n')
	move = strings.TrimSuffix(move, "\n")
	fmt.Print(ansi.CUU, ansi.DL)
	return move
}

func ShowFrame(posX int, posY int) {
	DrawBorders(x, y)
	ResetCursor()
	WriteHead(posX, posY)
}

func main() {
	// defer FixGreat()
	var (
		posX = 0
		posY = 0
	)
	ShowFrame(posX, posY)
	FixGreat()

	for {
		switch ReadMove() {
		case "S":
			if posY < (y - 1) {
				posY++
			}
		case "D":
			if posX < (x - 1) {
				posX++
			}
		case "W":
			if (posY - 1) >= 0 {
				posY--
			}
		case "A":
			if (posX - 1) >= 0 {
				posX--
			}
		}
		ShowFrame(posX, posY)
		FixGreat()
	}

}

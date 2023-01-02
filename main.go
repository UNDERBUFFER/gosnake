package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pborman/ansi"
	"golang.org/x/exp/slices"
)

type button string

const (
	y             = 10
	x             = 20
	TOP    button = "W"
	BOTTOM button = "S"
	LEFT   button = "A"
	RIGHT  button = "D"
	CENTER button = ""
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
	fmt.Print(ansi.CSI, y+4, ";0f")
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

func ReadMove() button {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your move: ")
	move, _ := reader.ReadString('\n')
	move = strings.ToUpper(strings.TrimSuffix(move, "\n"))
	if slices.Contains([]button{TOP, BOTTOM, RIGHT, LEFT}, button(move)) {
		return button(move)
	}
	return CENTER
}

func ShowFrame(posX int, posY int) {
	DrawBorders(x, y)
	ResetCursor()
	WriteHead(posX, posY)
}

func UpdatePosition(move button, posX *int, posY *int, mapX int, mapY int) {
	switch move {
	case TOP:
		if (*posY - 1) >= 0 {
			*posY--
		}
	case BOTTOM:
		if (*posY + 1) < mapY {
			*posY++
		}
	case LEFT:
		if (*posX - 1) >= 0 {
			*posX--
		}
	case RIGHT:
		if (*posX + 1) < mapX {
			*posX++
		}
	}
}

func main() {
	var (
		posX = 0
		posY = 0
	)

	for {
		ShowFrame(posX, posY)
		FixGreat()
		UpdatePosition(ReadMove(), &posX, &posY, x, y)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/pborman/ansi"
	"golang.org/x/exp/slices"
)

type Button string
type GameObject [][2]int

const (
	MAP_HEIGHT        = 10
	MAP_WIDTH         = 20
	TOP        Button = "W"
	BOTTOM     Button = "S"
	LEFT       Button = "A"
	RIGHT      Button = "D"
	CENTER     Button = ""
)

var SKIP = [2]int{-1, -1}

func DrawBorders() {
	fmt.Print(ansi.RIS, "\r")

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

func ResetCursor() {
	fmt.Print(ansi.CSI, "2;2f")
}

func FixGreat() {
	fmt.Print(ansi.CSI, MAP_HEIGHT+4, ";0f")
}

func WriteImage(x int, y int, img string) {
	ResetCursor()
	fmt.Print(ansi.CSI, y+2, ";", x+2, "f")
	fmt.Print(img)
}

func ReadMove() Button {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your move: ")
	move, _ := reader.ReadString('\n')
	move = strings.ToUpper(
		strings.TrimSuffix(move, "\n"),
	)
	if slices.Contains([]Button{
		TOP,
		BOTTOM,
		RIGHT,
		LEFT,
	}, Button(move)) {
		return Button(move)
	}
	return CENTER
}

func ShowFrame(player GameObject) {
	DrawBorders()
	ResetCursor()
	for _, val := range player {
		WriteImage(val[0], val[1], "0")
	}
}

func UpdatePosition(move Button, player GameObject) [2]int {
	var (
		x = player[0][0]
		y = player[0][1]
	)
	switch move {
	case TOP:
		if (y - 1) >= 0 {
			y--
		} else {
			y = MAP_HEIGHT - 1
		}
	case BOTTOM:
		if (y + 1) < MAP_HEIGHT {
			y++
		} else {
			y = 0
		}
	case LEFT:
		if (x - 1) >= 0 {
			x--
		} else {
			x = MAP_WIDTH - 1
		}
	case RIGHT:
		if (x + 1) < MAP_WIDTH {
			x++
		} else {
			x = 0
		}
	}
	coords := [2]int{x, y}

	for _, val := range player {
		if reflect.DeepEqual(val, coords) {
			return SKIP
		}
	}
	return [2]int{x, y}
}

func MovePlayer(x int, y int, player *GameObject) {
	*player = append(
		GameObject{{x, y}},
		(*player)[:len((*player))-1]...,
	)
}

func main() {
	player := GameObject{{0, 0}, {1, 0}, {1, 1}}

	ShowFrame(player)
	FixGreat()

	for {
		ShowFrame(player)
		FixGreat()

		move := ReadMove()

		coords := UpdatePosition(move, player)
		if !reflect.DeepEqual(coords, SKIP) {
			MovePlayer(coords[0], coords[1], &player)
		}
	}
}

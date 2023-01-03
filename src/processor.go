package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strings"

	"golang.org/x/exp/slices"
)

func ReadMove() Button {
	InputCursor()
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

func UpdatePosition(
	move Button,
	player GameObject,
) [2]int {
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
	return coords
}

func MoveObjects(
	x int,
	y int,
	player *GameObject,
	fruits *GameObject,
) {
	head := [2]int{x, y}
	length := len(*player) - 1

	for idx, val := range *fruits {
		if reflect.DeepEqual(val, head) {
			length = len(*player)
			*fruits = append(
				(*fruits)[:idx],
				(*fruits)[idx+1:]...,
			)
			break
		}
	}

	*player = append(
		GameObject{head},
		(*player)[:length]...,
	)
}

func SetupFruits(
	fruits *GameObject,
	player GameObject,
) {
Loop:
	for len(*fruits) < FRUITS_COUNT {
		coords := [2]int{
			rand.Intn(MAP_WIDTH),
			rand.Intn(MAP_HEIGHT),
		}

		for _, val := range player {
			if reflect.DeepEqual(val, coords) {
				continue Loop
			}
		}

		*fruits = append(
			*fruits,
			coords,
		)
	}
}

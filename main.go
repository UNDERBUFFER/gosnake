package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer ClearDisplay()

	player := GameObject{{0, 0}}
	fruits := GameObject{}

	logger := Logger{}

	for {
		logger.Info(fmt.Sprint(player, "|", fruits))
		SetupFruits(&fruits, player)
		ShowFrame(player, fruits)
		logger.GetLogs()

		move := ReadMove()
		coords := UpdatePosition(move, player)

		if !reflect.DeepEqual(coords, SKIP) {
			MoveObjects(
				coords[0],
				coords[1],
				&player,
				&fruits,
			)
		}
	}
}

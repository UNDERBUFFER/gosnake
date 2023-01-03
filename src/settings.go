package main

type Button string
type Position int
type GameObject [][2]int

const (
	MAP_HEIGHT          = 10
	MAP_WIDTH           = 20
	FRUITS_COUNT        = 5
	LOGS_COUNT          = 3
	TOP          Button = "W"
	BOTTOM       Button = "S"
	LEFT         Button = "A"
	RIGHT        Button = "D"
	CENTER       Button = ""
)

var SKIP = [2]int{-1, -1}

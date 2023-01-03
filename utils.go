package main

import (
	"fmt"
)

type Logger struct {
	messages []string
}

func (l Logger) GetLogs() {
	CursorLogs()

	border := ""
	for i := 0; i < MAP_WIDTH+2; i++ {
		border += "="
	}
	fmt.Println(border)

	length := len(l.messages)
	idx := 0
	if length > LOGS_COUNT {
		idx = length - LOGS_COUNT
	}
	for _, log := range l.messages[idx:] {
		fmt.Println(log)
	}
	fmt.Println(border)
}

func (l *Logger) Info(msg string) {
	l.messages = append(l.messages, fmt.Sprintf(
		"%d: %s",
		len(l.messages),
		msg,
	))
}

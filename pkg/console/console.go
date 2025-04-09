package console

import (
	"github.com/fatih/color"
)

var Red = color.New(color.FgRed).SprintfFunc()

var Green = color.New(color.FgGreen).SprintfFunc()

var Yellow = color.New(color.FgYellow).SprintfFunc()

func GreenPrint(msg string) {
	Green("%s", msg)
}

func YellowPrint(msg string) {
	Yellow("%s", msg)
}

func RedPrint(msg string) {
	Red("%s", msg)
}

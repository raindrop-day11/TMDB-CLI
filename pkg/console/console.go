package console

import (
	"os"

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

func ExitIf(err error) {
	if err != nil {
		RedPrint(err.Error())
		os.Exit(1)
	}
}

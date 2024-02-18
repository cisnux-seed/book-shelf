package utils

import (
	"os"
	"os/exec"
	"runtime"
)

var clearCommands = map[string]func(){
	"linux": func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	},
	"darwin": func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	},
	"windows": func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	},
}

func ClearScreen() {
	clearCommands[runtime.GOOS]()
}

func Exit() {
	println("Good Bye")
	os.Exit(0)
}

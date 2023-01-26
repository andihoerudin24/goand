package main

import (
	"errors"
	goand "github.com/andihoerudin24/goand"
	"github.com/fatih/color"
	"log"
	"os"
)

const version = "1.0.0"

var cel goand.Goand

func main() {
	arg1, arg2, arg3, err := validateInput()
	if err != nil {
		exitGraceFuly(err)
	}

	switch arg1 {
	case "help":
		showHelp()
	case "version":
		color.Yellow("Application version:" + version)
	default:
		log.Println(arg2, arg3)
	}
}

func validateInput() (string, string, string, error) {
	var arg1, arg2, arg3 string
	if len(os.Args) > 1 {
		arg1 = os.Args[1]

		if len(os.Args) >= 3 {
			arg2 = os.Args[2]
		}

		if len(os.Args) >= 4 {
			arg3 = os.Args[3]
		}
	} else {
		color.Red("Error : command required")
		showHelp()
		return "", "", "", errors.New("command required")
	}

	return arg1, arg2, arg3, nil
}

func showHelp() {
	color.Yellow(`Available commands:
		help    - show the help commands
		version - print application version
	`)
}

func exitGraceFuly(err error, msg ...string) {
	message := ""
	if len(msg) > 0 {
		message = msg[0]
	}

	if err != nil {
		color.Red("Error : %v\n", err)
	}

	if len(message) > 0 {
		color.Yellow(message)
	} else {
		color.Green("Finished!")
	}

	os.Exit(0)
}

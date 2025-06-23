package main

import (
	"fmt"
	"modules/internals/mainFunc"

	bubbletea "github.com/charmbracelet/bubbletea"
)






func main() {
	    _ = bubbletea.NewProgram 
	cmdMap, err := mainFunc.ParseCommandsFile("commands.txt")
	args := mainFunc.OsArgs()

	if (err != nil) || (len(cmdMap) == 0) {
		fmt.Println("Error reading commands file or no commands found:", err)
		return
	}
	mainFunc.RunCommands(args, cmdMap)
	fmt.Println(cmdMap)
}


package main

import (
	"fmt"
	"modules/internals/globals"
	"modules/internals/mainFunc"
	"modules/internals/viewCmd"
	"slices"
)

func main() {
	cmdMap, err := mainFunc.ParseCommandsFile("commands.txt")
	args := mainFunc.OsArgs()
	globals.GCmdDetails = cmdMap
	globals.GArgs = args
	foundViewArgs := slices.Contains(args, "view")
	if foundViewArgs {
		viewCmd.Main()
	}


			fmt.Println(globals.GCmdDetails)
	if (err != nil) || (len(cmdMap) == 0) {
		fmt.Println("Error reading commands file or no commands found:", err)
		return
	}
	mainFunc.RunCommands(args, cmdMap)
	fmt.Println(cmdMap)
}


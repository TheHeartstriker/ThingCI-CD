package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Defines what it takes and returns
func parseCommandsFile(filename string) ([]map[string]string, error) {
    var allLines []map[string]string
    file, err := os.Open(filename)
    if err != nil {
		return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" || strings.HasPrefix(line, "//") {
            continue
        }
        parts := strings.SplitN(line, ":", 2)
        if len(parts) != 2 {
            continue
        }
        name := strings.TrimSpace(parts[0])
        cmds := strings.Split(parts[1], "|")
        cmdCount := 1

        lineMap := make(map[string]string)
        lineMap["name"] = name
        for _, cmd := range cmds {
            c := strings.TrimSpace(cmd)
            if c != "" {
                key := fmt.Sprintf("command%d", cmdCount)
                lineMap[key] = c
                cmdCount++
            }
        }
        allLines = append(allLines, lineMap)
    }
    if err := scanner.Err(); err != nil {
        return nil, err
    }
    return allLines, nil
}

//Givin a name run commands based on map
func runCommands(cmdNames []string, allLines []map[string]string) error {
    repeat := make(map[string]bool)
    for _, cmdName := range cmdNames {
        if repeat[cmdName] {
            continue
        }
        for _, lineMap := range allLines {
            if lineMap["name"] == cmdName {
                for key, command := range lineMap {
                    if strings.HasPrefix(key, "command") {
                        fmt.Println("Running:", command)
                        // Run the command here if needed
                    }
                }
                repeat[cmdName] = true
                break
            }
        }
    }
    return nil
}

func osArgs() []string {
	if len(os.Args) < 2 {
		return []string{}
	}
	return os.Args[1:]
}

func main() {
	cmdMap, err := parseCommandsFile("commands.txt")
	args := osArgs()

	if (err != nil) || (len(cmdMap) == 0) {
		fmt.Println("Error reading commands file or no commands found:", err)
		return
	}
	runCommands(args, cmdMap)
	fmt.Println(cmdMap)
}

// func main() {
//     if len(os.Args) < 2 {
//         fmt.Println("Usage: thing <command>")
//         return
//     }

//     if os.Args[1] == "ls" {
//         var cmd *exec.Cmd
//         if runtime.GOOS == "windows" {
//             cmd = exec.Command("cmd", "/C", "dir")
//         } else {
//             cmd = exec.Command("ls")
//         }
//         output, err := cmd.Output()
//         if err != nil {
//             fmt.Println("Error running command:", err)
//             return
//         }
//         fmt.Print(string(output))
//     } else {
//         fmt.Println("Unknown command:", os.Args[1])
// 	}
// }

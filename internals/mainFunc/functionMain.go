package mainFunc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Defines what it takes and returns
func ParseCommandsFile(filename string) ([]map[string]string, error) {
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
func RunCommands(cmdNames []string, allLines []map[string]string) error {
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

func OsArgs() []string {
	if len(os.Args) < 2 {
		return []string{}
	}
	return os.Args[1:]
}
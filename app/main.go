package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var builtinCommands = map[string]struct{}{
	"echo": {},
	"exit": {},
	"type": {},
}

func main() {
	for {
		fmt.Print("$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("error reading input:" + err.Error())
			os.Exit(1)
		}
		command = strings.TrimSpace(command)
		if command == "exit" {
			os.Exit(0)
		} else if strings.HasPrefix(command, "echo") {
			fmt.Println(strings.TrimSpace(strings.TrimPrefix(command, "echo")))
		} else if strings.HasPrefix(command, "type") {
			arg := strings.TrimSpace(strings.TrimPrefix(command, "type"))
			if _, ok := builtinCommands[arg]; ok {
				fmt.Println(arg + " is a shell builtin")
			} else {
				fmt.Println(arg + ": command not found")
			}
		} else {
			fmt.Println(command + ": command not found")
		}
	}
}

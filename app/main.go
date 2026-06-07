package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var builtinCommands = map[string]struct{}{
	"echo": {},
	"exit": {},
	"type": {},
	"pwd":  {},
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
		} else if command == "pwd" {
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(dir)
		} else if strings.HasPrefix(command, "echo") {
			fmt.Println(strings.TrimSpace(strings.TrimPrefix(command, "echo")))
		} else if strings.HasPrefix(command, "type") {
			arg := strings.TrimSpace(strings.TrimPrefix(command, "type"))
			if _, ok := builtinCommands[arg]; ok {
				fmt.Println(arg + " is a shell builtin")
			} else {
				path, err := exec.LookPath(arg)
				if err != nil {
					fmt.Println(arg + ": not found")
					continue
				}
				fmt.Println(arg + " is " + path)
			}
		} else if strings.HasPrefix(command, "cd") {
			arg := strings.TrimSpace(strings.TrimPrefix(command, "cd"))
			err = os.Chdir(arg)
			if err != nil {
				fmt.Println("cd: "+ arg + ": No such file or directory")
			}
		} else {
			args := strings.Split(command, " ")
			cmd := args[0]
			args = args[1:]
			_, err := exec.LookPath(cmd)
			if err != nil {
				fmt.Println(cmd + ": command not found")
				continue
			}
			cmdStruct := exec.Command(cmd, args...)
			cmdStruct.Stdout = os.Stdout
			cmdStruct.Stderr = os.Stderr
			err = cmdStruct.Run()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

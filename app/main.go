package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
		} else {
			fmt.Println(command + ": command not found")
		}
	}
}

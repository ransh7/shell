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
		switch command {
		case "exit":
			os.Exit(0)
		default:
			fmt.Println(command + ": command not found")
		}
	}
}

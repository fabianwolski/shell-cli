package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var builtIns = []string{"exit", "exit 0", "echo"}

func executeCommand(input string) bool {

	if len(input) == 0 {
		return true
	}

	if input == "exit 0" {
		os.Exit(0)
	}

	var parts = strings.Split(input, " ")
	var cmd = parts[0]
	var argument = parts[1:]

	isBuiltIn := false
	for _, builtIn := range builtIns {
		if cmd == builtIn {
			isBuiltIn = true
			break
		}
	}

	if !isBuiltIn {
		fmt.Printf("%s: command not found\n", cmd)
		return true
	}

	switch cmd {
	case "exit":
		os.Exit(0)
	case "echo":
		if len(argument) > 0 {
			fmt.Println(strings.Join(argument, " "))
		} else {
			fmt.Println()
		}
		return true
	}

	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		command, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		command = strings.TrimSpace(command)

		executeCommand(command)

	}

}

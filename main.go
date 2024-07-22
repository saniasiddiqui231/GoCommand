package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("* ")
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err = execution(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}

func execution(input string) error {
	input = strings.TrimSuffix(input, "\r\n") //strings.TrimRight(input, "\r\n")

	args := strings.Fields(input)
	cmdName := args[0]
	cmdArgs := args[1:]
        switch args[0] {
	case "exit":
    		os.Exit(0)
}
	//fmt.Printf("Executing command: %s %v\n", cmdName, cmdArgs)

	command := exec.Command(cmdName, cmdArgs...)

	command.Stderr = os.Stderr
	command.Stdout = os.Stdout

	err := command.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Command execution failed: %v\n", err)
	}
	//return command.Run()
	return err

}

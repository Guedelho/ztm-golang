//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	nonBlank := 0
	comands := 0

	for {
		input, inputErr := r.ReadString('\n')
		n := strings.TrimSpace(input)
		if n == "" {
			continue
		} else {
			nonBlank += 1
			comands += 1
		}
		if n == "Q" || n == "q" {
			fmt.Println(nonBlank)
			fmt.Println(comands)
			os.Exit(0)
		}
		switch n {
		case "hello":
			fmt.Println("Hello there")
		case "bye":
			fmt.Println("bey there")
		}
		if inputErr == io.EOF {
			break
		}
		if inputErr != nil {
			fmt.Println("Error")
		}
	}

}

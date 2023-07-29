package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define the command-line flags.
	helpFlag := flag.Bool("help", false, "Print this help message")
	flag.Parse()

	// If --help flag is provided, display the help message and exit.
	if *helpFlag {
		displayHelp()
		return
	}

	// Read a string input from the user.
	fmt.Print("Enter a string: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputString := scanner.Text()
		fmt.Println("You entered:", inputString)
		displayHelp()
	} else {
		// Handle any potential error.
		err := scanner.Err()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func displayHelp() {
	fmt.Println("NAME")
	fmt.Println("    my-command - A simple CLI application")
	fmt.Println("")
	fmt.Println("SYNOPSIS")
	fmt.Println("    my-command [OPTIONS]")
	fmt.Println("")
	fmt.Println("OPTIONS")
	fmt.Println("    -h, --help    Print this help message")
	fmt.Println("")
}

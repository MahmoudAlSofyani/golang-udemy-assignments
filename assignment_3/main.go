package main

import (
	"fmt"
	"os"
)

func main() {

	// get command line arguments
	args := os.Args[1:]

	// check if no arguments were passed
	if len(args) == 0 {
		fmt.Println("Please pass in a filename")
		os.Exit(1)
	}

	// read the file
	bs, err := os.ReadFile(args[0])
	checkForError(err)

	// convert the byte slice to a string then log it to the terminal
	fmt.Println(string(bs))

}

// helper function to check if there is an error
func checkForError(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
		os.Exit(1)
	}
}

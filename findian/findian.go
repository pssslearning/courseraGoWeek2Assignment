// --------------------------------------------------------------
// Source code mantained at Github repository for Learning
// https://github.com/pssslearning/courseraGoWeek2Assignment
// --------------------------------------------------------------

package main

import (
	"fmt"
	"os"
	"strings"
)

const dashes = "--------------------------------------------------------------"

func showUSerPrompt() {
	fmt.Println(dashes)
	fmt.Println("Please enter a string to test ... ")
	fmt.Println("      (press <CTRL+C> or 'END' to exit.) ")
	fmt.Println(dashes)
}

func main() {
	// Show Welcome message at start
	fmt.Println("\nWelcome user to 2nd. Assignment program for Week 2,\n")

	var aRuneInInput rune
	var aString string

	for {

		aString = ""

		showUSerPrompt()

		for {
			_, err := fmt.Scanf("%c", &aRuneInInput)

			if err != nil {
				fmt.Println("An error has happenned.")
				fmt.Printf("The error detected was [%s]\n", err)
				fmt.Println("The program will be exited.")
				fmt.Println("Please note that may be remaining input not consumed and sent to your input. ")
				os.Exit(1)
			}

			if aRuneInInput == rune('\n') {
				break
			}

			aString += string(aRuneInInput)
		}

		// Make string entered to UPPER CASE
		aString = strings.ToUpper(aString)

		if aString == "END" {
			// Exit loop
			break
		}

		// Do the test

		if strings.HasPrefix(aString, "I") && strings.Index(aString, "A") > 0 && strings.HasSuffix(aString, "N") {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}

	}

	// Exit with return code 0 to OS.
	fmt.Println("\nGoodbye !!!\n")
	os.Exit(0)
}

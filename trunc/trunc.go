// --------------------------------------------------------------
// Source code mantained at Github repository for Learning
// https://github.com/pssslearning/courseraGoWeek2Assignment
// --------------------------------------------------------------

package main

import (
	"fmt"
	"os"
)

const dashes = "--------------------------------------------------------------"

func showUSerPrompt() {
	fmt.Println(dashes)
	fmt.Println("Please enter a float number to trunc ... ")
	fmt.Println("      (press <CTRL+C> or '0.0' to exit.) ")
	fmt.Println(dashes)
}

func main() {
	// Show Welcome message at start
	fmt.Println("\nWelcome user to Assignment program for Week 2,\n")

	var numFloat float64

	for {

		showUSerPrompt()

		_, err := fmt.Scanf("%f\n", &numFloat)

		if err != nil {
			fmt.Println("An error has happenned. You probably entered something that cannot be interpreted as a float number")
			fmt.Printf("The error detected was [%s]\n", err)
			fmt.Println("The program will be exited.")
			fmt.Println("Please note that may be remaining input not consumed and sent to your input. ")
			os.Exit(1)
		}

		if numFloat == float64(0.0) {
			// Exit loop
			break
		}

		fmt.Printf("The number you entered once correctly parsed as a float number is now: [%f]\n", numFloat)
		fmt.Printf("The number you entered once transformed in integer and then truncated is now: [%d]\n", int64(numFloat))

	}

	// Exit with return code 0 to OS.
	fmt.Println("\nGoodbye !!!\n")
	os.Exit(0)
}

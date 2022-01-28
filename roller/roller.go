// This app simply provides a dice roll.  It's nothing special, but the first go i've written.  I'm sure it will grow in time.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// If we don't get an argument as a dice, then throw an error and advise user
	if len(os.Args) < 2 {
		fmt.Println("Please place a valid die number as your first and only argument")
		os.Exit(1)
	}

	no_of_die := 1

	if len(os.Args) > 2 {
		var err error
		no_of_die, err = strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("There was an issue converting the number of dice", err)
			os.Exit(1)
		}
	}

	// Args is in string format.  Convert to Int using Atoi
	die, err := strconv.Atoi(os.Args[1])

	// If there's some issue in converting the string to an int, let the user know
	if err != nil {
		fmt.Println("Please enter a valid die number", err)
		os.Exit(1)
	}

	for i := 0; i < no_of_die; i++ {
		roll_it(die)
	}
}

func roll_it(die int) (roll int) {
	// provides a moving seed for rand in order to ensure a different result each time.  I might replace this with a hash in the future.
	rand.Seed(time.Now().UnixNano())

	// Provide a random number (starts at 0, hence the +1)
	roll = rand.Intn(die) + 1
	fmt.Println("Your dice roll:", roll)

	return roll
}

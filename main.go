package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rollDice() (rollDice int) {
	setDice := 0
	for {
		setDice++
		die1 := rand.Intn(6) + 1
		if die1 == 1 {
			rollDice = setDice
			return
		}
	}
}

// print winner

// main
func main() {
	fmt.Print("Insert player :")
	var player int
	fmt.Scanln(&player)
	if player < 2 {
		fmt.Print("The players needed is more than 1")
	} else {
		fmt.Print("Insert Round :")
		var rounds int
		fmt.Scanln(&rounds)
		rand.Seed(time.Now().UTC().UnixNano())
		fmt.Printf("%s", "===============")
		for round := 1; round <= rounds; round++ {
			fmt.Printf("\n%s %d %s", "Round", round, ":")
			for i := 1; i <= 1; i++ {
				for j := 1; j <= player; j++ {
					fmt.Printf("\n%s%d%s", "player#", j, ": ")
					for l := 1; l < 5; l++ {
						rollNumber := rollDice()
						fmt.Print(rollNumber, ",")
						for m := 0; m < 4; m++ {
							fmt.Printf("")
						}
					}
				}
			}
			fmt.Printf("\n%s", "After evaluation:")
			// set number player who have dice
			for i := 1; i <= player; i++ {
				fmt.Printf("\n%s%d%s", "player#", i, ": ")
				for j := 1; j < 5; j++ {
					rollNumber := rollDice()
					fmt.Print(rollNumber, ",")
					for k := 0; k < 4; k++ {
						fmt.Printf("")
					}
				}
			}

			fmt.Printf("\n%s", "=============")
		}
	}
	// print high score and winner
	var winner int
	var highScore int
	for i := 1; i <= player; i++ {
		fmt.Printf("\n%s%d%s", "player#", i, ": ")
		for j := 1; j < 5; j++ {
			rollNumber := rollDice()
			fmt.Print(rollNumber, ",")
			for k := 0; k < 4; k++ {
				fmt.Printf("")
			}
			if highScore < rollNumber {
				highScore = rollNumber
				winner = i
			}
		}

	}
	fmt.Printf("\n%s%d", "Winner is : player# ", winner)

}

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	MAX_ATTEMPTS = 6
	MIN          = 1
	MAX          = 100
)

func welcome() {
	fmt.Println("HI LO")
	fmt.Printf("CREATIVE COMPUTING  MORRISTOWN, NEW JERSEY\n\n\n")
	fmt.Printf("THIS IS THE GAME OF HI LO.\n")
	fmt.Printf("YOU WILL HAVE %d TRIES TO GUESS THE AMOUNT OF MONEY IN THE\n", MAX_ATTEMPTS)
	fmt.Printf("HI LO JACKPOT, WHICH IS BETWEEN %d AND %d DOLLARS.  IF YOU\n", MIN, MAX)
	fmt.Println("GUESS THE AMOUNT, YOU WIN ALL THE MONEY IN THE JACKPOT!")
	fmt.Println("THEN YOU GET ANOTHER CHANCE TO WIN MORE MONEY.  HOWEVER,")
	fmt.Printf("IF YOU DO NOT GUESS THE AMOUNT, THE GAME ENDS.\n\n\n")
}

func play() {
	total_winnings := 0
	for {
		secret := randomNumber()
		guessed_correctly := false

		for i := 1; i <= MAX_ATTEMPTS; i++ {

			var in int
			fmt.Printf("YOUR GUESS: ")
			_, err := fmt.Scanf("%d", &in)
			if err != nil {
				panic(err)
			}

			if in == secret {
				fmt.Printf("GOT IT!!!!!!!!!!   YOU WIN %d DOLLARS\n", secret)
				guessed_correctly = true
				break
			} else if in > secret {
				fmt.Printf("YOUR GUESS IS TOO HIGH.\n")
			} else if in < secret {
				fmt.Printf("YOUR GUESS IS TOO LOW.\n")
			}
		}
		if guessed_correctly {
			total_winnings += secret
			fmt.Printf("YOUR TOTAL WINNINGS ARE NOW %d DOLLARS\n", total_winnings)
		} else {
			fmt.Printf("YOU BLEW IT...TOO BAD...THE NUMBER WAS %d\n", secret)
		}

		fmt.Printf("\n")
		fmt.Printf("PLAY AGAIN (YES OR NO)?: ")
		var answer string
		_, err := fmt.Scanf("%s", &answer)
		if err != nil {
			panic(err)
		}
		if strings.ToUpper(answer) != "YES" {
			fmt.Printf("\nSO LONG.  HOPE YOU ENJOYED YOURSELF!!!\n")
			break
		} else {

		}

	}
}

func randomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(MAX-MIN+1) + MIN
}

func main() {
	welcome()
	play()
}

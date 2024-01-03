package exercises

/*
Write a number-guessing game.
Guide the user with  "Try a higher/lower number" warnings
*/
import (
	"fmt"
	"math/rand"
)

func NumberGuessingGame() {

	// Generate secret (random number)
	secret := generateSecret(100)
	// fmt.Println("DEBUG secret:", secret)

	for {
		// Ask the user for their guess (input)
		var guess int
		fmt.Print("Enter an integer: ")
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Check if guess == secret
		if secret == guess {
			fmt.Println("You win!")
			break
		} else if secret < guess {
			fmt.Println("Try a lower number.")
		} else {
			fmt.Println("Try a higher number.")
		}
	}

}

func generateSecret(limit int) int {

	// Since Intn returns a non-negative pseudo-random number in the half-open interval [0,n)
	// we add 1 to create (1, limit)
	return rand.Intn(limit) + 1
}

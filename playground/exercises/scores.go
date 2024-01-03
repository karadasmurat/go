package exercises

import "fmt"

/**
Calculate the max, min and average of n scores.
*/

// read scores from the console.
// func readScores(s *[]int) {
func readScores(s []int) []int {

	score := 0
	for {
		fmt.Print("Enter the score, -1 to quit: ")
		_, err := fmt.Scanf("%d", &score)
		if err != nil {
			fmt.Println("Error reading score:", err)
			return nil
		}

		if score == -1 {
			break
		}

		// v1 - pass a pointer to slice
		// Modifying the original slice through the pointer
		//*s = append(*s, score)

		// v2 - pass slice
		s = append(s, score)

	}

	// v2 the modified slice is returned, or reallocations will not be visible without accepting a pointer to slice instead.
	return s
}

func calculateAverage(numbers []int) float64 {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

func ScoreManager() {
	fmt.Println("Welcome to Score Manager.")

	// Creating a slice with a length of 0 and no specified initial capacity
	scores := make([]int, 0)

	// v1 - pass a pointer to slice
	// readScores(&scores)

	// v2 - pass slice
	scores = readScores(scores)

	fmt.Println(scores)

	avg := calculateAverage(scores)

	fmt.Printf("Average: %.2f\n", avg)
}

package mastermind

import (
	"strings"
)

var Alphabet = []byte{'A', 'B', 'C', 'D', 'E', 'F'}

func Guess(secret, guess string) (positions, letters int) {
	rightPositions := findRightPositions(secret, guess)
	commonLetters := findCommonLetters(secret, guess)

	return rightPositions, commonLetters - rightPositions
}

func findRightPositions(secret, guess string) int {
	count := 0
	for i, _ := range secret {
		if secret[i] == guess[i] {
			count++
		}
	}
	return count
}

func findCommonLetters(secret, guess string) int {
	sum := 0
	for _, letter := range Alphabet {
		secretCount := strings.Count(secret, string(letter))
		guessCount := strings.Count(guess, string(letter))

		if secretCount < guessCount {
			sum += secretCount
		} else {
			sum += guessCount
		}
	}
	return sum
}

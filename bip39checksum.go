package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/tyler-smith/go-bip39"
	"os"
	"strings"
)

const (
	prompt = "Please enter 11 or 23 BIP39 words: "
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(color.BlueString(prompt))

	input, err := reader.ReadString('\n')
	if err != nil {
		color.Red("Error reading input:", err)
		return
	}

	mnemonic := strings.TrimSpace(input)
	nWords := countWords(mnemonic)

	switch nWords {
	case 11, 23:
		possibleWords := brute(mnemonic)
		color.Blue("\nThere are %v last words valid for this mnemonic: \n", len(possibleWords))
		printWords(possibleWords)
	case 12, 24:
		color.Green("You already have the last word!")
	default:
		color.Red("The mnemonic phrase you entered is not valid.")
	}

}

func brute(mnemonic string) []string {
	english := bip39.GetWordList()
	var validWords []string

	for _, word := range english {
		current := mnemonic + " " + word
		isValid := bip39.IsMnemonicValid(current)

		if isValid {
			validWords = append(validWords, word)
		}
	}
	return validWords
}

func printWords(arr []string) {
	for _, w := range arr {
		fmt.Printf("- %s\n", w)
	}
	fmt.Println()

}

func countWords(s string) int {
	words := strings.Fields(s)
	return len(words)
}

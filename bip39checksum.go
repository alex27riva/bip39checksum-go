package main

import (
	"bufio"
	"fmt"
	"github.com/tyler-smith/go-bip39"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter 11 or 23 BIP39 words: ")
	input, _ := reader.ReadString('\n')
	words := brute(input)
	mnemonicLen := len(words)

	if mnemonicLen > 0 {
		fmt.Printf("\nThere are %v last words valid for this mnemonic: \n\n", mnemonicLen)
		printWords(words)
	} else {
		fmt.Println("The mnemonic phrase you entered is not valid.")
	}

}

func brute(mnemonic string) []string {
	english := bip39.GetWordList()
	var valid_words []string

	for _, word := range english {
		current := mnemonic + " " + word
		isValid := bip39.IsMnemonicValid(current)

		if isValid {
			valid_words = append(valid_words, word)
		}
	}
	return valid_words
}

func printWords(arr []string) {
	for _, w := range arr {
		fmt.Printf("- %s\n", w)
	}
	fmt.Println()

}

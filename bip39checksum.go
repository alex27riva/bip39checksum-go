package main

import (
	"bufio"
	"fmt"
	"github.com/tyler-smith/go-bip39"
	"os"
)

// 23 words for testing
// loud omit domain valve topic engine velvet chat foil few wrap unable practice snap gift version brass board broom loud amateur cabin toss

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please insert your mnemonic phrase: ")
	input, _ := reader.ReadString('\n')
	words := brute(input)
	n := len(words)

	if n > 0 {
		fmt.Println("\nPossible last words for this mnemonic: ")
		PrintWords(words)
	} else {
		fmt.Println("Mnemonic invalid")
	}

}

func brute(mnemonic string) []string {
	english := bip39.GetWordList()
	var valid_words []string

	for _, word := range english {
		current := mnemonic + " " + word
		valid := bip39.IsMnemonicValid(current)

		if valid {
			valid_words = append(valid_words, word)
		}
	}
	return valid_words
}

func PrintWords(arr []string) {
	for _, w := range arr {
		fmt.Printf("- %s\n", w)
	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/alex27riva/go-bip39"
	"github.com/fatih/color"
)

const (
	prompt = "Please enter your mnemonic phrase to recover the last word: "
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
	case 11, 14, 17, 20, 23:
		possibleWords := bip39.FindLastWords(mnemonic)
		color.Blue("\nThere are %v last words valid for this mnemonic: \n", len(possibleWords))
		printWords(possibleWords)
	case 12, 15, 18, 21, 24:
		color.Green("You already have the last word!")
	default:
		color.Red("The mnemonic phrase you entered is not valid.")
	}

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

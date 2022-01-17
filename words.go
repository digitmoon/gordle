package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

type placement int8

const (
	correct placement = iota
	wrongplace
	notpresent
	unchecked
)

type letterCell struct {
	ch rune
	pl placement
}

func printLetterCell(l letterCell) {
	green := color.New(color.BgGreen)
	green = green.Add(color.FgBlack, color.Bold)

	yellow := color.New(color.BgYellow)
	yellow = yellow.Add(color.FgBlack, color.Bold)
	white := color.New(color.BgWhite)
	white = white.Add(color.FgBlack, color.Bold)

	black := color.New(color.BgBlack)
	black = black.Add(color.FgWhite, color.Bold)

	switch l.pl {
	case correct:
		green.Printf("%c", l.ch)
		break
	case wrongplace:
		yellow.Printf("%c", l.ch)
		break
	case notpresent:
		white.Printf("%c", l.ch)
		break
	case unchecked:
		black.Printf("%c", l.ch)
	}
}

func checkGuess(word string, guess string) []letterCell {
	wordLength := len(word)
	ret := make([]letterCell, wordLength)
	dupes := make(map[rune]int)
	for _, r := range word {
		_, found := dupes[r]
		if found {
			dupes[r]++
		} else {
			dupes[r] = 1
		}
	}
	for i := 0; i < len(word); i++ {
		g := guess[i]
		if word[i] == g {
			ret[i] = letterCell{rune(g), correct}
			dupes[rune(g)]--
		}
	}
	for i := 0; i < len(word); i++ {
		g := guess[i]
		if word[i] == g {
			continue

		} else if strings.ContainsRune(word, rune(g)) && dupes[rune(g)] > 0 {
			ret[i] = letterCell{rune(g), wrongplace}
			dupes[rune(g)]--
		} else {
			ret[i] = letterCell{rune(g), notpresent}
		}
	}
	return ret

}

func getGuessInput(dict []string) string {
	wordLength := len(dict[0])
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if len(text) != wordLength+1 {
		fmt.Printf("%v not in dictionary\n", text)
		return getGuessInput(dict)
	}
	wordLengthChars := text[0:wordLength]
	if GuessInSortedDict(wordLengthChars, dict) {
		return wordLengthChars
	} else {
		fmt.Printf("%v not in dictionary\n", wordLengthChars)
		return getGuessInput(dict)
	}
}

func uncheckedAlphabet() []letterCell {
	letters := make([]letterCell, 26)
	for i, v := range "abcdefghijklmnopqrstuvwxyz" {
		letters[i] = letterCell{v, unchecked}
	}
	return letters
}

func colorAlphabet(letters []letterCell, alphabet []letterCell) []letterCell {
	for _, v := range letters {
		alphIndex := v.ch - 'a'
		if alphabet[alphIndex].pl == unchecked || (alphabet[alphIndex].pl == wrongplace && v.pl == correct) {
			alphabet[alphIndex].pl = v.pl
		}
	}

	return alphabet
}

func printAlphabet(alphabet []letterCell) {
	for _, v := range alphabet {
		printLetterCell(v)
	}
	fmt.Println()
}

func RunGame(dict []string) {
	word := NewWord(dict)
	alphabet := uncheckedAlphabet()
	printAlphabet(alphabet)
	defer func() { fmt.Println(word) }()
	for i := 0; i < 6; i++ {
		guess := getGuessInput(dict)
		a := checkGuess(word, guess)
		alphabet = colorAlphabet(a, alphabet)
		corr := true
		for _, c := range a {
			printLetterCell(c)
			if c.pl != correct {
				corr = false
			}
		}
		fmt.Println()
		printAlphabet(alphabet)
		if corr {
			fmt.Printf("Correct in %d guesses!\n", i+1)
			break
		}
	}

}

func main() {

	rand.Seed(time.Now().UnixNano())
	dict := GetDict(5)
	RunGame(dict)
}

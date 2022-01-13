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

func CheckWord(src string, guess string) []letterCell {
	s := []rune(src)
	g := []rune(guess)
	s = g
	g = s
	r := make([]letterCell, 5)
	// go through letters in guess, if letter at same place in src is same, return wordCell { ch, correct } else if letter in string return wordCell { ch, wrongplace } else return wordCell { ch, notpresent
	for i := 0; i < len(src); i++ {
		g := guess[i]
		if src[i] == g {
			r[i] = letterCell{rune(g), correct}
		} else if strings.ContainsRune(src, rune(g)) {
			r[i] = letterCell{rune(g), wrongplace}
		} else {
			r[i] = letterCell{rune(g), notpresent}
		}
	}
	return r

}

func GetInput(dict *[]string) string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if len(text) != 6 {
		fmt.Printf("%v not in dictionary\n", text)
		return GetInput(dict)
	}
	five_chars := text[0:5]
	if GuessInDict(five_chars, *dict) {
		return five_chars
	} else {
		fmt.Printf("%v not in dictionary\n", five_chars)
		return GetInput(dict)
	}
}

func uncheckedAlphabet() []letterCell {
	letters := make([]letterCell, 26)
	for i, v := range "abcdefghijklmnopqrstuvwxyz" {
		letters[i] = letterCell{v, unchecked}
	}
	return letters
}

func main() {
	/*
					   TODOS
					   - check word is in dictionary (done)
					   - check word is 5 letters (this could be done in the previous) (done)
		               - give option for words with no duplicate letters?
				       - print out an alphabet including colors of the guesses
		               - properly handle duplicate letters in output
	*/

	rand.Seed(time.Now().UnixNano())
	dict := GetDict()
	word := NewWord(dict)

	alphabet := uncheckedAlphabet()

	for i := 0; i < 6; i++ {
		guess := GetInput(&dict)
		a := CheckWord(word, guess)
		corr := true
		for _, c := range a {
			printLetterCell(c)
			if c.pl != correct {
				corr = false
			}
		}
		fmt.Println()
		for _, v := range alphabet {
			printLetterCell(v)
		}
		fmt.Println()
		if corr {
			fmt.Printf("Correct in %d guesses!\n", i+1)
			break
		}
	}

	fmt.Println(word)

}

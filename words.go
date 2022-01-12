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

func GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	five_chars := text[0:5]
	return five_chars
}

func main() {
	/*
			   TODOS
			   - check word is in dictionary
			   - check word is 5 letters (this could be done in the previous)
		       - print out an alphabet including colors of the guesses
	*/

	rand.Seed(time.Now().UnixNano())
	word := NewWord()

	for i := 0; i < 6; i++ {
		guess := GetInput()
		a := CheckWord(word, guess)
		corr := true
		for _, c := range a {
			printLetterCell(c)
			if c.pl != correct {
				corr = false
			}
		}
		fmt.Println()
		if corr {
			fmt.Printf("Correct in %d guesses!\n", i+1)
			break
		}
	}

	fmt.Println(word)

}

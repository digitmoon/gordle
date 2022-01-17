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
	ret := make([]letterCell, 5)

	dupes := make(map[rune]int)
	for _, r := range src {
		_, found := dupes[r]
		if found {
			dupes[r]++
		} else {
			dupes[r] = 1
		}
	}
	for i := 0; i < len(src); i++ {
		g := guess[i]
		if src[i] == g {
			ret[i] = letterCell{rune(g), correct}
			dupes[rune(g)]--
		}
	}
	for i := 0; i < len(src); i++ {
		g := guess[i]
		if src[i] == g {
			continue

		} else if strings.ContainsRune(src, rune(g)) && dupes[rune(g)] > 0 {
			ret[i] = letterCell{rune(g), wrongplace}
			dupes[rune(g)]--
		} else {
			ret[i] = letterCell{rune(g), notpresent}
		}
	}
	return ret

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

func colorAlphabet(letters []letterCell, alphabet []letterCell) []letterCell {

	for _, v := range letters {
		alphIndex := v.ch - 'a'
		if alphabet[alphIndex].pl == unchecked || (alphabet[alphIndex].pl == wrongplace && v.pl == correct) {
			alphabet[alphIndex].pl = v.pl
		}
	}

	return alphabet
}

func main() {

	rand.Seed(time.Now().UnixNano())
	dict := GetDict()
	word := NewWord(dict)

	alphabet := uncheckedAlphabet()
	defer func() { fmt.Println(word) }()
	for i := 0; i < 6; i++ {
		guess := GetInput(&dict)
		a := CheckWord(word, guess)
		alphabet = colorAlphabet(a, alphabet)
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
}

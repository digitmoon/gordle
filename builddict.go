package main

import (
	"bufio"
	"math/rand"
	"os"
	"regexp"
)

func noDups(s string) bool {
	m := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		r := s[i]
		_, ok := m[r]
		if !ok {
			m[r] = true
		} else {
			return false
		}
	}
	return true
}

func GuessInDict(guess string, dict []string) bool {
	for _, word := range dict {
		if guess == word {
			return true
		} else if word > guess {
			return false
		}
	}
	return false
}

func GetDict() []string {
	fiveLetters := regexp.MustCompile("^[a-z]{5}$")
	f, _ := os.Open("./words.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	words := make([]string, 0)
	for scanner.Scan() {
		word := scanner.Text()
		if fiveLetters.MatchString(word) { //&& noDups(word) {
			words = append(words, word)
		}
	}
	return words
}

func NewWord(dict []string) string {
	n := len(dict)
	i := rand.Intn(n)
	randomWord := dict[i]
	return randomWord
}

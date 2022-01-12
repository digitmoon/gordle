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

func GetDict() []string {
	fiveLetters := regexp.MustCompile("^[a-z]{5}$")
	f, _ := os.Open("./words.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	words := make([]string, 0)
	for scanner.Scan() {
		word := scanner.Text()
		if fiveLetters.MatchString(word) && noDups(word) {
			words = append(words, word)
		}
	}
	return words
}

func NewWord() string {
	//rand.Seed(time.Now().UnixNano())
	d := GetDict()
	n := len(d)
	i := rand.Intn(n)
	randomWord := d[i]
	return randomWord
}

/*

func main() {
	randomWord := NewWord()
	fmt.Println(randomWord)

}

*/

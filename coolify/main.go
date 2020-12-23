// play with vowels in the word to make it cool.
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	duplicateVowel bool = true
	removeVowel    bool = false
)

func randBool() bool {
	return rand.Intn(2) == 0
}

func main() {
	rand.Seed(time.Now().UnixNano())
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		word := []byte(s.Text())
		if randBool() {
			var index = -1
			for i, char := range word {
				switch char {
				case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
					if randBool() {
						index = i
					}
				}
			}
			if index >= 0 {
				switch randBool() {
				case duplicateVowel:
					word = append(word[:index+1], word[index:]...)
				case removeVowel:
					word = append(word[:index], word[index+1:]...)
				}
			}
		}
		fmt.Println(string(word))
	}
}

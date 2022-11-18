/**
A pangram is a sentence where every letter of the English alphabet appears at least once.

Given a string sentence containing only lowercase English letters, return true if sentence is a pangram, or false otherwise.

Example 1:

Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
Output: true
Explanation: sentence contains at least one of every letter of the English alphabet.
Example 2:

Input: sentence = "leetcode"
Output: false

Constraints:

1 <= sentence.length <= 1000
sentence consists of lowercase English letters.
*/

package main

import (
	"fmt"
	"os"
)

type void struct{}

func checkIfPangram(sentence string) bool {
	var alphabetLength = int('z') - int('a') + 1
	var set = make(map[rune]struct{})
	var member void

	var count int
	for _, letter := range sentence {
		if _, letterExists := set[letter]; !letterExists {
			set[letter] = member
			count++
		}
	}

	return count == alphabetLength
}

func main() {
	fmt.Fprintf(os.Stdout, "%v = True\n", checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	fmt.Fprintf(os.Stdout, "%v = False\n", checkIfPangram("leetcode"))
}

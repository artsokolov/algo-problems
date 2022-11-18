/**
Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.

You can use each character in text at most once. Return the maximum number of instances that can be formed.

Example 1:

Input: text = "nlaebolko"
Output: 1

Example 2:

Input: text = "loonbalxballpoon"
Output: 2

Example 3:

Input: text = "leetcode"
Output: 0


Constraints:

1 <= text.length <= 104
text consists of lower case English letters only.
*/
package main

import (
	"fmt"
	"math"
	"os"
)

func maxNumberOfBalloons(text string) int {
	var countLetters = make(map[rune]int)
	for _, letter := range text {
		countLetters[letter]++
	}

	var word = map[rune]int{'b': 1, 'a': 1, 'l': 2, 'o': 2, 'n': 1}
	var ans = math.Inf(1)
	for let, occurrence := range word {
		ans = math.Min(float64(countLetters[let]/occurrence), ans)
	}

	return int(ans)
}

func main() {
	fmt.Fprintf(os.Stdout, "%v\n", maxNumberOfBalloons("leetcode"))
}

/**
Given an integer array nums, return the largest integer that only occurs once. If no integer occurs once, return -1.

Example 1:

Input: nums = [5,7,3,9,4,9,8,3,1]
Output: 8
Explanation: The maximum integer in the array is 9 but it is repeated. The number 8 occurs only once, so it is the answer.

Example 2:

Input: nums = [9,9,8,8]
Output: -1
Explanation: There is no number that occurs only once.


Constraints:

1 <= nums.length <= 2000
0 <= nums[i] <= 1000
*/
package main

import (
	"fmt"
	"math"
	"os"
)

func largestUniqueNumber(nums []int) int {
	var numsOccur = make(map[int]int)
	for _, num := range nums {
		numsOccur[num]++
	}

	max := -1
	for num, occures := range numsOccur {
		if occures == 1 {
			max = int(math.Max(float64(num), float64(max)))
		}
	}

	return max
}

func main() {
	fmt.Fprintf(os.Stdout, "%v\n", largestUniqueNumber([]int{5, 7, 3, 9, 4, 9, 8, 3, 1}))
	fmt.Fprintf(os.Stdout, "%v\n", largestUniqueNumber([]int{9, 9, 8, 8}))
}

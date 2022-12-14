/**
Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

Example 1:

Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].
Example 2:

Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]

Constraints:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums is sorted in non-decreasing order.
*/

package main

import (
	"fmt"
	"math"
	"os"
)

func sortedSquares(nums []int) []int {
	var n = len(nums) - 1
	var ans = make([]int, len(nums))
	i, j := 0, n
	sqr := 0

	for ; n > 0; n-- {
		if math.Abs(float64(nums[i])) > math.Abs(float64(nums[j])) {
			sqr = nums[i]
			i++
		} else {
			sqr = nums[j]
			j--
		}
		ans[n] = sqr * sqr
	}

	return ans
}

func main() {
	k := sortedSquares([]int{-4, -1, 0, 3, 10})
	fmt.Fprintf(os.Stdout, "%v\n", k)
}

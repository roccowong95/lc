/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (29.14%)
 * Likes:    6941
 * Dislikes: 411
 * Total Accepted:    1.2M
 * Total Submissions: 4.1M
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string, find the length of the longest substring without repeating
 * characters.
 *
 *
 * Example 1:
 *
 *
 * Input: "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3.
 * ⁠            Note that the answer must be a substring, "pwke" is a
 * subsequence and not a substring.
 */
package main

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	offset := make(map[byte]int)
	left, right := 0, 1
	current, max := 1, 1
	offset[s[left]] = left

	for right < len(s) && left < len(s) {
		idx, ok := offset[s[right]]
		if ok {
			if left == idx {
				offset[s[right]] = right
				left++
				right++
				continue
			}
			if left < idx && idx < right {
				offset[s[left]] = left
				left++
				continue
			}
		}

		offset[s[left]] = left
		offset[s[right]] = right
		current = right - left + 1
		if current > max {
			max = current
		}
		right++
	}

	return max
}

// @lc code=end

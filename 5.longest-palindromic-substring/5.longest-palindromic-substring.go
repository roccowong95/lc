/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */
func longestPalindrome(s string) string {
	// expand from center
	/*
	length := len(s)
	left, right := 0, 0
	for center := 0; center < length; center++ {
		// odd
		l, r := center, center
		for l >= 0 && r < length && s[l]==s[r] {
			l--
			r++
		}
		// s[l] != s[r] when out of loop

		l++

		if r-l > right-left {
			right, left = r, l
		}

		// even
		l, r = center, center+1
		for l >= 0 && r < length && s[l]==s[r] {
			l--
			r++
		}

		l++

		if r-l > right-left {
			right, left = r, l
		}
	}
	return s[left:right]
	*/

	// dp:
	// i: left most.
	// j: right most.
	// dp[i,j]=1, if i==j
	// dp[i,j]=dp[i+1,j-1], if s[i]==s[i]
	// dp[i,j]=0, if s[i]!=s[j]
}


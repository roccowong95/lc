/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */
func reverse(x int) int {
	var (
		sign = 1
		ret  = 0
	)

	if x == 0 {
		return 0
	}

	if (x+1)>>31 > 0 {
		return 0
	}

	if x < 0 {
		sign = -1
		x = -x
	}

	for ; x > 0; x /= 10 {
		ret = ret*10 + x%10
	}

	if (ret+1)>>31 > 0 {
		return 0
	}

	return sign * ret
}

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */
func reverse(x int) int {
	var (
		sign = 1
		pow = 0
		digits []int
	)

	if x < 0 {
		sign = -1
		x = -x
	}

	for ;x>0;x/=10 {
	}
}


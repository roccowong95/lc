/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 *
 * https://leetcode-cn.com/problems/unique-paths/description/
 *
 * algorithms
 * Medium (57.31%)
 * Likes:    353
 * Dislikes: 0
 * Total Accepted:    52.2K
 * Total Submissions: 91.1K
 * Testcase Example:  '3\n2'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
 *
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
 *
 * 问总共有多少条不同的路径？
 *
 *
 *
 * 例如，上图是一个7 x 3 的网格。有多少可能的路径？
 *
 * 说明：m 和 n 的值均不超过 100。
 *
 * 示例 1:
 *
 * 输入: m = 3, n = 2
 * 输出: 3
 * 解释:
 * 从左上角开始，总共有 3 条路径可以到达右下角。
 * 1. 向右 -> 向右 -> 向下
 * 2. 向右 -> 向下 -> 向右
 * 3. 向下 -> 向右 -> 向右
 *
 *
 * 示例 2:
 *
 * 输入: m = 7, n = 3
 * 输出: 28
 *
 */

package main

import "fmt"

// @lc code=start
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	// factorial O(m+n)
	/*

		all := m + n - 2
		pick := m - 1

		if all-pick < pick {
			pick = all - pick
		}

		//c(pick,all)

		up := 1
		for i := all; i > all-pick; i-- {
			up *= i
		}
		down := 1
		for i := 2; i <= pick; i++ {
			down *= i
		}
		return up / down
	*/

	// dp
	// dp[i, j]=dp[i-1,j]+dp[i,j-1]
	// dp[0,j]=1
	// dp[i,0]=1
	left := 1
	lastRow := make([]int, m)
	currentRow := make([]int, m)

	for i := 0; i < m; i++ {
		lastRow[i] = 1
	}

	current := 0
	for row := 1; row < n; row++ {
		for col := 1; col < m; col++ {
			current = left + lastRow[col]
			left = current
			currentRow[col] = current
		}
		lastRow = currentRow
		currentRow = make([]int, m)
		left = 1
	}
	return current
}

// @lc code=end

func main() {
	fmt.Println(uniquePaths(100, 3))
}

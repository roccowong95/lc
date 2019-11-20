/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if _, ok := m[n]; ok {
			return []int{m[n], i}
		}
		m[target-n] = i
	}
	return nil
}


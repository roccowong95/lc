/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
func threeSum(nums []int) [][]int {
	// 1. brutal
	var result [][]int
	sort.Ints(nums)
	l := len(nums)
	for left:= 0; left<l-2 && nums[left]<=0; left++ {
		if left > 0 && nums[left] == nums[left-1] {
			continue
		}
		for right:=left+1; right<l; right++ {
			if right>left+1 && nums[right] == nums[right-1] {
				continue
			}
			idx := sort.SearchInts(nums[right+1:], -nums[left]-nums[right])
			if right+idx+1>=l {
				continue
	 		} else {
				 if nums[left]+nums[right]+nums[right+idx+1] == 0 {
					result = append(result, []int{nums[left], nums[right], nums[right+idx+1]})
				 }
			}
		}
	}
	return result
}


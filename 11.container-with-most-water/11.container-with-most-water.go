/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
func maxArea(height []int) int {
	// 1. brutal
	/*
	l := len(height)
	max := 0
	for width := 1; width < l; width++ {
		widthMax := 0
		for idx := 0; idx+width<l; idx++ {
			min := height[idx]
			if height[idx+width] < min {
				min = height[idx+width]
			}
			if min > widthMax {
				widthMax = min
			}
		}
		area := width * widthMax
		if area > max {
			max = area
		}
	}
	return max
	*/

	// 2. improved
	// proof:
	// assumptions:
	// a. max area lies on i, j
	// b. h[i]<h[j] (identical)
	// then max_area=(j-i)*h[j]-h[i].
	//
	// it can be then inferred that:
	// c. for any start<i, h[start]<h[i]
	// d. for any end>j, h[end]<h[i]
	//
	// so we should proove: the situation where start=i and end=j
	// would happen during the algorithm.
	//
	// in this algorithm, start and end moves 1 index at a time,
	// so either will eventually happen:
	// 1.  start=i, end>j
	// 2.  end=j, start<i
	//
	// if 1 happens, according to assumption b d, 
	// h[end]<h[i]=h[start]<h[j]
	// so until end==j, start will not move, which means we can 
	// come across where start==i and end==j.
	//
	// if 2 happens, h[start]<h[i]<h[j]=h[end]
	// until start==i, end will not move, so the conclusion is same.
	max, start, end := 0, 0, len(height) - 1
	for start<end {
		min := height[start]
		if height[end] < min {
			min = height[end]
		}
		current := min * (end-start)
		if current > max {
			max = current
		}
		if height[start] > height[end] {
			end--
		} else {
			start++
		}
	}
	return max
}


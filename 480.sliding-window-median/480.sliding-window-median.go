/*
 * @lc app=leetcode id=480 lang=golang
 *
 * [480] Sliding Window Median
 *
 * https://leetcode.com/problems/sliding-window-median/description/
 *
 * algorithms
 * Hard (34.32%)
 * Likes:    542
 * Dislikes: 56
 * Total Accepted:    34.1K
 * Total Submissions: 99.3K
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * Median is the middle value in an ordered integer list. If the size of the
 * list is even, there is no middle value. So the median is the mean of the two
 * middle value.
 * Examples:
 * [2,3,4] , the median is 3
 * [2,3], the median is (2 + 3) / 2 = 2.5
 *
 * Given an array nums, there is a sliding window of size k which is moving
 * from the very left of the array to the very right. You can only see the k
 * numbers in the window. Each time the sliding window moves right by one
 * position. Your job is to output the median array for each window in the
 * original array.
 *
 * For example,
 * Given nums = [1,3,-1,-3,5,3,6,7], and k = 3.
 *
 *
 * Window position                Median
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       1
 * ⁠1 [3  -1  -3] 5  3  6  7       -1
 * ⁠1  3 [-1  -3  5] 3  6  7       -1
 * ⁠1  3  -1 [-3  5  3] 6  7       3
 * ⁠1  3  -1  -3 [5  3  6] 7       5
 * ⁠1  3  -1  -3  5 [3  6  7]      6
 *
 *
 * Therefore, return the median sliding window as [1,-1,-1,3,5,6].
 *
 * Note:
 * You may assume k is always valid, ie: k is always smaller than input array's
 * size for non-empty array.
 */
package main

import "fmt"

// @lc code=start

// TODO change to interface

type heap struct {
	arr []int
	lt  func(int, int) bool // "minimum" heap with lt, lt(a,b)=true means a"<"b
}

func newHeap(arr []int) *heap {
	return &heap{arr: arr}
}

func left(idx int) int {
	return 2*idx + 1
}

func right(idx int) int {
	return 2*idx + 2
}

func parent(idx int) int {
	return (idx - 1) / 2
}

// l: max index
func (h *heap) siftDown(idx int, l int) {
	l++
	for idx < l {
		lidx := left(idx)
		if lidx >= l {
			return
		}
		smaller := lidx
		ridx := right(idx)
		if ridx < l {
			if h.lt(h.arr[ridx], h.arr[smaller]) {
				smaller = ridx
			}
		}
		if !h.lt(h.arr[idx], h.arr[smaller]) {
			//fmt.Printf("swapping [%d]:%d with [%d]:%d\n", idx, h.arr[idx], smaller, h.arr[smaller])
			h.arr[idx], h.arr[smaller] = h.arr[smaller], h.arr[idx]
		}
		//fmt.Printf("after %v\n", h.arr)
		idx = smaller
	}
}

func (h *heap) siftUp(idx int) {
	for idx > 0 {
	}
}

func (h *heap) sort() {
	idx := len(h.arr) - 1
	for idx >= 0 {
		h.arr[0], h.arr[idx] = h.arr[idx], h.arr[0]
		h.siftDown(0, idx-1)
		idx--
	}
}

func (h *heap) heapify() {
	idx := parent(len(h.arr) - 1)
	for idx >= 0 {
		h.siftDown(idx, len(h.arr)-1)
		idx--
	}
}

func medianSlidingWindow(nums []int, k int) []float64 {
	var ret []float64
	return ret
}

// @lc code=end

func less(a, b int) bool {
	return a < b
}

func main() {
	//fmt.Println(medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	arr := []int{3, 2, 1, 0, -1, -2, -3, -4}
	fmt.Printf("original \t%v\n", arr)
	h := newHeap(arr)
	h.lt = less
	h.heapify()
	fmt.Printf("heapified\t%v\n", h.arr)
	h.sort()
	fmt.Printf("sorted   \t%v\n", h.arr)
}

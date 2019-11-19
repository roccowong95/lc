/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (59.45%)
 * Total Accepted:    568.2K
 * Total Submissions: 955.7K
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given a binary tree, return the inorder traversal of its nodes' values.
 *
 * Example:
 *
 *
 * Input: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * Output: [1,3,2]
 *
 * Follow up: Recursive solution is trivial, could you do it iteratively?
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func inorderTraversal(root *TreeNode) []int {
	// 1. recursive.
	/*
		var ret []int
		if nil == root {
			return ret
		}
		if nil != root.Left {
			ret = append(ret, inorderTraversal(root.Left)...)
		}
		ret = append(ret, root.Val)
		if nil != root.Right {
			ret = append(ret, inorderTraversal(root.Right)...)
		}
		return ret
	*/

	// 2. non-recursive.
	var stack []*TreeNode
	var l int
	var ret []int
	ptr := root

	for ptr != nil || l != 0 {
		for ptr != nil {
			stack = append(stack, ptr)
			l++
			ptr = ptr.Left
		}
		ptr = stack[l-1]
		ret = append(ret, ptr.Val)
		stack = stack[:l-1]
		l--
		ptr = ptr.Right
	}

	return ret
}

// @lc code=end

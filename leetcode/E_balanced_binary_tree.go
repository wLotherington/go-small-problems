// Easy
// https://leetcode.com/problems/balanced-binary-tree/

import "math"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	} else if root.Right == nil && root.Left == nil {
		return true
	} else if root.Right == nil {
		return isBalanced(root.Left) && height(root.Left) <= 1
	} else if root.Left == nil {
		return isBalanced(root.Right) && height(root.Right) <= 1
	}

	return isBalanced(root.Right) && isBalanced(root.Left) && heightDiff(root.Right, root.Left) <= 1
}

func height(node *TreeNode) int {
	if node.Right == nil && node.Left == nil {
		return 1
	} else if node.Right == nil {
		return 1 + height(node.Left)
	} else if node.Left == nil {
		return 1 + height(node.Right)
	}

	return 1 + heightMax(node.Right, node.Left)
}

func heightDiff(node1, node2 *TreeNode) int {
	return int(math.Abs(float64(height(node1) - height(node2))))
}

func heightMax(node1, node2 *TreeNode) int {
	return int(math.Max(float64(height(node1)), float64(height(node2))))
}
package code

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var dfs func(node *TreeNode, vals *[]int)
	dfs = func(root *TreeNode, vals *[]int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			*vals = append(*vals, root.Val)
		}
		dfs(root.Left, vals)
		dfs(root.Right, vals)
	}

	vals1 := []int{}
	vals2 := []int{}
	dfs(root1, &vals1)
	dfs(root2, &vals2)

	if len(vals1) != len(vals2) {
		return false
	}

	for k, v := range vals1 {
		if v != vals2[k] {
			return false
		}
	}
	return true
}

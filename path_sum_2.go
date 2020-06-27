package main

// TreeNode is the definition of a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {

	allPaths := [][]int{}
	paths(root, 0, sum, []int{}, &allPaths)
	return allPaths
}

func paths(root *TreeNode, partialSum, targetSum int, nodes []int, allPaths *[][]int) {
	if root == nil {
		return
	}

	partialSum += root.Val

	nodes = append(nodes, root.Val)

	// Leaf node
	if root.Right == nil && root.Left == nil {
		if partialSum == targetSum {
			*allPaths = append(*allPaths, append([]int{}, nodes...))
		}
		return
	}

	paths(root.Left, partialSum, targetSum, nodes, allPaths)
	paths(root.Right, partialSum, targetSum, nodes, allPaths)

	return
}

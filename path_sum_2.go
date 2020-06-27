package main

import "fmt"

/*
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pprint(ll *[][]int) {
	for i := range *ll {
		fmt.Printf("Node: %v, %p, %p, %v \n", i, ll, (*ll)[i], (*ll)[i])
	}
}

func pathSum(root *TreeNode, sum int) [][]int {

	return paths(root, 0, sum, []int{}, [][]int{})
}

func paths(root *TreeNode, partialSum, targetSum int, nodes []int, allPaths [][]int) [][]int {
	if root == nil {
		return allPaths
	}

	// fmt.Printf("root: %v, partial: %v, target: %v, nodes: %v, allPaths 1: %+v\n", root.Val, partialSum, targetSum, nodes, allPaths)
	partialSum += root.Val
	// fmt.Printf("root: %v, partial: %v, target: %v, nodes: %v, allPaths 2: %#v\n", root.Val, partialSum, targetSum, nodes, allPaths)

	fmt.Printf("before append nodes %p %v\n", &nodes, nodes)
	pprint(&allPaths)
	nodes = append(nodes, root.Val)
	fmt.Printf("after append %p, %v \n", &nodes, nodes)
	pprint(&allPaths)

	// fmt.Printf("root: %v, partial: %v, target: %v, nodes: %v, allPaths 3: %#v\n", root.Val, partialSum, targetSum, nodes, allPaths)

	// Leaf node
	if root.Right == nil && root.Left == nil {
		if partialSum == targetSum {
			fmt.Println("SUM MATCH BEGIN")
			fmt.Println(nodes)
			fmt.Println("SUM MATCH END")
			allPaths = append(allPaths, nodes)
		}
		return allPaths
	}

	allPaths = paths(root.Left, partialSum, targetSum, nodes, allPaths)
	fmt.Printf("allPaths after left %v\n", allPaths)
	allPaths = paths(root.Right, partialSum, targetSum, nodes, allPaths)
	fmt.Printf("allPaths after right %v\n", allPaths)

	return allPaths
}

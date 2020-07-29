package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {

}

func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}

	if root.Left == nil && root.Right == nil{
		return 1
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return 1 + leftDepth
	}else {
		return 1 + rightDepth
	}
}

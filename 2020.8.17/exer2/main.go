package main

type TreeNode struct {
	Val int
	Right *TreeNode
	Left *TreeNode
}

/*
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。
*/

func main()  {

}

func isBalanced(root *TreeNode) bool {
	if root == nil{
		return true
	}

	if !isBalanced(root.Left) || !isBalanced(root.Right){
		return false
	}

	leftH := maxDepth(root.Left) + 1
	rightH := maxDepth(root.Right) + 1
	if abs(leftH - rightH) > 1{
		return false
	}

	return true
}

func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b{
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0{
		return -a
	}

	return a
}
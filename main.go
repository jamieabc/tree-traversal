package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var root *TreeNode

func init() {
	//      1
	//   2      3
	//  4 5    6 7

	n7 := &TreeNode{Val: 7}
	n6 := &TreeNode{Val: 6}
	n5 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 4}

	n3 := &TreeNode{
		Val:   3,
		Left:  n6,
		Right: n7,
	}

	n2 := &TreeNode{
		Val:   2,
		Left:  n4,
		Right: n5,
	}

	root = &TreeNode{
		Val:   1,
		Left:  n2,
		Right: n3,
	}
}

//   N
// L   R
func main() {
	fmt.Println("in-order:")
	inOrder(root)
	fmt.Println("")
	fmt.Println("pre-order:")
	preOrder(root)
	fmt.Println("")
	fmt.Println("post-order:")
	postOrder(root)
	fmt.Println("")
	fmt.Println("same-level:")
	sameLevel(root)
}

func inOrder(node *TreeNode) {
	// L - N - R
	if node == nil {
		return
	}

	if node.Left != nil {
		inOrder(node.Left)
	}

	fmt.Printf(" %d ", node.Val)

	if node.Right != nil {
		inOrder(node.Right)
	}
}

func preOrder(node *TreeNode) {
	// N - L - R
	if node == nil {
		return
	}

	fmt.Printf(" %d ", node.Val)

	if node.Left != nil {
		preOrder(node.Left)
	}

	if node.Right != nil {
		preOrder(node.Right)
	}
}

func postOrder(node *TreeNode) {
	// L - R - N
	if node == nil {
		return
	}

	if node.Left != nil {
		postOrder(node.Left)
	}

	if node.Right != nil {
		postOrder(node.Right)
	}

	fmt.Printf(" %d ", node.Val)
}

func sameLevel(node *TreeNode) {
	length := 7
	nodes := make([]*TreeNode, length)
	nodes[0] = node

	// no queue, use slice instead
	for i, j := 0, 1; i < length; i++ {
		n := nodes[i]
		if n.Left != nil {
			nodes[j] = n.Left
			j++
		}

		if n.Right != nil {
			nodes[j] = n.Right
			j++
		}
	}

	for _, n := range nodes {
		fmt.Printf(" %d ", n.Val)
	}
}

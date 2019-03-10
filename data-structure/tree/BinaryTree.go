package tree

import (
	"fmt"
)

// BinaryTree 二叉树
type BinaryTree struct {
	root *Node
}

// NewBinaryTree 构建二叉树
func NewBinaryTree(rootV interface{}) *BinaryTree {
	return &BinaryTree{NewNode(rootV)}
}

// PreOrderTraverse 前序遍历 中左右
func (b *BinaryTree) PreOrderTraverse(n *Node) {
	fmt.Printf("%+v ", n.data)
	if n.left != nil {
		b.PreOrderTraverse(n.left)
	}
	if n.right != nil {
		b.PreOrderTraverse(n.right)
	}
}

// InOrderTraverse 中序遍历，递归写法，左中右
func (b *BinaryTree) InOrderTraverse(n *Node) {
	if n.left != nil {
		b.InOrderTraverse(n.left)
	}
	fmt.Printf("%+v ", n.data)
	if n.right != nil {
		b.InOrderTraverse(n.right)
	}
}

// PostOrderTraverse 后序遍历，递归写法，左右中
func (b *BinaryTree) PostOrderTraverse(n *Node) {
	if n.left != nil {
		b.PostOrderTraverse(n.left)
	}
	if n.right != nil {
		b.PostOrderTraverse(n.right)
	}
	fmt.Printf("%+v ", n.data)
}

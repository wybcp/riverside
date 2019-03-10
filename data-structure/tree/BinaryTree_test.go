package tree

import "testing"

func GetBinaryTree() *BinaryTree {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.left = NewNode(3)
	binaryTree.root.right = NewNode(4)
	binaryTree.root.right.left = NewNode(5)
	binaryTree.root.right.right = NewNode(6)
	return binaryTree
}
var binaryTree = GetBinaryTree()
func TestBinaryTree_InOrderTraverse(t *testing.T) {
	binaryTree.InOrderTraverse(binaryTree.root)
}

func TestBinaryTree_PreOrderTraverse(t *testing.T) {
	binaryTree.PreOrderTraverse(binaryTree.root)
}
func TestBinaryTree_PostOrderTraverse(t *testing.T) {
	binaryTree.PostOrderTraverse(binaryTree.root)
}

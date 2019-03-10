package tree

import "testing"

var compareFunc = func(v, nodeV interface{}) int {
	vInt := v.(int)
	nodeVInt := nodeV.(int)

	if vInt > nodeVInt {
		return 1
	} else if vInt < nodeVInt {
		return -1
	}
	return 0
}

func GetBST() *BST {
	bst := NewBST(6, compareFunc)

	bst.Insert(9)
	bst.Insert(4)
	bst.Insert(8)
	bst.Insert(10)
	bst.Insert(2)
	bst.Insert(5)
	bst.Insert(12)
	return bst
}

var bst = GetBST()

func TestBST_Find(t *testing.T) {
	t.Log(bst.Find(9))
	// 由小到大，一次输出
	bst.InOrderTraverse(bst.BinaryTree.root)
}

func TestBST_Insert(t *testing.T) {
	bst.InOrderTraverse(bst.BinaryTree.root)
}

func TestBST_Min(t *testing.T) {
	t.Log(bst.Min())
}

func TestBST_Max(t *testing.T) {
	t.Log(bst.Max())
}

func TestBST_DeleteA(t *testing.T) {
	// 不存在的节点
	t.Log(bst.Delete(7))
	bst.InOrderTraverse(bst.BinaryTree.root)
}

func TestBST_DeleteB(t *testing.T) {
	// 删除叶子节点
	t.Log(bst.Delete(8))

	bst.InOrderTraverse(bst.BinaryTree.root)
}

func TestBST_DeleteC(t *testing.T) {
	// 删除有一个子节点的节点
	t.Log(bst.Delete(10))

	bst.InOrderTraverse(bst.BinaryTree.root)
}

func TestBST_DeleteD(t *testing.T) {
	// 删除有两个子节点的节点
	t.Log(bst.Delete(9))

	bst.InOrderTraverse(bst.BinaryTree.root)
}

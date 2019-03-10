package tree

// BST 二叉查找树，在树中的任意一个节点，其左子树中的每个节点的值，都要小于这个节点的值，而右子树节点的值都大于这个节点的值。
type BST struct {
	*BinaryTree
	//比对函数，0:v==nodeV,正数:v>nodeV,负数:v<nodeV
	compare func(v, nodeV interface{}) int
}

// NewBST 新建二叉查找树
func NewBST(rootV interface{}, compare func(v, nodeV interface{}) int) *BST {
	if nil == compare {
		return nil
	}
	return &BST{BinaryTree: NewBinaryTree(rootV), compare: compare}
}

// Find 二叉查找树查找
func (b *BST) Find(v interface{}) *Node {
	n := b.root
	for n != nil {
		compareResult := b.compare(v, n.data)
		if compareResult == 0 {
			return n
		} else if compareResult > 0 {
			//v > nodeV
			n = n.right
		} else {
			//v < nodeV
			n = n.left
		}
	}
	return nil
}

// Insert 二叉查找树插入,如果插入的值已存在，则失败
func (b *BST) Insert(v interface{}) bool {
	n := b.root
	for n != nil {
		compareResult := b.compare(v, n.data)
		if compareResult == 0 {
			return false
		} else if compareResult > 0 {
			if n.right == nil {
				n.right = NewNode(v)
				break
			}
			n = n.right
		} else {
			if n.left == nil {
				n.left = NewNode(v)
				break
			}
			n = n.left
		}
	}
	return true
}

// Delete 二叉查找树删除
func (b *BST) Delete(v interface{}) bool {

	// 删除的节点的父节点
	var fatherN *Node
	deleteLeft := false
	n := b.root
	for n != nil {
		compareResult := b.compare(v, n.data)
		if compareResult == 0 {
			break
		} else if compareResult > 0 {
			//v > nodeV
			fatherN = n
			deleteLeft = false
			n = n.right
		} else {
			//v < nodeV
			fatherN = n
			deleteLeft = true
			n = n.left
		}
	}
	//需要删除的节点不存在
	if n == nil {
		return false
	}

	switch {
	//删除的是一个叶子节点
	case n.left == nil && n.right == nil:
		if fatherN != nil {
			if deleteLeft {
				fatherN.left = nil
			} else {
				fatherN.right = nil
			}
		} else { //根节点
			b.root = nil
		}
	case n.right != nil && n.left != nil:
		//删除的是一个有左右节点的节点
		// 找到这个节点的右子树中的最小节点，把它替换到要删除的节点上。
		// 然后再删除掉这个最小节点，因为最小节点肯定没有左子节点（如果有左子结点，那就不是最小节点了），可以应用上面两条规则来删除这个最小节点。

		// 节点的右子树的根节点
		rightN := n.right
		rightNFather := n
		for rightN.left != nil {
			rightNFather = rightN
			rightN = rightN.right
		}
		n.data = rightN.data
		if rightN.right == nil {
			rightNFather.right = nil
		} else {
			rightNFather.right = rightN.right
		}
	default:
		//删除的是只有一个节点的节点
		if n.left == nil {
			if deleteLeft {
				fatherN.left = n.right
			} else {
				fatherN.right = n.right
			}
		} else {
			if deleteLeft {
				fatherN.left = n.left
			} else {
				fatherN.right = n.left
			}
		}
	}

	return true
}

// Min 二叉查找树最小值
func (b *BST) Min() *Node {
	n := b.root
	for n.left != nil {
		n = n.left
	}
	return n
}

// Max 二叉查找树最大值
func (b *BST) Max() *Node {
	n := b.root
	for n.right != nil {
		n = n.right
	}
	return n
}

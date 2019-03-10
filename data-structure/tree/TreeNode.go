package tree

import "fmt"

// Node 树节点
type Node struct {
	data  interface{}
	left  *Node
	right *Node
}

// NewNode 创建新节点
func NewNode(data interface{}) *Node {
	return &Node{data: data}
}

//String 字符串输出
func (n *Node) String() string {
	return fmt.Sprintf("v:%+v, left:%+v, right:%+v", n.data, n.left, n.right)
}

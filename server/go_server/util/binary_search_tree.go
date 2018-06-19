package util

import (
	"errors"
	"fmt"
)

// BinarySearchTree 树
type BinarySearchTree struct {
	root *Node
}

// Insert 插入
func (bst *BinarySearchTree) Insert(key string, value interface{}) error {
	node := &Node{key, value, nil, nil}
	if bst.root == nil {
		bst.root = node
		return nil
	}
	return bst.root.Insert(node)
}

// InOrderTraverse 中序遍历
func (bst *BinarySearchTree) InOrderTraverse(f func(*Node)) {
	bst.root.InOrderTraverse(f)
}

// String 打印
func (bst *BinarySearchTree) String() {
	sep := "------------------------------------------------"
	fmt.Println(sep)
	bst.root.Stringify(0)
	fmt.Println(sep)
}

// Search 查找
func (bst *BinarySearchTree) Search(key string) (value interface{}, ok bool) {
	return bst.root.Search(key)
}

// Remove 删除
func (bst *BinarySearchTree) Remove(key string) {
	bst.root = bst.root.Remove(key)
}

// Node 节点
type Node struct {
	key   string
	value interface{}
	left  *Node
	right *Node
}

// Insert 插入
func (n *Node) Insert(c *Node) error {
	if n == nil {
		return errors.New("node could not be nil")
	}

	switch {
	case n.key == c.key: // 已存在该key，忽略它
		return nil
	case n.key > c.key:
		if n.left == nil { // 左节点为空
			n.left = c
			return nil
		}
		return n.left.Insert(c)
	case n.key < c.key:
		if n.right == nil { // 右节点为空
			n.right = c
			return nil
		}
		return n.right.Insert(c)
	}
	return nil
}

// InOrderTraverse 中序遍历
func (n *Node) InOrderTraverse(f func(*Node)) {
	if n != nil {
		n.left.InOrderTraverse(f)
		f(n)
		n.right.InOrderTraverse(f)
	}
}

// Stringify 打印
func (n *Node) Stringify(level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		n.left.Stringify(level)
		fmt.Printf(format+"%v(%v)\n", n.key, n.value)
		n.right.Stringify(level)
	}
}

// Search 查找
func (n *Node) Search(key string) (value interface{}, ok bool) {
	if n == nil {
		return
	}
	// 二分法查找
	switch {
	case n.key == key:
		return n.value, true
	case n.key > key:
		return n.left.Search(key)
	case n.key < key:
		return n.right.Search(key)
	}
	return
}

// Remove 删除
func (n *Node) Remove(key string) *Node {
	if n == nil {
		return nil
	}
	// 不匹配当前节点，往子树递归
	if n.key > key {
		n.left = n.left.Remove(key)
		return n
	}
	if n.key < key {
		n.right = n.right.Remove(key)
		return n
	}
	// 匹配当前节点，分三种情况
	if n.left == nil && n.right == nil { // 1 叶子节点
		n = nil
		return nil
	}
	if n.left == nil { // 2(1) 左子节点为空
		n = n.right
		return n
	}
	if n.right == nil { // 2(2) 右子节点为空
		n = n.left
		return n
	}

	// 3 左右子节点均存在，找出右子树最小值/左子树最大值
	minNodeInRight := n.right // 在右子树找出最小值，将最小值放到被删除节点的位置上
	for {
		// 最小值肯定是在左节点
		if minNodeInRight != nil && minNodeInRight.left != nil {
			minNodeInRight = minNodeInRight.left
		} else {
			break
		}
	}
	n.key, n.value = minNodeInRight.key, minNodeInRight.value
	n.right = n.right.Remove(n.key) // 这个key所在的Node已经被移到n的位置上，所以要将它从子树里删除

	return n
}

func (n *Node) rotateLeft() {
	// TODO
}

func (n *Node) rotateRight() {
	// TODO
}

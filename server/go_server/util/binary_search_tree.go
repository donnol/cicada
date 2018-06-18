package util

// BinarySearchTree 二叉搜索树
type BinarySearchTree struct {
	root *Node
	size int
}

// Node 树节点
type Node struct {
	parent *Node
	left   *Node
	right  *Node
	value  string
}

// Search 搜索
func (bst *BinarySearchTree) Search(value string) (ok bool) {
	node := bst.root
	for node != nil {
		if node.value == value {
			ok = true
			return
		} else if node.value > value {
			node = node.left
		} else {
			node = node.right
		}
	}
	return
}

// Insert 寻找合适的位置插入
func (bst *BinarySearchTree) Insert(value string) (err error) {
	node := bst.root
	for node != nil {
	}
	return
}

// Remove 删除
func (bst *BinarySearchTree) Remove(value string) (ok bool) {
	node := bst.root
	for node != nil {
		if node.value == value {
			// TODO
			if node.parent.value > node.value { // node在左边
				// node.parent.left = node.right
				// node.right.parent = node.parent
			} else { // node在右边
				// node.parent.right = node.left
				// node.left.parent = node.parent
			}
			node.parent = nil
			node.left = nil
			node.right = nil
			ok = true
			return
		} else if node.value > value {
			node = node.left
		} else {
			node = node.right
		}
	}
	return
}

func (n *Node) rotateLeft() {
	// TODO
}

func (n *Node) rotateRight() {
	// TODO
}

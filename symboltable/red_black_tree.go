package symboltable

type rbtNode struct {
	red   bool
	left  *rbtNode
	right *rbtNode
	key   Comparable
	value interface{}
}

type redBlackTreeSt struct {
	root *rbtNode
	size int
}

func NewRedBlackTreeSt() SymbolTable{
	return &redBlackTreeSt{}
}

func (rb *redBlackTreeSt) Get(key Comparable) interface{} {
	return rb.get(key, rb.root)
}

func (rb *redBlackTreeSt) get(key Comparable, node *rbtNode) interface{} {
	if node == nil {
		return nil
	}

	compareResult := node.key.CompareTo(key)
	if compareResult > 0 {
		return rb.get(key, node.left)
	} else if compareResult < 0 {
		return rb.get(key, node.right)
	} else {
		return node.value
	}
}

func (rb *redBlackTreeSt) Set(key Comparable, value interface{}) {
	rb.root = rb.set(rb.root, key, value)
	rb.root.red = false
}

func (rb *redBlackTreeSt) set(node *rbtNode, key Comparable, value interface{}) *rbtNode {
	if node == nil {
		rb.size++
		return &rbtNode{
			red:   true,
			key:   key,
			value: value,
		}
	}

	compareResult := node.key.CompareTo(key)
	if compareResult > 0 {
		node.left = rb.set(node.left, key, value)
	} else if compareResult < 0 {
		node.right = rb.set(node.right, key, value)
	} else {
		node.value = value
		return node
	}

	if (node.left == nil || !node.left.red) && node.right != nil && node.right.red {
		return rb.leftRotate(node)
	} else if node.left != nil && node.left.red && node.left.left != nil && node.left.left.red {
		return rb.flipDoubleRed(rb.rightRotate(node))
	} else if node.left != nil && node.left.red && node.right != nil && node.right.red {
		return rb.flipDoubleRed(node)
	}

	return node
}

func (rb *redBlackTreeSt) leftRotate(node *rbtNode) *rbtNode {
	right := node.right
	node.right = right.left
	right.left = node
	right.red = node.red
	node.red = true
	return right
}

func (rb *redBlackTreeSt) rightRotate(node *rbtNode) *rbtNode {
	left := node.left
	node.left = left.right
	left.right = node
	left.red = node.red
	node.red = true
	return left
}

func (rb *redBlackTreeSt) flipDoubleRed(node *rbtNode) *rbtNode {
	node.left.red = false
	node.right.red = false
	node.red = true
	return node
}

func (rb *redBlackTreeSt) Size() int {
	return rb.size
}

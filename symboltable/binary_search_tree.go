package symboltable

import "log"

type binarySearchTreeNode struct {
	left  *binarySearchTreeNode
	right *binarySearchTreeNode
	key   Comparable
	value interface{}
}

type binarySearchTreeSt struct {
	root *binarySearchTreeNode
	size int
}

func NewBinarySearchTreeSt() SymbolTable {
	return &binarySearchTreeSt{root: nil}
}

func (bst *binarySearchTreeSt) Get(key Comparable) interface{} {
	return bst.get(key, bst.root)
}

func (bst *binarySearchTreeSt) get(key Comparable, fromNode *binarySearchTreeNode) interface{} {
	if fromNode == nil {
		return nil
	}

	compareResult := fromNode.key.CompareTo(key)
	if compareResult > 0 {
		return bst.get(key, fromNode.left)
	} else if compareResult < 0 {
		return bst.get(key, fromNode.right)
	} else {
		return fromNode.value
	}
}

func (bst *binarySearchTreeSt) Set(key Comparable, value interface{}) {
	if bst.root == nil {
		bst.root = &binarySearchTreeNode{
			key:   key,
			value: value,
		}
		bst.size++
		return
	}

	currentNode := bst.root
	for {
		compareResult := currentNode.key.CompareTo(key)
		if compareResult > 0 {
			if currentNode.left == nil {
				currentNode.left = &binarySearchTreeNode{
					key:   key,
					value: value,
				}
				bst.size++
				return
			} else {
				currentNode = currentNode.left
				continue
			}
		} else if compareResult < 0 {
			if currentNode.right == nil {
				currentNode.right = &binarySearchTreeNode{
					key:   key,
					value: value,
				}
				bst.size++
				break
			} else {
				currentNode = currentNode.right
				continue
			}
		} else {
			currentNode.value = value
			break
		}
	}
}

func (bst *binarySearchTreeSt) printTree(node *binarySearchTreeNode) {
	if node == nil {
		return
	}

	bst.printTree(node.left)
	log.Println(node.key, node.value)
	bst.printTree(node.right)
}

func (bst *binarySearchTreeSt) Size() int {
	return bst.size
}

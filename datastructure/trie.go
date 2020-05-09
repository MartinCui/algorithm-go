package datastructure

type Trie struct {
	root             *trieNode
	keyCharRangeSize int
}

func NewTrie(keyCharRangeSize int) *Trie {
	return &Trie{root: &trieNode{subNodes: make([]*trieNode, keyCharRangeSize)}, keyCharRangeSize: keyCharRangeSize}
}

type trieNode struct {
	value    interface{}
	subNodes []*trieNode
}

func (t *Trie) Get(key []int) interface{} {
	if len(key) == 0 {
		return nil
	}

	node := t.root
	for i := 0; i < len(key); i++ {
		if node == nil {
			break
		}

		node = node.subNodes[key[i]]
	}

	if node == nil {
		return nil
	} else {
		return node.value
	}
}

func (t *Trie) Set(key []int, value interface{}) {
	parentNode := t.root
	for i := 0; i < len(key); i++ {
		k := key[i]
		currentNode := parentNode.subNodes[k]
		if currentNode == nil {
			currentNode = &trieNode{
				value:    nil,
				subNodes: make([]*trieNode, t.keyCharRangeSize),
			}
			parentNode.subNodes[k] = currentNode
		}

		parentNode = currentNode
	}

	parentNode.value = value
}

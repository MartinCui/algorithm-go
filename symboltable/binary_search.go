package symboltable

type binarySearchSt struct {
	keys   []Comparable
	values []interface{}
}

func NewBinarySearchSt() SymbolTable {
	return &binarySearchSt{
		keys:   make([]Comparable, 0),
		values: make([]interface{}, 0),
	}
}

func (bs *binarySearchSt) Get(key Comparable) interface{} {
	rk, found := bs.rank(key)
	if found {
		return bs.values[rk]
	}

	return nil
}

func (bs *binarySearchSt) rank(key Comparable) (position int, found bool) {
	if len(bs.keys) == 0 {
		return 0, false
	}

	low := 0
	high := len(bs.keys) - 1
	for low <= high && bs.keys[low].CompareTo(key) <= 0 && bs.keys[high].CompareTo(key) >= 0 {
		mid := low + (high-low)/2
		compareResult := bs.keys[mid].CompareTo(key)
		if compareResult > 0 {
			high = mid - 1
		} else if compareResult < 0 {
			low = mid + 1
		} else {
			return mid, true
		}
	}

	if bs.keys[low].CompareTo(key) > 0 {
		return low, false
	} else {
		return high + 1, false
	}
}

func (bs *binarySearchSt) Set(key Comparable, value interface{}) {
	rk, found := bs.rank(key)
	if found {
		bs.values[rk] = value
		return
	}

	bs.keys = append(bs.keys, key)
	bs.values = append(bs.values, value)
	for i := len(bs.keys) - 1; i > rk && i > 0; i-- {
		bs.keys[i] = bs.keys[i-1]
		bs.values[i] = bs.values[i-1]
	}

	bs.keys[rk] = key
	bs.values[rk] = value
}

func (bs *binarySearchSt) Size() int {
	return len(bs.keys)
}

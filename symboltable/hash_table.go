package symboltable

type hashTable struct {
	size       int
	bucketSize int
	keys       []Comparable
	values     []interface{}
}

func NewHashTable() SymbolTable {
	return _hashTableNew(97)
}

func _hashTableNew(bucketSize int) *hashTable {
	return &hashTable{
		size:       0,
		bucketSize: bucketSize,
		keys:       make([]Comparable, bucketSize),
		values:     make([]interface{}, bucketSize),
	}
}

func (ht *hashTable) increaseBucketSize(newBucketSize int) {
	newTable := _hashTableNew(newBucketSize)
	for i, key := range ht.keys {
		if key != nil {
			newTable.Set(key, ht.values[i])
		}
	}

	ht.bucketSize = newTable.bucketSize
	ht.values = newTable.values
	ht.size = newTable.size
	ht.keys = newTable.keys
}

func (ht *hashTable) hashPositionInBucket(key Comparable) int {
	return int(key.HashCode() & 0x7fffffff) % ht.bucketSize
}

func (ht *hashTable) Get(key Comparable) interface{} {
	for position := ht.hashPositionInBucket(key); ht.keys[position] != nil; position = (position + 1) % ht.bucketSize {
		if ht.keys[position].CompareTo(key) == 0{
			return ht.values[position]
		}
	}

	return nil
}

func (ht *hashTable) Set(key Comparable, value interface{}) {
	if ht.size >= (ht.bucketSize/2){
		ht.increaseBucketSize(ht.bucketSize * 2)
	}

	position := ht.hashPositionInBucket(key)
	for ; ht.keys[position] != nil; position = (position + 1) % ht.bucketSize {
		if ht.keys[position].CompareTo(key) == 0{
			ht.values[position] = value
			return
		}
	}

	ht.keys[position] = key
	ht.values[position] = value
	ht.size++
}

func (ht *hashTable) Size() int {
	return ht.size
}

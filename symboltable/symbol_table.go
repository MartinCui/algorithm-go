package symboltable

type Comparable interface {
	CompareTo(other interface{}) int
	HashCode() int32
}

type SymbolTable interface {
	Get(key Comparable) interface{}
	Set(key Comparable, value interface{})
	Size() int
}

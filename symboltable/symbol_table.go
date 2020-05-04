package symboltable

type Comparable interface {
	CompareTo(other interface{}) int
}

type SymbolTable interface {
	Get(key Comparable) interface{}
	Set(key Comparable, value interface{})
	Size() int
}

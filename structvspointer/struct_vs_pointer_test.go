package structvspointer

import (
	"testing"
	"time"
)

const runTimes = 100

type smallStruct struct {
	id            int
	name          string
	address       string
	id0, id1, id2 int
}

type bigStruct struct {
	id0, id1, id2, id3, id4, id5, id6, id7, id8, id9 int
	//name0, name1, name2, name3, name4, name5, name6, name7, name8, name9                               string
	//address0, address1, address2, address3, address4, address5, address6, address7, address8, address9 string
	t0, t1, t2, t3, t4, t5, t6, t7, t8, t9 time.Time
	// d0, d1, d2, d3, d4, d5, d6, d7, d8, d9                                                             time.Duration
}

func smallStructAdd(a, b smallStruct, times int) int {
	if times <= 0 {
		return a.id + b.id
	} else {
		return smallStructAdd(a, b, times-1)
	}
}

func smallPointerAdd(a, b *smallStruct, times int) int {
	if times <= 0 {
		return a.id + b.id
	} else {
		return smallPointerAdd(a, b, times-1)
	}
}

func bigStructAdd(a, b bigStruct, times int) int {
	if times <= 0 {
		return a.id0 + b.id1
	} else {
		return bigStructAdd(a, b, times-1)
	}
}

func bigPointerAdd(a, b *bigStruct, times int) int {
	if times <= 0 {
		return a.id0 + b.id1
	} else {
		return bigPointerAdd(a, b, times-1)
	}
}

func BenchmarkSmallStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = smallStructAdd(smallStruct{}, smallStruct{}, runTimes)
	}
}

func BenchmarkSmallPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = smallPointerAdd(&smallStruct{}, &smallStruct{}, runTimes)
	}
}

func BenchmarkBigStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = bigStructAdd(bigStruct{}, bigStruct{}, runTimes)
	}
}

func BenchmarkBigPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = bigPointerAdd(&bigStruct{}, &bigStruct{}, runTimes)
	}
}

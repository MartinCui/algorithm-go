package algorithm_go

import (
	"math/rand"
	"time"
)

func Exchange(arr []int, i, j int){
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func MakeRandomSliceNoLimit(size int) []int {
	return MakeRandomSlice(size, false)
}

func MakeRandomSlice(size int, limitMax bool) []int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, size)
	for i := 0; i < len(arr); i++ {
		if limitMax {
			arr[i] = rnd.Intn(100)
		} else {
			arr[i] = int(rnd.Int31())
		}
	}
	return arr
}
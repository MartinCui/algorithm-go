package sort

func shellSort(arr []int){
	hopRoot := 3
	var hops []int
	for hop := 1; hop < len(arr); hop = hopRoot*hop + 1 {
		hops = append(hops, hop)
	}

	for hopIndex := len(hops) - 1; hopIndex >= 0; hopIndex-- {
		hopLength := hops[hopIndex]
		for i := hopLength; i < len(arr); i++ {
			for j := i; j >= hopLength && arr[j] < arr[j-hopLength]; j -= hopLength {
				tmp := arr[j-hopLength]
				arr[j-hopLength] = arr[j]
				arr[j] = tmp
			}
		}
	}
}

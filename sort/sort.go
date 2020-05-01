package sort

func exchange(arr []int, i, j int){
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

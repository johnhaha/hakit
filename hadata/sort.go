package hadata

type Sort interface {
	GetValue() int
}

func partition(arr []Sort, low, high int) ([]Sort, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j].GetValue() < pivot.GetValue() {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func QuickSort(arr []Sort, low, high int) []Sort {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = QuickSort(arr, low, p-1)
		arr = QuickSort(arr, p+1, high)
	}
	return arr
}

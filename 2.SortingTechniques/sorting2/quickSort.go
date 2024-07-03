package sorting2

func QuickSort(arr []int, low int, high int) []int {

	if low < high {
		pI := Partition(arr, low, high)
		QuickSort(arr, low, pI-1)
		QuickSort(arr, pI+1, high)
	}
	return arr
}
func Partition(arr []int, low int, high int) int {
	pivot := arr[low]
	i := low
	j := high

	for i < j {
		for arr[i] <= pivot && i <= high-1 {
			i++
		}
		for arr[j] > pivot && j >= low+1 {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[low], arr[j] = arr[j], arr[low]
	return j
}

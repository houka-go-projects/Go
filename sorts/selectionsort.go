//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

func SelectionSort(arr []int) []int {

	length := len(arr)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		tmp := arr[i]
		arr[i] = arr[min]
		arr[min] = tmp
	}
	return arr
}

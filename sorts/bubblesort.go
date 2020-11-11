// Implementation of basic bubble sort algorithm
// Reference: https://en.wikipedia.org/wiki/Bubble_sort

package sorts

func bubbleSort(arr []int) []int {
	swapped := true
	length := len(arr)-1
	for swapped {
		swapped = false
		i := 0
		for i; i < length; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}
	return arr
}

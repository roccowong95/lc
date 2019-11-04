package heap

import "sort"

// Heapify makes a sort.Interface heapified.
func Heapify(arr sort.Interface) {
	l := arr.Len()
	for i := l/2 - 1; i >= 0; i-- {
		adjust(arr, i, l)
	}
}

// Sort sorts a sort.Interface using heap.
func Sort(arr sort.Interface) {
	Heapify(arr)
	for j := arr.Len() - 1; j >= 0; j-- {
		arr.Swap(0, j)
		adjust(arr, 0, j)
	}
}

// only need max len
func adjust(arr sort.Interface, idx, l int) {
	for k := left(idx); k < l; k = left(k) {
		if k+1 < l && arr.Less(k, k+1) {
			k++
		}
		if arr.Less(idx, k) {
			arr.Swap(idx, k)
			idx = k
		}
	}
}

func left(idx int) int {
	return idx*2 + 1
}

func right(idx int) int {
	return idx*2 + 2
}

package heap_test

import (
	"sort"
	"testing"

	"github.com/roccowong95/lc/containers/heap"
)

type ints []int

func (arr ints) Len() int {
	return len(arr)
}

func (arr ints) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr ints) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func Test_heapify(t *testing.T) {
	i := ints([]int{3, 1, 2, 4, 65, 9, 33})
	heap.Heapify(i)
	t.Log(i)
}

func Test_sort(t *testing.T) {
	i := ints([]int{3, 1, 2, 4, 65, 9, 33})
	result := make([]int, len(i))

	copy(result, i)
	sort.Ints(result)
	heap.Sort(i)

	for idx := range result {
		if result[idx] != i[idx] {
			t.Fatal("wrong")
		}
	}
}

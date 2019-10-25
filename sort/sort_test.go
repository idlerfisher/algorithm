package sort_test

import (
	"github.com/idlerfisher/algorithm/sort"
	"testing"
)

func TestInsertSort(t *testing.T) {
	data := []int{5, 2, 1, 3, 4, 8, 7, 0}
	t.Log(data)
	sort.InsertSort(data)
	t.Log(data)
}

func TestMergeSort(t *testing.T) {
	data := []int{5, 2, 1, 3, 4, 8, 7, 0, 2, 11, 50, 20, 21, 13, 100}
	t.Log(data)
	ret := sort.MergeSort(data)
	t.Log(ret)
}

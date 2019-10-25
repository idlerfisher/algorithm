/**
 * Created by Fisher on 2019/10/24.
 * 归并排序
 * 平均复杂度O(n log n) 最好情况O(n log n) 最坏情况O(n log n)  空间复杂度O(n) 稳定
 */

package sort

func MergeSort(data []int) []int {
	dataLen := len(data)
	if dataLen < 2 {
		return data
	}
	return merge(MergeSort(data[:dataLen/2]), MergeSort(data[dataLen/2:]))
}

func merge(left []int, right []int) []int {
	leftLen := len(left)
	rightLen := len(right)
	totalLen := leftLen + rightLen
	result := make([]int, totalLen)
	for n, i, j := 0, 0, 0; n < leftLen+rightLen; n++ {
		if i >= leftLen {
			result[n] = right[j]
			j++
		} else if j >= rightLen {
			result[n] = left[i]
			i++
		} else {
			if left[i] < right[j] {
				result[n] = left[i]
				i++
			} else {
				result[n] = right[j]
				j++
			}
		}
	}
	return result
}

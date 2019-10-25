/**
 * Created by Fisher on 2019/10/20.
 * 插入排序
 * 平均复杂度O(n^2) 最好情况O(n) 最坏情况O(n^2)  空间复杂度O(1) 稳定
 */

package sort

func InsertSort(data []int) {
	n := len(data)
	if n <= 1 {
		return
	}

	j, tmp := 0, 0
	for i := 1; i < n; i++ {
		if data[i] < data[i-1] {
			tmp = data[i]
			data[i] = data[i-1]
			for j = i - 2; j >= 0 && data[j] > tmp; j-- {
				data[j+1] = data[j]
			}
			data[j+1] = tmp
		}
	}
}

package sort

func count(arr []int) []int {
	n := len(arr)
	//均值变化不大的，O(n+m)
	min, max := 0, 0
	for i := 0; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	d := max - min + 1
	cnt := make([]int, d)
	//统计
	for i := 0; i < n; i++ {
		idx := arr[i] - min
		cnt[idx]++
	}
	//还原
	k := 0
	for i := 0; i < d; i++ {
		for j := cnt[i]; j > 0; j-- {
			arr[k] = i + min
			k++
		}
	}
	return arr
}

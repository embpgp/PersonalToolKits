package sort

import "fmt"

// 假设里面的值最大为10**6
// time:O(n),space O(1)
func PosiNegaSwap(arr []int) {

	// 1 遍历，所有正数乘以最大值，并加上相对位置
	pos := 0
	maxValue := 100000
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] > 0 {
			arr[i] = arr[i]*maxValue + pos
			pos++
		}

	}
	//最终pos表示正数的个数
	//负数排序完成后，正数的起始位置应该是n-pos

	//2 swap正负数
	slow, fast := 0, 0
	for ; fast < n; fast++ {
		if arr[fast] < 0 {
			arr[fast], arr[slow] = arr[slow], arr[fast]
			slow++
		}

	}
	fmt.Printf("%+v\n", arr)
	//3 还原正数的位置，抽屉原理，O(n)
	start := n - pos
	for i := start; i < n; i++ {
		for {
			idx := arr[i] % maxValue
			if idx+start == i {
				break
			}
			arr[start+idx], arr[i] = arr[i], arr[start+idx]
		}

	}
	// 4 还原正数
	for i := start; i < n; i++ {
		//fmt.Printf("i:%d, arr[i]:%d\n", i, arr[i])
		arr[i] = arr[i] / maxValue
	}
	fmt.Printf("%+v\n", arr)

}

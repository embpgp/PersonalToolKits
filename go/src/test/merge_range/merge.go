package merge

import "fmt"

func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}

func mergeRange(arr [][]int) [][]int {
	n := len(arr)
	//sort
	for i := 0; i < n; i++ {
		isSwaped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j][0] > arr[j+1][0] {
				arr[j], arr[j+1] = arr[j+1], arr[j] //swap
				isSwaped = true
			}

		}
		if !isSwaped {
			break
		}
	}
	fmt.Printf("sort:%+v\n", arr)
	var merged [][]int
	merged = append(merged, arr[0])
	for i := 1; i < n; i++ {
		l, r := arr[i][0], arr[i][1]
		//fmt.Printf("L:%d,R:%d\n", l, r)
		last := len(merged) - 1
		if merged[last][1] < l {
			merged = append(merged, arr[i])
		} else {
			merged[last][1] = max(merged[last][1], r)
		}

	}
	return merged

}

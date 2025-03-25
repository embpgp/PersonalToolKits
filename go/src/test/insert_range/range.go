package insertrange

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func InsertRange(arr [][]int, insert []int) [][]int {

	res := make([][]int, 0)
	i, l := 0, len(arr)
	//左边没交集
	for i < l && arr[i][1] < insert[0] {
		res = append(res, arr[i])
		i++
	}
	//有交集
	for i < l && arr[i][0] <= insert[1] {
		insert[0] = min(arr[i][0], insert[0])
		insert[1] = max(arr[i][1], insert[1])
		i++
	}
	res = append(res, insert)
	for i < l {
		res = append(res, arr[i])
		i++
	}
	return res

}

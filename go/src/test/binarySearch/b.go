package binarysearch

// 若有重复，返回最左边的值
func BinarySearchInsertion(nums []int, target int) int {
	n := len(nums)
	i, j := 0, n-1
	for i <= j {
		mid := (j-i)/2 + i
		if nums[mid] < target {
			i = mid + 1
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			// return mid 若相等
			j = mid - 1
			//若找到最右边的呢
			//i = mid + 1
		}
	}
	return i
}

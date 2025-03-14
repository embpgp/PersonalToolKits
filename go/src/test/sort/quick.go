package sort

import "fmt"

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := partition(nums, left, right)
	quickSort(nums, left, mid-1)
	quickSort(nums, mid+1, right)

}

// 返回调整完毕之后的pivot的位置，便于之后对两边的数组进行快排
func partition(nums []int, left, right int) int {
	// choose a base num
	pivot := nums[right]
	for left < right {

		for left < right && nums[left] <= pivot {
			left++
		}
		fmt.Printf("a[%d]=%d,a[%d]=%d\n", left, nums[left], right, nums[right])
		nums[right] = nums[left]
		for right > left && nums[right] >= pivot {
			right--
		}
		fmt.Printf("a[%d]=%d,a[%d]=%d\n", left, nums[left], right, nums[right])
		nums[left] = nums[right]
	}
	nums[left] = pivot
	return left
}

func quickSortV1(a []int) {
	partitionV1(a, 0, len(a)-1)
}

func partitionV1(a []int, left, right int) {
	if left >= right {
		return

	}
	pivot := a[left]
	i, j := left, right
	for i < j {
		for ; i < j && a[j] >= pivot; j-- {
		}
		a[i] = a[j]
		for ; i < j && a[i] <= pivot; i++ {
		}
		a[j] = a[i]
	}
	a[j] = pivot
	partitionV1(a, left, j-1)
	partitionV1(a, j+1, right)

}

func findTopKByQuick(a []int, k int) int {
	if k >= len(a) {
		k = len(a) - 1
	}
	return partitionTopK(a, 0, len(a)-1, k)
}

func partitionTopK(a []int, left, right, k int) int {
	if left >= right {
		return a[left]

	}
	pivot := a[left]
	i, j := left, right
	for i < j {
		for ; i < j && a[j] >= pivot; j-- {
		}
		a[i] = a[j]
		for ; i < j && a[i] <= pivot; i++ {
		}
		a[j] = a[i]
	}
	a[j] = pivot
	if j == k {
		return a[j]
	}
	if j < k {
		return partitionTopK(a, j+1, right, k)
	} else {
		return partitionTopK(a, left, j-1, k)
	}

}

package sort

import "fmt"

func heapSort(a []int) {
	n := len(a)
	//先构建一个最大堆，从最后一个元素开始
	for i := n/2 - 1; i >= 0; i-- {
		heapify(a, n, i)
	}
	for i := n - 1; i > 0; i-- {
		a[i], a[0] = a[0], a[i]
		heapify(a, i, 0) //从0开始调整
	}
}

// heapify 调整以i为根的子树为最大堆
func heapify(a []int, n, i int) {
	largest := i // 初始化当前节点为最大值
	left := 2*i + 1
	right := 2*i + 2
	if left < n && a[left] > a[largest] {
		largest = left
	}
	if right < n && a[right] > a[largest] {
		largest = right
	}
	if i != largest {
		a[i], a[largest] = a[largest], a[i]
		heapify(a, n, largest)
	}

}

func findTopKbyHeap(a []int, k int) int {
	n := len(a)
	//先构建一个最大堆，从最后一个元素开始
	for i := n/2 - 1; i >= 0; i-- {
		heapify(a, n, i)
	}
	for i := n - 1; i >= k; i-- {
		a[i], a[0] = a[0], a[i]
		heapify(a, i, 0) //从0开始调整
	}
	return a[k]
}

type streamMinHeap struct {
	k    int
	data []int
}

// 从i节点下沉
func (h *streamMinHeap) Down(i int) {
	small := i
	left := 2*i + 1
	right := 2*i + 2
	if left < len(h.data) && h.data[left] < h.data[i] {
		small = left
	}
	if right < len(h.data) && h.data[right] < h.data[i] {
		small = right
	}
	if small != i {
		h.data[small], h.data[i] = h.data[i], h.data[small]
		h.Down(small)
	}
}

// 从i节点上浮
func (h *streamMinHeap) Up(i int) {
	for {
		p := (i - 1) / 2
		if p >= 0 && h.data[i] < h.data[p] {
			h.data[i], h.data[p] = h.data[p], h.data[i]
			i = p
		} else {
			break
		}

	}

}

func (h *streamMinHeap) Add(num int) {
	if len(h.data) < h.k {
		h.data = append(h.data, num)
		h.Up(len(h.data) - 1)
	} else {
		if h.data[0] < num {
			h.data[0] = num
			h.Down(0)
		}
	}
}

func (h *streamMinHeap) Print() {
	fmt.Printf("%v\n", h.data)
}

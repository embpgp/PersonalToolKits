package slice

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_f(t *testing.T) {
	p := make([]int, 1023)
	p = append(p, 1)
	f1(p)
	//(*reflect.SliceHeader)(unsafe.Pointer(&p)).Len = 5 //强制修改slice长
	//(*reflect.SliceHeader)(unsafe.Pointer(&p)).Cap = 3 //强制修改slice长
	fmt.Printf("%v,%+v,%v,%v\n", unsafe.Pointer(&p), p, cap(p), len(p))
	f(&p)
	fmt.Printf("%+v\n", p)

	/*var i1 int = 1
	var i2 int8 = 2
	var i3 int16 = 3
	var i4 int32 = 4
	var i5 int64 = 5
	fmt.Println(unsafe.Sizeof(i1))
	fmt.Println(unsafe.Sizeof(i2))
	fmt.Println(unsafe.Sizeof(i3))
	fmt.Println(unsafe.Sizeof(i4))
	fmt.Println(unsafe.Sizeof(i5))
	*/
}

func Test_grow(t *testing.T) {
	arr := make([]int, 0)
	lastCap := cap(arr)
	for i := 0; i < 2000; i++ {
		if cap(arr) != lastCap {
			fmt.Printf("len:%d,cap:%d\n", len(arr), cap(arr))
			lastCap = cap(arr)
		}

		arr = append(arr, i)
	}
}

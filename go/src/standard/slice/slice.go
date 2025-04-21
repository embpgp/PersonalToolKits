package slice

import (
	"fmt"
	"unsafe"
)

func f(p *[]int) {
	*p = append(*p, 6)
}

func f1(p []int) {
	fmt.Printf("%v,%+v,%v,%v\n", unsafe.Pointer(&p), p, cap(p), len(p))
	p = append(p, 4)
	fmt.Printf("%v,%+v,%v,%v\n", unsafe.Pointer(&p), p, cap(p), len(p))
}

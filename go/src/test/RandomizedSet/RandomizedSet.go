package randomizedset

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	arr        []int
	length     int
	val2IdxMap map[int]values
}

type values struct {
	val int
	idx int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		arr:        make([]int, 0),
		length:     0,
		val2IdxMap: make(map[int]values),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.val2IdxMap[val]; ok {
		return false
	}
	this.arr = append(this.arr, val)
	this.val2IdxMap[val] = values{val: val, idx: this.length}
	this.length++
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	vNode, ok := this.val2IdxMap[val]
	if !ok {
		return false
	}
	delIdx := vNode.idx
	last := this.arr[this.length-1]
	//fill
	this.arr[delIdx] = last

	this.val2IdxMap[last] = values{val: last, idx: delIdx}
	this.arr = this.arr[:this.length-1]
	this.length--
	delete(this.val2IdxMap, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	v := this.arr[rand.Intn(this.length)]
	return v
}

func (this *RandomizedSet) Printf() {
	fmt.Printf("arr:%+v, len:%d, map:%+v\n", this.arr, this.length, this.val2IdxMap)
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

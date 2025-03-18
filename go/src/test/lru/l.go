package lru

import (
	"fmt"
	"sync"
)

type node struct {
	Key   string
	Value int
	Prev  *node
	Next  *node
}

type cacheLru struct {
	lock     sync.RWMutex
	MaxSize  int
	DataMap  map[string]*node
	DataHead node
	DataTail node
}

func NewLRU(len int) *cacheLru {
	m := &cacheLru{}
	m.Init(len)
	return m
}

func (m *cacheLru) Init(len int) {
	m.MaxSize = len
	m.DataMap = make(map[string]*node, len)
	m.DataHead.Prev = &m.DataTail
	m.DataHead.Next = &m.DataTail
	m.DataTail.Prev = &m.DataHead
	m.DataTail.Next = &m.DataHead

}

func (m *cacheLru) Put(key string, value int) {
	m.lock.Lock()
	defer m.lock.Unlock()
	v := m.getNode(key)
	if v != nil { //重复key,更新value
		v.Value = value
		return
	}
	var tmp node
	tmp.Value = value
	tmp.Key = key
	if len(m.DataMap) >= m.MaxSize {
		clean := m.DataTail.Prev
		clean.Prev.Next = &m.DataTail
		m.DataTail.Prev = clean.Prev
		fmt.Printf("clean: key:%s, value: %d\n", clean.Key, clean.Value)
		delete(m.DataMap, clean.Key)
	}
	curr := m.DataHead.Next
	tmp.Next = curr
	curr.Prev = &tmp
	tmp.Prev = &m.DataHead
	m.DataHead.Next = &tmp
	if len(m.DataMap) == 0 {
		m.DataTail.Prev = &tmp
	}
	m.DataMap[key] = &tmp

}

func (m *cacheLru) getNode(key string) *node {
	v, ok := m.DataMap[key]
	if !ok {
		return nil
	}
	//exist ,swap 先脱离
	v.Prev.Next = v.Next
	v.Next.Prev = v.Prev
	//再插入head后面
	v.Next = m.DataHead.Next
	v.Prev = &m.DataHead
	m.DataHead.Next = v

	return v
}

func (m *cacheLru) Get(key string) int {
	m.lock.Lock()
	defer m.lock.Unlock()
	v := m.getNode(key)
	if v == nil {
		return -1
	}
	return v.Value
}

func (m *cacheLru) Print() {
	m.lock.RLock()
	defer m.lock.RUnlock()
	v := m.DataHead.Next
	for {
		if v == &m.DataTail {
			break
		}
		fmt.Printf("[%s] -> %d\n", v.Key, v.Value)
		v = v.Next
	}
	fmt.Printf("len: %d\n", len(m.DataMap))

}

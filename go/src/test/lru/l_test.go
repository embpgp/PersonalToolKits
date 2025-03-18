package lru

import (
	"fmt"
	"testing"
)

func TestNewLRU(t *testing.T) {
	m := NewLRU(5)
	m.Put("one", 1)
	m.Put("two", 2)
	//m.Print()
	m.Put("three", 3)
	m.Put("four", 4)

	m.Put("five", 5)
	m.Put("fix", 6)
	m.Print()
	fmt.Printf("get:%d\n", m.Get("three"))
	m.Print()

}

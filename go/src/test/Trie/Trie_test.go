package trie

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {

	root := Constructor()
	root.Insert("hello")
	root.Insert("hdd")
	fmt.Printf("%+v\n", root.Search("hellofd"))
}

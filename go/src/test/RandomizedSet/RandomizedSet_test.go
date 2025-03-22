package randomizedset

import (
	"fmt"
	"testing"
)

func TestRandomizedSet_GetRandom(t *testing.T) {

	sets := Constructor()
	sets.Insert(0)
	sets.Insert(1)
	sets.Printf()
	sets.Remove(0)
	sets.Printf()
	sets.Insert(2)
	sets.Printf()
	sets.Remove(1)
	sets.Printf()
	fmt.Printf("%d\n", sets.GetRandom())

}

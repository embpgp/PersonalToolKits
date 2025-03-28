package trie

import "strings"

type Trie struct {
	ch    [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	w := strings.ToLower(word)
	node := this
	for i := 0; i < len(w); i++ {
		c := w[i] - 'a'
		if node.ch[c] == nil {
			node.ch[c] = &Trie{}
		}
		node = node.ch[c]
	}
	node.isEnd = true

}

func (this *Trie) Search(word string) bool {

	trie := this.searchNode(word)
	return trie != nil && trie.isEnd == true

}

func (this *Trie) searchNode(prefix string) *Trie {
	w := strings.ToLower(prefix)
	var i int
	var c byte
	node := this
	for i = 0; i < len(w); i++ {
		c = w[i] - 'a'
		if node.ch[c] == nil {
			return nil
		}
		node = node.ch[c]
	}
	return node

}

func (this *Trie) StartsWith(prefix string) bool {
	return this.searchNode(prefix) != nil

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

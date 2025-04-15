package main

type Node struct {
	children [26]*Node
	endHere  bool
}

type WordDictionary struct {
	root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{root: &Node{}}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.root
	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &Node{}
		}
		curr = curr.children[idx]
	}
	curr.endHere = true
}

func (this *WordDictionary) Search(word string) bool {
	return searchHelper(word, 0, this.root)
}

func searchHelper(word string, idx int, curr *Node) bool {
	if curr == nil {
		return false
	}
	if idx == len(word) {
		return curr.endHere
	}
	ch := word[idx]
	if ch == '.' {
		for i := 0; i < 26; i++ {
			if searchHelper(word, idx+1, curr.children[i]) {
				return true
			}
		}
		return false
	} else {
		return searchHelper(word, idx+1, curr.children[ch-'a'])
	}
}

for i >= j {
	3
}




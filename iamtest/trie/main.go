package main

import (
	"fmt"
	"sort"
	"strings"
)

func (this *trie) Print() {
	printNode(this.root, 0)
}

func printNode(n *node, level int) {
	// Get all characters at this level and sort them
	keys := make([]rune, 0, len(n.children))
	for k := range n.children {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for _, char := range keys {
		child := n.children[char]

		// Create the indentation
		indent := strings.Repeat("  ", level)

		// Check if this character completes a word
		marker := ""
		if child.isEnd {
			marker = " (✓)"
		}

		fmt.Printf("%s└─ %c%s\n", indent, char, marker)

		// Recurse to the next level
		printNode(child, level+1)
	}
}

func main() {
	t := getTrie()

	t.Insert("ashish")
	t.Insert("anjna")
	t.Insert("anjali")

	t.Insert("apple")
	t.Insert("app")
	t.Insert("apricot")
	t.Insert("banana")
	fmt.Println("a", t.StartWith("ashu"))
	fmt.Println("b", t.Search("ashi"))
	fmt.Println("C", t.Search("anjna"))
	fmt.Println("d", t.GetSuggestions("a"))
}

type node struct {
	children map[rune]*node
	isEnd    bool
}

type trie struct {
	root *node
}

func getTrie() trie {
	return trie{
		root: &node{
			children: make(map[rune]*node),
		},
	}
}

func (t *trie) Insert(word string) {
	nd := t.root
	for _, c := range word {
		if _, ok := nd.children[c]; !ok {
			nd.children[c] = &node{children: make(map[rune]*node)}
		}
		nd = nd.children[c]

	}
	nd.isEnd = true
}
func (t *trie) Search(word string) bool {
	root := t.root
	for _, c := range word {
		if _, ok := root.children[c]; !ok {
			return false
		}
		root = root.children[c]
	}
	return root.isEnd
}
func (t *trie) StartWith(word string) bool {
	root := t.root
	for _, c := range word {
		if _, ok := root.children[c]; !ok {
			return false
		}
		root = root.children[c]
	}
	return true
}

func (t *trie) GetSuggestions(word string) []string {
	res := make([]string, 0)
	root := t.root
	for _, c := range word {
		if next, ok := root.children[c]; !ok {
			return res
		} else {
			root = next
		}
	}
	t.dfs(root, word, &res)
	return res
}

func (t *trie) dfs(root *node, word string, res *[]string) {
	if len(*res) == 3 {
		return
	}
	if root.isEnd {
		*res = append(*res, word)
	}

	for char, child := range root.children {
		t.dfs(child, word+string(char), res)
	}
}

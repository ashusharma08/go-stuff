package main

import (
	"fmt"
)

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {

	wordSet := make(map[string]bool, 0)
	for _, item := range wordList {
		wordSet[item] = true
	}
	if _, ok := wordSet[endWord]; !ok {
		return 0
	}

	queue := []string{beginWord}

	dis := 1

	for len(queue) > 0 {
		level := len(queue)

		for i := 0; i < level; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr == endWord {
				return dis
			}

			for j := 0; j < len(curr); j++ {
				for c := 'a'; c <= 'z'; c++ {
					if byte(c) == curr[j] {
						continue
					}
					newW := curr[:j] + string(c) + curr[j+1:]
					if _, ok := wordSet[newW]; ok {
						queue = append(queue, newW)
						delete(wordSet, newW)
					}
				}
			}
		}
		dis++
	}

	return 0
}

func ladderLength_1(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 0
	}
	queue := make([]struct {
		Level int
		Val   string
	}, 0, len(wordList))
	wordList = append(wordList, beginWord)
	graph := giveMeGraph(wordList)
	queue = append(queue, struct {
		Level int
		Val   string
	}{Level: 1, Val: beginWord})

	visited := map[string]bool{beginWord: true}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.Val == endWord {
			return curr.Level
		}

		for i := 0; i < len(curr.Val); i++ {
			pattern := curr.Val[:i] + "*" + curr.Val[i+1:]
			for _, item := range graph[pattern] {

				if !visited[item] {
					visited[item] = true
					queue = append(queue, struct {
						Level int
						Val   string
					}{
						Val: item, Level: curr.Level + 1,
					})
				}
			}
		}
		// for _, neigh := range graph[curr.Val] {
		// 	if !visited[neigh] {
		// 		visited[neigh] = true
		// 		queue = append(queue, struct {
		// 			Level int
		// 			Val   string
		// 		}{Val: neigh, Level: curr.Level + 1})
		// 	}
		// }
	}
	return 0
}

func giveMeGraph(wordList []string) map[string][]string {
	graph := make(map[string][]string, len(wordList))

	for _, w := range wordList {
		for i := 0; i < len(w); i++ {
			pattern := w[:i] + "*" + w[i+1:]
			graph[pattern] = append(graph[pattern], w)
		}
	}
	return graph
}

func shouldAppend(val1, val2 string) bool {
	if len(val1) != len(val2) {
		return false
	}
	sl1 := []rune(val1)
	diff := 0
	for i, item := range val2 {
		if sl1[i] != item {
			diff++
		}
		if diff > 1 {
			break
		}
	}
	return diff == 1
}

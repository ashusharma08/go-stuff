package main

import "fmt"

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	dictionary := make(map[string]bool, len(wordList))
	for _, item := range wordList {
		dictionary[item] = true
	}

	if _, ok := dictionary[endWord]; !ok {
		return -1
	}

	queue := []string{beginWord}
	dis := 1
	for len(queue) > 0 {
		level := len(queue)
		for range level {
			curr := queue[0]
			queue = queue[1:]

			if curr == endWord {
				return dis
			}

			for i := 0; i < len(curr); i++ {
				for j := 'a'; j <= 'z'; j++ {
					if curr[i] == byte(j) {
						continue
					}
					char := curr[:i] + string(j) + curr[i+1:]
					if _, ok := dictionary[char]; ok {
						queue = append(queue, char)
						delete(dictionary, char)
					}
				}
			}
		}
		dis++
	}

	return dis
}

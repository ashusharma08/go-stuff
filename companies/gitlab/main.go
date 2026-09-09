package main

import (
	"fmt"
	"sync"
)

func main() {
	// wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	// begin, end := "hit", "cog"

	// fmt.Println(twoWayWordLadder(wordList, begin, end))
	treeLCA()
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeLCA() {
	// Build tree:
	//         3
	//        / \
	//       5   1
	//      / \ / \
	//     6  2 0  8
	//       / \
	//      7   4
	root := &Node{Val: 3}
	root.Left = &Node{Val: 5}
	root.Right = &Node{Val: 1}

	root.Left.Left = &Node{Val: 6}
	root.Left.Right = &Node{Val: 2}
	root.Right.Left = &Node{Val: 0}
	root.Right.Right = &Node{Val: 8}

	root.Left.Right.Left = &Node{Val: 7}
	root.Left.Right.Right = &Node{Val: 4}

	p := root.Left             // 5
	q := root.Left.Right.Right // 4

	fmt.Println(lowestCommonAncestor(root, p, q))
}

func lowestCommonAncestor(node, p, q *Node) *Node {
	//traverse left
	//traverse right..
	// find the common root.. hoW? both branches return whether left and right.. so we are norrowing the bra ch

	if node == nil || node == p || node == q {
		fmt.Println("returning n ", node)
		return node
	}

	left := lowestCommonAncestor(node.Left, p, q)
	right := lowestCommonAncestor(node.Right, p, q)
	fmt.Println("left right  ", left, right)

	if left != nil && right != nil {
		return node
	}
	if left != nil {
		return left
	}

	return right
}

func twoWayWordLadder(wordList []string, begin, end string) int {
	exists := make(map[string]bool, 0)

	for _, item := range wordList {
		exists[item] = true
	}

	beginSet, endSet := map[string]bool{begin: true}, map[string]bool{end: true}
	currIdx := 1
	for len(beginSet) != 0 && len(endSet) != 0 {
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}

		newSet := make(map[string]bool, 0)
		for curr := range beginSet {
			for j := 0; j < len(curr); j++ {
				for c := 'a'; c <= 'z'; c++ {
					newW := curr[:j] + string(c) + curr[j+1:]
					if endSet[newW] {
						return currIdx + 1
					}
					if exists[newW] {
						newSet[newW] = true
						delete(exists, newW)
					}
				}
			}
		}
		beginSet = newSet
		currIdx++
	}
	return 0
}

func wordLadder(wordList []string, begin, end string) int {

	exists := make(map[string]bool, 0)

	for _, item := range wordList {
		exists[item] = true
	}

	queue := []string{begin}
	dis := 1
	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr == end {
				return dis
			}
			for j := 0; j < len(curr); j++ {
				for x := 'a'; x <= 'z'; x++ {
					if byte(x) == curr[j] {
						continue
					}
					nw := curr[:j] + string(x) + curr[j+1:]
					if exists[nw] {
						queue = append(queue, nw)
						delete(exists, nw)
					}
				}
			}
		}
		dis++
	}
	return 0
}
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		results <- j * 2
	}
}

func matestin() {
	jobs := make(chan int, 5)
	results := make(chan int)
	var wg sync.WaitGroup
	for w := 0; w < 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	fmt.Println("Done sending jobs")
	done := make(chan struct{})
	go func() {
		for r := range results {
			fmt.Println(r)
		}
		done <- struct{}{}
	}()
	wg.Wait()
	close(results)
	<-done
}

package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct{}

// func (s *Solution) isPalindrome(head *ListNode) bool {
// 	// TODO: Write your code here
// 	//find middle
// 	if head == nil || head.Next == nil {
// 		return true
// 	}
// 	fast := head
// 	slow := head
// 	for fast != nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 	}
// 	reversed := s.reverse(slow)
// 	cpr := reversed

// 	for head != nil && reversed != nil {
// 		if head.Val != reversed.Val {
// 			break
// 		}
// 		head = head.Next
// 		reversed = reversed.Next
// 	}
// 	s.reverse(cpr)
// 	if head == nil || reversed == nil { // if both halves match
// 		return true
// 	}

// 	return false
// }

// func (s *Solution) reverse(head *ListNode) *ListNode {
// 	var prev *ListNode
// 	for head != nil {
// 		next := head.Next
// 		head.Next = prev
// 		prev = head
// 		head = next
// 	}
// 	return prev
// }

// 0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1
func (s *Solution) LongestSubarrayReplaceOne(arr []int, k int) int {
	start, max := 0, 0
	occ := make(map[int]int, 0)
	for winEnd, val := range arr {
		occ[val]++
		if occ[0] > k {
			occ[arr[start]]--
			start++
		}
		if max < winEnd-start+1 {
			max = winEnd - start + 1
		}
	}
	return max
}

// "oidbcaf", "abc"
func (s *Solution) findPermutation(str string, pattern string) bool {
	// TODO: Write your code here
	patternMap := make(map[rune]int)
	for _, r := range pattern {
		patternMap[r]++
	}
	matched := 0
	start := 0
	for winEnd, val := range str {
		if _, ok := patternMap[val]; ok {
			patternMap[val]--
			if patternMap[val] == 0 {
				matched++
			}
		}

		if matched == len(patternMap) {
			return true
		}

		if winEnd >= len(pattern)-1 {
			if v, ok := patternMap[rune(str[start])]; ok {
				if v == 0 {
					matched--
				}
				patternMap[rune(str[start])]++
			}
			start++
		}
	}
	return false
}

func (s *Solution) findLongestSubstring(str string, k int) int {
	start, maxLen, maxLetter := 0, 0, 0
	occ := make(map[rune]int, 0)
	for winEnd, val := range str {
		occ[val]++
		if maxLetter < occ[val] {
			maxLetter = occ[val]
		}
		if k < winEnd-start+1-maxLetter {
			occ[rune(str[start])]--
			if occ[rune(str[start])] == 0 {
				delete(occ, rune(str[start]))
			}
			start++
		}
		if maxLen < winEnd-start+1 {
			maxLen = winEnd - start + 1
		}
	}
	return maxLen
}

// ["A", "B", "C", "A", "C"]
func (s Solution) fruitBasket(arr []string) int {
	maxLength := 0
	start := 0
	ooc := make(map[string]int)
	for windowEnd, val := range arr {
		ooc[val]++
		if len(ooc) > 2 {
			ooc[arr[start]]--
			if ooc[arr[start]] == 0 {
				delete(ooc, arr[start])
			}
			start++
		}
		maxLength = int(math.Max(float64(maxLength), float64(windowEnd-start+1)))
	}
	return maxLength
}

// Input: str="araaci", K=2

func (s *Solution) findLength(str string, k int) int {
	start, max := 0, 0
	mp := make(map[rune]int, 0)
	for windowEnd, value := range str {
		mp[value]++
		if len(mp) > k {
			toRemove := rune(str[start])
			mp[toRemove]--
			if mp[toRemove] == 0 {
				delete(mp, toRemove)
			}
			start++
		}
		max = int(math.Max(float64(max), float64(windowEnd-start+1)))
	}

	return max
}

func (s *Solution) findMaxSumSubArray(k int, arr []int) int {
	if len(arr) < k {
		panic("invalidArgument")
	}
	maxSum, newSum := 0, 0
	start, end := 0, 0
	for end < len(arr) {
		newSum = newSum + arr[end]
		if end >= k-1 {
			if newSum > maxSum {
				maxSum = newSum
			}
			newSum -= arr[start]
			start++
		}
		end++
	}
	return maxSum
}

func main() {
	// sol := Solution{}
	// head := &ListNode{Val: 1}
	// head.Next = &ListNode{Val: 2}
	// head.Next.Next = &ListNode{Val: 2}
	// head.Next.Next.Next = &ListNode{Val: 1}
	// fmt.Printf("Is palindrome: %v\n", sol.isPalindrome(head))

	// head.Next.Next.Next.Next = &ListNode{Val: 2}
	// fmt.Printf("Is palindrome: %v\n", sol.isPalindrome(head))
	// s := Solution{}
	// fmt.Println("Maximum sum of a subarray of size K: ",
	// 	s.findMaxSumSubArray(3, []int{2, 1, 5, 1, 3, 2}))
	// fmt.Println("Maximum sum of a subarray of size K: ",
	// 	s.findMaxSumSubArray(2, []int{2, 3, 4, 1, 5}))
	// fmt.Println("Maximum sum of a subarray of size K: ",
	// 	s.findMaxSumSubArray(1, []int{1, 2, 3, 4, 5}))

	// fmt.Println(s.findLength("araaci", 2))
	// fmt.Println(s.findLength("araaci", 1))
	// fmt.Println(s.findLength("cbbebi", 3))

	// fmt.Println(s.fruitBasket([]string{"A", "B", "C", "A", "C"}))
	// fmt.Println(s.fruitBasket([]string{"A", "B", "C", "B", "B", "C"}))

	// fmt.Println(s.findLongestSubstring("aabccbb", 2))
	// fmt.Println(s.findLongestSubstring("aabccabb", 2))
	// fmt.Println(s.findLongestSubstring("abbcb", 1))
	// fmt.Println(s.findLongestSubstring("abccde", 1))

	// fmt.Println(s.LongestSubarrayReplaceOne([]int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}, 2))
	// fmt.Println(s.LongestSubarrayReplaceOne([]int{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1}, 3))
	// fmt.Println(s.LongestSubarrayReplaceOne([]int{1, 0, 0, 1, 1, 0, 1, 1}, 2))

	// fmt.Println("Permutation exist:", s.findPermutation("oidbcaf", "abc"))
	// fmt.Println("Permutation exist:", s.findPermutation("odicf", "dc"))
	// fmt.Println("Permutation exist:", s.findPermutation("bcdxabcdy", "bcdyabcdx"))
	// fmt.Println("Permutation exist:", s.findPermutation("aaacb", "abc"))
	// fmt.Println("Permutation exist:", s.findPermutation("anothertestcase", "acest"))
	// concurrencyH2o()
	// concurrenyFooBar()
	// concurrenyFooBarv2()
	diningPhilosphers()
}

type Fork struct {
	sync.Mutex
}
type Philosopher struct {
	id        int
	leftFork  *Fork
	rightFork *Fork
}

func (p *Philosopher) Eat(wg *sync.WaitGroup) {
	defer wg.Done()
	for range 3 {
		p.leftFork.Lock()
		p.rightFork.Lock()

		fmt.Println(p.id, " is eating")
		time.Sleep(1 * time.Second) //eating for a second

		p.leftFork.Unlock()
		p.rightFork.Unlock()

		fmt.Println(p.id, " is thinking")
		time.Sleep(1 * time.Second) //thinking

	}
}

func diningPhilosphers() {
	var ph [5]Philosopher
	var f [5]Fork
	i := 0
	for range 5 {
		ph[i] = Philosopher{
			id:        i,
			leftFork:  &f[i],
			rightFork: &f[(i+1)%5],
		}
		i++
	}
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go ph[i].Eat(&wg)
	}

	wg.Wait()
}

func (f *FooBar) Printfoo(num int) {
	for range num {
		<-f.foo
		fmt.Println("foo")
		f.bar <- struct{}{}
	}
}
func (f *FooBar) PrintBar(num int) {
	for range num {
		<-f.bar
		fmt.Println("bar")
		f.foo <- struct{}{}
	}
}

func concurrenyFooBarv2() {
	fb := &FooBar{
		foo: make(chan struct{}, 1),
		bar: make(chan struct{}),
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fb.Printfoo(5)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fb.PrintBar(5)
	}()
	fb.foo <- struct{}{}
	wg.Wait()
}

type FooBar struct {
	foo chan struct{}
	bar chan struct{}
	wg  *sync.WaitGroup
}

func (f *FooBar) fooPrinter() {
	f.foo <- struct{}{}
	fmt.Println("FOO")
	f.wg.Done()
}
func (f *FooBar) barPrinter() {
	f.bar <- struct{}{}
	f.wg.Add(1)
	f.wg.Wait()
	fmt.Println("BAR")
	<-f.foo
	<-f.bar
}

func concurrenyFooBar() {
	fb := &FooBar{
		foo: make(chan struct{}, 1),
		bar: make(chan struct{}, 1),
		wg:  &sync.WaitGroup{},
	}
	var wg sync.WaitGroup
	for range 10 {
		wg.Add(2)
		go func() {
			defer wg.Done()
			fb.fooPrinter()
		}()
		go func() {
			defer wg.Done()
			fb.barPrinter()
		}()
	}
	wg.Wait()
}

type Water struct {
	Hydrogen chan struct{}
	Oxygen   chan struct{}
	Wg       *sync.WaitGroup
}

func newWater() *Water {
	return &Water{
		Hydrogen: make(chan struct{}, 2),
		Oxygen:   make(chan struct{}, 1),
		Wg:       &sync.WaitGroup{},
	}
}

func (w *Water) oxygen(printer func()) {
	w.Oxygen <- struct{}{}
	w.Wg.Add(2)
	w.Wg.Wait()
	printer()
	<-w.Hydrogen
	<-w.Hydrogen
	<-w.Oxygen
}
func (w *Water) hydrogen(printer func()) {
	w.Hydrogen <- struct{}{}
	printer()
	w.Wg.Done()

}
func concurrencyH2o() {

	w := newWater()
	oxygenPrinter := func() { fmt.Println("O") }
	hydPrinter := func() { fmt.Println("H") }

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			w.hydrogen(hydPrinter)
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			w.hydrogen(hydPrinter)

		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			w.oxygen(oxygenPrinter)
		}()
	}
	wg.Wait()

}

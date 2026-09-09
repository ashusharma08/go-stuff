package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("title", "paper"))
	fmt.Println(isIsomorphic("ab", "cc`"))
}

func isIsomorphic(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	mp1 := make(map[rune]byte)
	mp2 := make(map[byte]rune)
	for idx, item := range str1 {
		v := str2[idx]
		if x, ok := mp1[item]; !ok {
			mp1[item] = v
		} else if v != x {
			return false
		}

		if x, ok := mp2[v]; !ok {
			mp2[v] = item
		} else if x != item {
			return false
		}
	}
	return true
}

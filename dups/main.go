package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	f, err := os.ReadFile("something.txt")
	if err != nil {
		fmt.Errorf("err %#v", err)
		return
	}
	items := strings.Split(string(f), "\n")
	slices.Sort(items)
	unique := slices.Compact(items)
	alphas := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	getCamelCase := func(str string) string {
		var allAlpha strings.Builder
		for _, x := range str {
			if strings.Contains(alphas, string(x)) || x == ' ' {
				allAlpha.WriteRune(x)
			}
		}
		words := strings.Split(allAlpha.String(), " ")
		key := strings.ToLower(words[0])
		for _, word := range words[1:] {
			key += cases.Title(language.AmericanEnglish).String(word)
		}
		return key
	}

	var bbb strings.Builder
	bbb.WriteString("chrs:= []*store.CharLabel{")
	for _, item := range unique {
		bbb.WriteString(fmt.Sprintf(`{Code: "%s", Description:"%s" },`, getCamelCase(item), item))
	}
	bbb.WriteString("}")
	fmt.Println(bbb.String())
}

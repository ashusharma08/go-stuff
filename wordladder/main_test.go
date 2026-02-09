package main

import "testing"

func Test_shouldAppend(t *testing.T) {
	input := []struct {
		Description string
		Val1        string
		Val2        string
		Expected    bool
	}{
		{
			Description: "Should return true",
			Val1:        "dog",
			Val2:        "cog",
			Expected:    true,
		},
		{
			Description: "Should return false",
			Val1:        "dog",
			Val2:        "cot",
			Expected:    false,
		},
		{
			Description: "Should return false",
			Val1:        "dog",
			Val2:        "dogo",
			Expected:    false,
		},
		{
			Description: "Should return true",
			Val1:        "doog",
			Val2:        "dooo",
			Expected:    true,
		},
		{
			Description: "Should return true",
			Val1:        "ass",
			Val2:        "cas",
			Expected:    true,
		},
	}
	t.Parallel()
	for _, item := range input {
		t.Run(item.Description, func(t *testing.T) {
			if res := shouldAppend(item.Val1, item.Val2); res != item.Expected {
				t.Fatalf("item %#v, expected %t, got %t", item, item.Expected, res)
			}
		})
	}
}

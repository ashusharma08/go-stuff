package main

import "fmt"

func main() {
	fmt.Println(numRollsToTarget(3, 8, 10))
}

// func numRollsToTarget(n int, k int, target int) int {
// 	// Optimization: If target is impossible (too small or too large)
// 	if target < n || target > n*k {
// 		return 0
// 	}

// 	const mod = 1_000_000_007

// 	// Step 1: Create our "History" slice.
// 	// We make it size target + 1 so the indices match the sums (e.g., index 5 = sum 5).
// 	prevWays := make([]int, target+1)

// 	// BASE CASE: With 0 dice, there is 1 way to get a sum of 0.
// 	prevWays[0] = 1

// 	// Step 2: Loop through each die from 1 to n.
// 	for i := 1; i <= n; i++ {
// 		fmt.Println("Dice", i)
// 		// Create a fresh slice for the current die's possibilities.
// 		currentWays := make([]int, target+1)

// 		// Step 3: For every possible total sum 'j' we could reach...
// 		for j := 1; j <= target; j++ {

// 			// Step 4: Check all faces 'f' of the current die (1 to k).
// 			for f := 1; f <= k; f++ {
// 				// If the current sum 'j' is greater than or equal to the face 'f',
// 				// it means we can look back at the history.
// 				if j >= f {
// 					// The number of ways to get sum 'j' with 'i' dice is the sum
// 					// of ways we got (j-f) with 'i-1' dice.
// 					currentWays[j] = (currentWays[j] + prevWays[j-f])
// 				}

// 			}
// 			fmt.Println("_", currentWays)
// 		}
// 		fmt.Println("E", currentWays)
// 		// Step 5: The "Current" results become the "History" for the next die roll.
// 		prevWays = currentWays
// 	}

// 	// Our answer is the number of ways to reach 'target' after n dice.
// 	return prevWays[target]
// }

func numRollsToTarget(n int, k int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1 // 0 dice, sum 0 -> 1 way

	for dice := 1; dice <= n; dice++ {
		newDP := make([]int, target+1)
		prefix := make([]int, target+2) // prefix[s+1] = sum dp[0..s]
		for s := 0; s <= target; s++ {
			prefix[s+1] = prefix[s] + dp[s]
		}

		for s := 1; s <= target; s++ {
			left := s - k
			if left < 0 {
				left = 0
			}
			newDP[s] = prefix[s] - prefix[left]
		}
		fmt.Println(newDP)
		dp = newDP
	}

	return dp[target]
}

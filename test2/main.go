package main

import (
	"fmt"
	"sort"
)

type Tuple struct {
	Packs int
	Combo []int
}

func minPacksToFulfill(orderQty int, packSizes []int) []Tuple {
	maxPackSize := 0
	for _, pack := range packSizes {
		if pack > maxPackSize {
			maxPackSize = pack
		}
	}

	dp := make([]Tuple, orderQty+maxPackSize+1)
	for i := range dp {
		dp[i] = Tuple{orderQty + maxPackSize, []int{}}
	}
	dp[0] = Tuple{0, []int{}}

	for i := 0; i < orderQty+maxPackSize+1; i++ {
		for _, pack := range packSizes {
			if pack <= i {
				if 1+dp[i-pack].Packs < dp[i].Packs {
					newCombo := append([]int(nil), dp[i-pack].Combo...)
					newCombo = append(newCombo, pack)
					dp[i] = Tuple{1 + dp[i-pack].Packs, newCombo}
				}
			}
		}
	}

	return dp[orderQty : orderQty+maxPackSize+1]
}

func main() {
	packSizes := []int{250, 500, 1000, 2000, 5000}
	orderQty := 751

	combinations := minPacksToFulfill(orderQty, packSizes)

	// Filter to get valid combinations and then sort them
	maxPackSize := 0
	for _, pack := range packSizes {
		if pack > maxPackSize {
			maxPackSize = pack
		}
	}

	var validCombinations []Tuple
	for _, combo := range combinations {
		if combo.Packs != orderQty+maxPackSize {
			validCombinations = append(validCombinations, combo)
		}
	}

	// Sort by the sum of packs
	sort.Slice(validCombinations, func(i, j int) bool {
		return sum(validCombinations[i].Combo) < sum(validCombinations[j].Combo)
	})

	// If you only want the best solution, it would be the first entry after sorting
	bestSolution := validCombinations[0]
	fmt.Println("\nBest solution:")
	fmt.Printf("%d packs: %v\n", bestSolution.Packs, bestSolution.Combo)
}

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

/**
You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Example 2:
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
**/

package easy

import (
	"fmt"
	"math"
)

// Dynamic programming
// Complexity: O(n) & Space O(1)
func maxProfit(prices []int) int {
	maxProfitN := 0
	minPrice := math.MaxInt32

	for _, price := range prices {
		minPrice = min(minPrice, price)
		maxProfitN = max(maxProfitN, price-minPrice)
	}
	return maxProfitN
}

func main() {
	res := maxProfit([]int{7, 6, 4, 3, 1})
	fmt.Println(res) // 0
	res = maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println(res) // 5
}

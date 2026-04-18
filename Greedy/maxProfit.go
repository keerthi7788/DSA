package greedy

/*
Input: prices = [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Total profit is 4.
Example 3:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.

Constraints:

1 <= prices.length <= 3 * 104
0 <= prices[i] <= 104
Go
*/
func maxProfit(prices []int) int {

	maxprofit := 0

	for i := 1; i < len(prices); i++ {

		if prices[i] > prices[i-1] {

			maxprofit += prices[i] - prices[i-1]

		}

	}

	return maxprofit

}

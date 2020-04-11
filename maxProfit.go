/*
Say you have an array prices for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

Example 1:

Input: [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
             Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Example 2:

Input: [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
             Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
             engaging multiple transactions at the same time. You must sell before buying again.
Example 3:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
 

Constraints:

1 <= prices.length <= 3 * 10 ^ 4
0 <= prices[i] <= 10 ^ 4
*/


func maxProfit(p []int) int {
    bed := false
    buy := 0
    prof := 0
    n := len(p)
    for i:=0;i<n;i++ {
        //fmt.Printf("Call for i= %d and profit=%d array = %+v\n",i,prof,p)
        if i+1 == n{
           if bed{
                prof += p[i]-buy
            }
            continue
        }
        if p[i] < p[i + 1] {
            if bed {
              continue  
            } 
            //fmt.Printf("Buy for i= %d and profit=%d array = %+v and buy=%d\n",i,prof,p,p[i])
            buy = p[i]
            bed = true
        }
        if p[i] > p[i+1]{
            if !bed{
                continue
            }
            //fmt.Printf("Sell for i= %d and profit=%d array = %+v and sell=%d \n",i,prof,p,p[i])
            prof += p[i]-buy
            bed=false
        }
    }
    return prof
    
}

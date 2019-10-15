func maxProfit(prices []int) int {
    if len(prices) < 1 || prices == nil {
        return 0
    }
    res := 0
    minv := prices[0]
    for i:=1;i<len(prices);i++ {
        res = max(res, prices[i]-minv)
        if minv > prices[i] {
            minv = prices[i]
        }
    }
    return res
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
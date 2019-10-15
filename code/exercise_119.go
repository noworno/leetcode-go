func getRow(rowIndex int) []int {
    res := make([]int, 34)
    last := make([]int, 34)
    res[0] = 1
    last[0] = 1
    for i:=0;i<rowIndex;i++ {
        res[i+1]=1
        for j:=1;j<i+1;j++ {
            res[j] = last[j-1] + last[j]
        }
        res, last = last, res
    }
    return last[:rowIndex+1]
}
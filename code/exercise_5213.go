func minCostToMoveChips(chips []int) int {
    num1 := 0
    num2 := 0
    for _,val := range chips {
        if val % 2 == 0 {
            num1++
        } else {
            num2 ++
        }
    }
    if num1 < num2{
        return num1
    } else {
        return num2
    }
}

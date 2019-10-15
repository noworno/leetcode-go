func removeElement(nums []int, val int) int {
    lastIndex := len(nums)-1
    for i:=lastIndex;i>=0;i-- {
        if nums[i] == val {
            if i != lastIndex {
                nums[lastIndex], nums[i] = nums[i], nums[lastIndex]
            }
            lastIndex--
        }
    }
    return lastIndex+1
}
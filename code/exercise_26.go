func removeDuplicates(nums []int) int {
    lastIndex := 0
    for i:=1; i<len(nums); i++ {
        if nums[i] != nums[lastIndex] {
            lastIndex++
            nums[lastIndex] = nums[i]
        }
    }
    return lastIndex+1
}
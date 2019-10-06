func removeElement(nums []int, val int) int {
    l := len(nums)
    count := 0
    interval := 0
    for i:=0;i<l;i++ {
        if nums[i] == val {
            count++
            interval++
        } else {
            nums[i-interval] = nums[i] 
        }
    }
    return l-count
}

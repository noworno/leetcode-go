// 二分法，查偶数下标的值是否与其相邻位的植相等（注意边界条件）
func singleNonDuplicate(nums []int) int {
    l := 0
    r := len(nums)-1
    index := 0
    for {
        if l==r {
            return nums[r]
        }
        index = (l+r)/2
        if index > 0 && nums[index] == nums[index-1] {
            if (index-2)%2==0 {
                r = index-2
            } else {
                l = index+1
            }
        } else if index < len(nums)-1 && nums[index] == nums[index+1] {
            if (index+2)%2==0 {
                l = index+2
            } else {
                r = index-1
            }
        } else {
            return nums[index]
        }
    }
    return -1
}
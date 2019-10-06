func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
    con := make([]int,1000)
    index :=0
    for _,v1 := range arr1 {
        count := 0
        for _,v2 := range arr2 {
            if v1==v2 {
                count++
                break
            }
        }
        for _,v3 := range arr3 {
            if v1==v3 {
                count++
                break
            }
        }
        if count==2 {
            con[index] = v1
            index++
        }
    }
    return con[:index]
}

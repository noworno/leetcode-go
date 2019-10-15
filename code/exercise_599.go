func findRestaurant(list1 []string, list2 []string) []string {
    m := make(map[string]int)
    res := make([]string, 0)
    leastNum := 2000
    last := 0
    for i,v := range list1 {
        m[v] = i
    }
    for i,v := range list2 {
        if index,ok := m[v]; ok {
            if index+i < leastNum {
                leastNum = index+i
                res = append(res, v)
                last = 0
            } else if index+i == leastNum {
                res = append(res, v)
                last++
            }
        }
    }
    return res[len(res)-1-last:]
}
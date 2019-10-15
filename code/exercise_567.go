func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }
    s1,s2 = s2,s1
    m1 := make(map[byte]int)
    m2 := make(map[byte]int)
    for i:=0;i<len(s2);i++{
        m2[s2[i]] += 1
        m1[s1[i]] += 1
    }
    if check(m1, m2) {
        return true
    }
    for j:=len(s2); j<len(s1); j++ {
        m1[s1[j]] += 1
        m1[s1[j-len(s2)]] -= 1
        if check(m1,m2) {
            return true
        }
    }
    return false
}

func check(a, b map[byte]int) bool {
    for k,v := range a {
        if b[k] != v {
            return false
        }
    }
    return true
}
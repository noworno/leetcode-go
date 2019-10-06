func lengthOfLastWord(s string) int {
    l := len(s)
    count := 0
    for i:=(l-1);i>=0;i-- {
        if s[i]==' '&&count>0 {
            break
        } else if s[i]!=' ' {
            count++
        }
    }
    return count
}

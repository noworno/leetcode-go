func findPoisonedDuration(timeSeries []int, duration int) int {
    res := 0
    lastAttackTime := -1
    for _,v := range timeSeries {
        if lastAttackTime >=0 && v-lastAttackTime<duration {
            res += v-lastAttackTime
        } else {
            res += duration
        }
        lastAttackTime = v
    }
    return res
}
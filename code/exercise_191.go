func hammingWeight(num uint32) int {
	s := strconv.FormatUint(uint64(num), 2)
	count := 0
	for _, c := range s {
		if c == '1' {
			count++
		}
	}
	return count
}
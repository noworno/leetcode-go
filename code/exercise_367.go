func isPerfectSquare(num int) bool {
    if num == 1{
        return true
    }
    z  := float64(num/2)
    n := float64(num)
    precision := 0.0001
    for {
        if z*z == n {
           break
        } else if z*z>n&&(z*z-n)<precision || z*z<n&&(n-z*z)<precision {
            break
        }
        z -= (z * z - n) / (2 * z)
    }
    nz := int(z)
    if nz*nz == num {
        return true
    } 
    return false
}

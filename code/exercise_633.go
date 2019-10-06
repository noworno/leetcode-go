func judgeSquareSum(c int) bool {
    for i:=0;i<=c/2;i++{
        if i*i > c {
            break
        }
        if judgeSquare(c-i*i) {
            return true
        }
    }
    return false
}

func judgeSquare(num int) bool{
    n := float64(num)
    z := 1.0
    precision := 0.0001
    for {
        if z*z==n{
            break
        } else if z*z > n && z*z-n<=precision || z*z < n && n-z*z<=precision {
            break
        }
        z -= (z*z-n) / (2*z)
    }
    nz := int(z)
    if nz*nz == num {
        return true
    }
    return false
}

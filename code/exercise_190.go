func reverseBits(num uint32) uint32 {
    var mask uint32 = 1
    var result uint32 = 0
    for i:=0;i<32;i++{
        result <<= 1 // 因为执行32次，所以先移位，后面加最后才有意义
        result += num&mask
         num >>= 1
    }
    return result
}
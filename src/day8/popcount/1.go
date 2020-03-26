/*
	统计二进制中，1bit数目
*/
package popcount

var pc [256]byte

//匿名函数初始化
/*
	var pc [256]byte = func (pc [256]byte){
		for i := range pc {
			pc[i]= pc[i/2] + byte(i&1)
		}

		return
	}
*/
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

//返回根据设定的x的统计
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

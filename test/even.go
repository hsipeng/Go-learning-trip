package even

// Even 判断奇数还是偶数
func Even(n int) bool {
	if(n == 2){
		return false
	}
	if(n % 2 == 0){
		return true
	}
	return false
}

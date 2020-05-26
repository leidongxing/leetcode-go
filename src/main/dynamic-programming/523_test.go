package dynamic_programming

//暴力破解
func checkSubarraySum1(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		sum := 0
		n := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			n++
			if n >= 2 {
				if k == 0 && sum == 0 {
					return true
				} else if k != 0 && sum%k == 0 {
					return true
				}
			}
		}
	}
	return false
}

//使用map存储
func checkSubarraySum(nums []int, k int) bool {
	l := len(nums)
	hashmap := make(map[int]int)
	hashmap[0] = -1 //对0特殊处理
	sum := 0
	for i := 0; i < l; i++ {
		sum += nums[i]
		if k != 0 {
			sum = sum % k
		}

		if v, ok := hashmap[sum]; ok {
			if i-v > 1 { //必须有2个数
				return true
			}
		} else {
			hashmap[sum] = i
		}
	}
	return false
}

package hash_table

import "testing"

//和可被K整除的子数组

//暴力破解会超时
//同余定理：如果两个整数 a, b 满足 (a-b)%K == 0，那么有 a%K == b%K。
//使用map存储所有  前缀和
func subarraysDivByK(A []int, K int) int {
	record := map[int]int{0: 1}
	sum, ans := 0, 0
	for _, elem := range A {
		sum += elem
		modulus := (sum%K + K) % K
		ans += record[modulus]
		record[modulus]++
	}
	return ans
}

func Test_974(t *testing.T) {
	t.Log(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5)) //7
}

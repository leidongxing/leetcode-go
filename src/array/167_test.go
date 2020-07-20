package array

//两数之和 II - 输入有序数组
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	result := make([]int, 2)
	for left < right {
		if numbers[left]+numbers[right] == target {
			result[0] = left + 1
			result[1] = right + 1
			break
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}
	return result
}

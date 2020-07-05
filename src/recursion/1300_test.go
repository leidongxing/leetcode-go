package recursion

import (
	"sort"
	"testing"
)

//转变数组后最接近目标值的数组和
func findBestValue(arr []int, target int) int {
	n := len(arr)
	sort.Ints(arr)

	//n个最大值都小于target
	if arr[n-1]*n <= target {
		return arr[n-1] // value为最大值
	}

	avg := target / n
	//n个最小值都大于target 填平所有值
	if arr[0]*n >= target {
		//除不尽的情况下判断是否要+1
		if (target%n)*2 > n {
			return avg + 1
		}
		return avg //value为比数组中最小值还小的平均值
	}

	//value为数组中的最小值和最大值之间
	var sum int
	tmpArr := []int{}
	for _, v := range arr {
		if v <= avg {
			//小于平均数的直接相加
			sum += v
		} else {
			tmpArr = append(tmpArr, v)
		}
	}
	//填平剩余的数组tmpArr
	return findBestValue(tmpArr, target-sum)
}

func Test_1300(t *testing.T) {
	t.Log(findBestValue([]int{2, 3, 5}, 10))                              //5
	t.Log(findBestValue([]int{2, 3, 5}, 1))                               //0
	t.Log(findBestValue([]int{4, 9, 3}, 10))                              //3
	t.Log(findBestValue([]int{60864, 25176, 27249, 21296, 20204}, 56803)) //11361
	t.Log(findBestValue([]int{60864, 25176, 27249, 21296, 20204}, 56800)) //11360
	t.Log(findBestValue([]int{60864, 25176, 27249, 21296, 20204}, 56802)) //11360
}

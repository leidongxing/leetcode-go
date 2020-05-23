package array_slice

import "testing"

func TestArrayInit(t *testing.T) {
	//数组中元素初始化为0值
	var arr [3]int
	t.Log(arr[1], arr[2])

	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 3, 4, 5}
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	//使用_表示占位符 idx
	for _, e := range arr {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	//取索引0到3的元素 不包含3
	arr1Sec := arr[:3]
	t.Log(arr1Sec)
}

func TestSliceInit(t *testing.T) {
	//切片可变长 没有长度
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	//len=3 cap=5
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	//s2[3],s2[4]不能被访问 未被初始化
	t.Log(s2[0], s2[1], s2[2])
	//填充s2下一个元素为1
	s2 = append(s2, 1)
	t.Log(s2[3])
	t.Log(len(s2), cap(s2))
}

//切片可变长
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2)
	t.Log(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)

	//a := []int{1, 2, 3, 4}
	//b := []int{1, 2, 3, 4}
	//slice 不能进行比较  slice can only be compared to nil
	//if a == b {
	//	t.Log(1)
	//}
}

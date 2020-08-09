package backtracking

import (
	"strconv"
	"testing"
)

//复原IP地址
var (
	result   []string
	segments []int
)

func restoreIpAddresses(s string) []string {
	segments = make([]int, 4)
	result = []string{}
	backtrackIp(s, 0, 0)
	return result
}

func backtrackIp(s string, segId, index int) {
	if segId == 4 {
		if index == len(s) {
			ipAddr := ""
			for i := 0; i < 4; i++ {
				ipAddr += strconv.Itoa(segments[i])
				if i != 3 {
					ipAddr += "."
				}
			}
			result = append(result, ipAddr)
		}
		return
	}

	//没有找到4段 提前回溯
	if index == len(s) {
		return
	}
	//不能有前导零 00
	if s[index] == '0' {
		segments[segId] = 0
		backtrackIp(s, segId+1, index+1)
	}
	//枚举每一种可能性
	addr := 0
	for segEnd := index; segEnd < len(s); segEnd++ {
		addr = addr*10 + int(s[segEnd]-'0')
		if addr > 0 && addr <= 255 {
			segments[segId] = addr
			backtrackIp(s, segId+1, segEnd+1)
		} else {
			break
		}
	}
}

func Test_93(t *testing.T) {
	t.Log(restoreIpAddresses("25525511135"))
}

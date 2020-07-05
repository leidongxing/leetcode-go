package array

import (
	"testing"
)

//螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	dir := 1
	row, col := 0, 0
	top, right, bottom, left := 0, len(matrix[0])-1, len(matrix)-1, 0
	var result []int
	//上下边界 左右边界
	for top <= bottom && left <= right {
		result = append(result, matrix[row][col])
		switch dir {
		//右运动
		case 1:
			if col == right {
				top++ //上边界下移
				dir = 2
				row++ //向下移动
				continue
			}
			col++
		//下运动
		case 2:
			if row == bottom {
				right-- //右边界左移
				dir = 3
				col-- //向左移动
				continue
			}
			row++
		//左运动
		case 3:
			if col == left {
				bottom-- //下边界上移
				dir = 4
				row-- //向上移动
				continue
			}
			col--
		//上运动
		case 4:
			if row == top {
				left++ //左边界右移
				dir = 1
				col++ //向右移动
				continue
			}
			row--
		}
	}
	return result
}

func Test_54(t *testing.T) {
	t.Log(spiralOrder([][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}))
	t.Log(spiralOrder([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}))
}

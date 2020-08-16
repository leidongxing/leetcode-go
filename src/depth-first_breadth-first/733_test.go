package depth_first_breadth_first

import "testing"

//图像渲染
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	visited := make([][]int, len(image))
	for i := 0; i < len(image); i++ {
		visited[i] = make([]int, len(image[0]))
	}
	draw(image, sr, sc, newColor, image[sr][sc], visited)
	return image
}

func draw(image [][]int, sr int, sc int, newColor int, oldColor int, visited [][]int) {
	image[sr][sc] = newColor
	visited[sr][sc] = 1

	//up
	if sr-1 >= 0 && visited[sr-1][sc] != 1 && image[sr-1][sc] == oldColor {
		draw(image, sr-1, sc, newColor, oldColor, visited)
	}

	//down
	if sr+1 < len(image) && visited[sr+1][sc] != 1 && image[sr+1][sc] == oldColor {
		draw(image, sr+1, sc, newColor, oldColor, visited)
	}

	//left
	if sc-1 >= 0 && visited[sr][sc-1] != 1 && image[sr][sc-1] == oldColor {
		draw(image, sr, sc-1, newColor, oldColor, visited)
	}

	//right
	if sc+1 < len(image[0]) && visited[sr][sc+1] != 1 && image[sr][sc+1] == oldColor {
		draw(image, sr, sc+1, newColor, oldColor, visited)
	}
	return
}

func Test_733(t *testing.T) {
	t.Log(floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))
	t.Log(floodFill([][]int{{0, 0, 0}, {0, 1, 1}}, 1, 1, 1))
	t.Log(floodFill([][]int{{0, 1, 1}, {0, 1, 1}}, 1, 1, 1))
}

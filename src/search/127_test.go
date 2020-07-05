package search

import "testing"

//单词接龙  返回转换序列长度

//双向广度优先 两个队列保存 从两端开始的每一层节点
func ladderLength(beginWord string, endWord string, wordList []string) int {
	result := 0
	startQueue := []string{beginWord}
	endQueue := []string{endWord}
	startUsed := make([]bool, len(wordList))
	endUsed := make([]bool, len(wordList))
	if i := searchIndex(endWord, wordList); i == -1 {
		return 0
	} else {
		//标记目标endWord
		endUsed[i] = true
	}

	for len(startQueue) > 0 {
		result++
		layer := len(startQueue)
		for i := 0; i < layer; i++ {
			for index, word := range wordList {
				if !startUsed[index] && hasOneDiffWord(startQueue[i], word) {
					startQueue = append(startQueue, word)
					startUsed[index] = true
					if endUsed[index] {
						//同一个节点被两个方向搜索均访问到 终止搜索 返回层次
						return result + 1
					}
				}
			}
		}
		startQueue = startQueue[layer:]
		//前后互换  双向搜索
		if len(startQueue) > len(endQueue) {
			startQueue, endQueue = endQueue, startQueue
			startUsed, endUsed = endUsed, startUsed
		}
	}
	return 0
}

//广度优先 队列保存每一层的节点
func ladderLength1(beginWord string, endWord string, wordList []string) int {
	//endWord不在字典中 无法转换
	if searchIndex(endWord, wordList) == -1 {
		return 0
	}
	result := 0
	//记录已经使用word  防止重复访问(出现环)
	used := make([]bool, len(wordList))
	//队列保存当前路径  初始化
	queue := []string{beginWord}
	for len(queue) > 0 {
		//遍历的层深度加一
		result++
		layer := len(queue)
		for i := 0; i < layer; i++ {
			//搜索到endWord
			if queue[i] == endWord {
				return result
			}
			for index, word := range wordList {
				//没有被使用过 且 和目标word只差一个字符  符合要求的一层元素入队
				if !used[index] && hasOneDiffWord(queue[i], word) {
					//入队
					queue = append(queue, word)
					used[index] = true
				}
			}
		}
		//当前元素出队
		queue = queue[layer:]
	}
	return 0
}

func searchIndex(keyword string, wordList []string) int {
	for i := 0; i < len(wordList); i++ {
		if keyword == wordList[i] {
			return i
		}
	}
	return -1
}

func hasOneDiffWord(a, b string) bool {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	if count == 1 {
		return true
	}
	return false
}

func Test_127(t *testing.T) {
	t.Log(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})) //5
	t.Log(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))        //0
}

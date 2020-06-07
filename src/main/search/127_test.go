package search

import "testing"

//单词接龙  返回转换序列长度
func ladderLength1(beginWord string, endWord string, wordList []string) int {
	//endWord不在字典中 无法转换
	if searchIndex(endWord, wordList) == -1 {
		return 0
	}
	result := 0
	//记录已经使用word
	used := make([]bool, len(wordList))
	//队列保存当前路径
	queue := []string{beginWord}
	for len(queue) > 0 {
		result++
		cur := len(queue)
		for i := 0; i < cur; i++ {
			//搜索到endWord
			if queue[i] == endWord {
				return result
			}
			for index, word := range wordList {
				//没有被使用过 且 和目标word只差一个字符
				if !used[index] && hasOneDiffWord(queue[i], word) {
					queue = append(queue, word)
					used[index] = true
				}
			}
		}
		queue = queue[cur:]
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

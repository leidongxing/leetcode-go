package string

import (
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s) //初始化为默认零值""

	s = "hello world" //11
	t.Log(len(s))

	s = "\xE4\xBA\xBB\xFF" //string可以存储任何二进制数据
	t.Log(s)
	t.Log(len(s)) //4
	s = "中"
	t.Log(len(s)) //1

	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0]) //字符集表示
	t.Logf("中 UTF-8 %x", s)      //物理存储
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		//每个输出格式符都是匹配第一个参数
		t.Logf("%[1]c %[1]x", c)
	}
}

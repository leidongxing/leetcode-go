package remote

import (
	//加载远程包 起一个别名 cm
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}

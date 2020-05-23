package _map

import "testing"

//使用map实现set
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	t.Log(len(mySet))

	delete(mySet, 1)

	t.Log(len(mySet))
}

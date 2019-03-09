package set

import "testing"

// 利用map实现go 里面的set结构

func TestSet(t *testing.T) {
	mySet := map[string]bool{}
	mySet["hello"] = true
	mySet["hi"] = true
	str := "hi"
	//判断是否存在
	if mySet[str] {
		t.Log("str is existing")
	} else {
		t.Log("str is not existing")
	}
	// set元素个数
	t.Log(len(mySet))
	// set删除元素
	delete(mySet, "hi")
}

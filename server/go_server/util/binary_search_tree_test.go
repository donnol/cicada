package util

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {

	testData := map[string]interface{}{
		"a": "apple",
		"b": "banana",
		"c": "cat",
		"d": "dog",
		"e": "egg",
		"f": "fox",
		"g": "god",
		"h": "hop",
		"i": "ioe",
		"j": "jog",
		"k": "koq",
	}

	// 新建
	bst := &BinarySearchTree{}
	if bst.root != nil {
		t.Fatalf("bad bst root, %+v", bst)
	}

	// 插入
	for key, value := range testData {
		// 插入顺序不同，生成的树也不同
		if err := bst.Insert(key, value); err != nil {
			t.Fatal(err)
		}
	}

	// 遍历
	bst.InOrderTraverse(func(n *Node) {
		fmt.Printf("%#v\n", n)
	})

	// 打印
	bst.String()

	// 搜索及删除
	for key, value := range testData {
		v, ok := bst.Search(key)
		if !ok || v.(string) != value {
			t.Fatalf("bad value, %+v, %v", v, ok)
		}

		bst.Remove(key)
		_, ok = bst.Search(key)
		if ok {
			t.Fatalf("bad value, %v, %v, %v, %+v", key, value, ok, bst)
		}
	}

	// 结束
	if bst.root != nil {
		t.Fatalf("bad bst root, %+v", bst)
	}
}

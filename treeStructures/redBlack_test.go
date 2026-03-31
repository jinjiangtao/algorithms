package treestructures

import (
	"fmt"
	"testing"
)

func TestRedBlackTree(t *testing.T) {
	// 创建红黑树
	tree := NewRedBlackTree()

	// 测试插入操作
	tree.Insert(10, "value10")
	tree.Insert(5, "value5")
	tree.Insert(15, "value15")
	tree.Insert(3, "value3")
	tree.Insert(7, "value7")
	tree.Insert(12, "value12")
	tree.Insert(18, "value18")

	// 测试中序遍历
	traversal := tree.InorderTraversal()
	expected := []int{3, 5, 7, 10, 12, 15, 18}
	for i, v := range traversal {
		if v != expected[i] {
			t.Errorf("Inorder traversal failed: expected %v, got %v", expected, traversal)
			break
		}
	}

	// 测试查找操作
	node := tree.Search(7)
	if node == tree.Nil || node.Value != "value7" {
		t.Error("Search failed for key 7")
	}

	node = tree.Search(20)
	if node != tree.Nil {
		t.Error("Search should return nil for key 20")
	}

	// 测试删除操作
	if !tree.Delete(15) {
		t.Error("Delete failed for key 15")
	}

	traversal = tree.InorderTraversal()
	expected = []int{3, 5, 7, 10, 12, 18}
	for i, v := range traversal {
		if v != expected[i] {
			t.Errorf("Inorder traversal after delete failed: expected %v, got %v", expected, traversal)
			break
		}
	}

	if tree.Delete(20) {
		t.Error("Delete should return false for non-existent key 20")
	}

	fmt.Println("RedBlackTree test passed!")
}

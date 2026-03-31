package treestructures

// 颜色定义
const (
	Red   = true
	Black = false
)

// Node 红黑树节点
type Node struct {
	Key    int
	Value  interface{}
	Color  bool
	Left   *Node
	Right  *Node
	Parent *Node
}

// RedBlackTree 红黑树结构
type RedBlackTree struct {
	Root *Node
	Nil  *Node
}

// NewRedBlackTree 创建一个新的红黑树
func NewRedBlackTree() *RedBlackTree {
	nilNode := &Node{Color: Black}
	return &RedBlackTree{
		Root: nilNode,
		Nil:  nilNode,
	}
}

// LeftRotate 左旋转操作
func (t *RedBlackTree) LeftRotate(x *Node) {
	y := x.Right
	x.Right = y.Left
	if y.Left != t.Nil {
		y.Left.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == t.Nil {
		t.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

// RightRotate 右旋转操作
func (t *RedBlackTree) RightRotate(x *Node) {
	y := x.Left
	x.Left = y.Right
	if y.Right != t.Nil {
		y.Right.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == t.Nil {
		t.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Right = x
	x.Parent = y
}

// Insert 插入操作
func (t *RedBlackTree) Insert(key int, value interface{}) {
	newNode := &Node{
		Key:    key,
		Value:  value,
		Color:  Red,
		Left:   t.Nil,
		Right:  t.Nil,
		Parent: t.Nil,
	}

	var parent *Node = t.Nil
	current := t.Root

	for current != t.Nil {
		parent = current
		if newNode.Key < current.Key {
			current = current.Left
		} else if newNode.Key > current.Key {
			current = current.Right
		} else {
			// 键已存在，更新值
			current.Value = value
			return
		}
	}

	newNode.Parent = parent
	if parent == t.Nil {
		t.Root = newNode
	} else if newNode.Key < parent.Key {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}

	// 修复红黑树性质
	t.insertFixup(newNode)
}

// insertFixup 修复插入操作后的红黑树性质
func (t *RedBlackTree) insertFixup(z *Node) {
	for z.Parent.Color == Red {
		if z.Parent == z.Parent.Parent.Left {
			y := z.Parent.Parent.Right
			if y.Color == Red {
				z.Parent.Color = Black
				y.Color = Black
				z.Parent.Parent.Color = Red
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Right {
					z = z.Parent
					t.LeftRotate(z)
				}
				z.Parent.Color = Black
				z.Parent.Parent.Color = Red
				t.RightRotate(z.Parent.Parent)
			}
		} else {
			y := z.Parent.Parent.Left
			if y.Color == Red {
				z.Parent.Color = Black
				y.Color = Black
				z.Parent.Parent.Color = Red
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Left {
					z = z.Parent
					t.RightRotate(z)
				}
				z.Parent.Color = Black
				z.Parent.Parent.Color = Red
				t.LeftRotate(z.Parent.Parent)
			}
		}
	}
	t.Root.Color = Black
}

// Delete 删除操作
func (t *RedBlackTree) Delete(key int) bool {
	node := t.Search(key)
	if node == t.Nil {
		return false
	}

	var y, x *Node
	if node.Left == t.Nil || node.Right == t.Nil {
		y = node
	} else {
		y = t.successor(node)
	}

	if y.Left != t.Nil {
		x = y.Left
	} else {
		x = y.Right
	}

	x.Parent = y.Parent
	if y.Parent == t.Nil {
		t.Root = x
	} else if y == y.Parent.Left {
		y.Parent.Left = x
	} else {
		y.Parent.Right = x
	}

	if y != node {
		node.Key = y.Key
		node.Value = y.Value
	}

	if y.Color == Black {
		t.deleteFixup(x)
	}

	return true
}

// deleteFixup 修复删除操作后的红黑树性质
func (t *RedBlackTree) deleteFixup(x *Node) {
	for x != t.Root && x.Color == Black {
		if x == x.Parent.Left {
			w := x.Parent.Right
			if w.Color == Red {
				w.Color = Black
				x.Parent.Color = Red
				t.LeftRotate(x.Parent)
				w = x.Parent.Right
			}
			if w.Left.Color == Black && w.Right.Color == Black {
				w.Color = Red
				x = x.Parent
			} else {
				if w.Right.Color == Black {
					w.Left.Color = Black
					w.Color = Red
					t.RightRotate(w)
					w = x.Parent.Right
				}
				w.Color = x.Parent.Color
				x.Parent.Color = Black
				w.Right.Color = Black
				t.LeftRotate(x.Parent)
				x = t.Root
			}
		} else {
			w := x.Parent.Left
			if w.Color == Red {
				w.Color = Black
				x.Parent.Color = Red
				t.RightRotate(x.Parent)
				w = x.Parent.Left
			}
			if w.Right.Color == Black && w.Left.Color == Black {
				w.Color = Red
				x = x.Parent
			} else {
				if w.Left.Color == Black {
					w.Right.Color = Black
					w.Color = Red
					t.LeftRotate(w)
					w = x.Parent.Left
				}
				w.Color = x.Parent.Color
				x.Parent.Color = Black
				w.Left.Color = Black
				t.RightRotate(x.Parent)
				x = t.Root
			}
		}
	}
	x.Color = Black
}

// Search 查找操作
func (t *RedBlackTree) Search(key int) *Node {
	current := t.Root
	for current != t.Nil && key != current.Key {
		if key < current.Key {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return current
}

// successor 查找节点的后继节点
func (t *RedBlackTree) successor(node *Node) *Node {
	if node.Right != t.Nil {
		return t.minimum(node.Right)
	}
	parent := node.Parent
	for parent != t.Nil && node == parent.Right {
		node = parent
		parent = parent.Parent
	}
	return parent
}

// minimum 查找子树中的最小节点
func (t *RedBlackTree) minimum(node *Node) *Node {
	for node.Left != t.Nil {
		node = node.Left
	}
	return node
}

// InorderTraversal 中序遍历
func (t *RedBlackTree) InorderTraversal() []int {
	result := []int{}
	t.inorderTraversal(t.Root, &result)
	return result
}

func (t *RedBlackTree) inorderTraversal(node *Node, result *[]int) {
	if node != t.Nil {
		t.inorderTraversal(node.Left, result)
		*result = append(*result, node.Key)
		t.inorderTraversal(node.Right, result)
	}
}

package treestructures

import (
	"fmt"
)

// TreeNode 二叉树节点结构
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// BinaryTree 二叉树结构
type BinaryTree struct {
	Root *TreeNode
}

// NewBinaryTree 创建新的二叉树
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{Root: nil}
}

// Insert 向二叉树中插入节点
func (bt *BinaryTree) Insert(value int) {
	bt.Root = insertRecursive(bt.Root, value)
}

// insertRecursive 递归插入节点
func insertRecursive(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return &TreeNode{Value: value}
	}

	if value < node.Value {
		node.Left = insertRecursive(node.Left, value)
	} else if value > node.Value {
		node.Right = insertRecursive(node.Right, value)
	}

	return node
}

// Search 在二叉树中查找节点
func (bt *BinaryTree) Search(value int) *TreeNode {
	return searchRecursive(bt.Root, value)
}

// searchRecursive 递归查找节点
func searchRecursive(node *TreeNode, value int) *TreeNode {
	if node == nil || node.Value == value {
		return node
	}

	if value < node.Value {
		return searchRecursive(node.Left, value)
	}
	return searchRecursive(node.Right, value)
}

// Delete 从二叉树中删除节点
func (bt *BinaryTree) Delete(value int) {
	bt.Root = deleteRecursive(bt.Root, value)
}

// deleteRecursive 递归删除节点
func deleteRecursive(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return nil
	}

	if value < node.Value {
		node.Left = deleteRecursive(node.Left, value)
	} else if value > node.Value {
		node.Right = deleteRecursive(node.Right, value)
	} else {
		// 找到要删除的节点
		// 情况1：叶子节点
		if node.Left == nil && node.Right == nil {
			return nil
		}

		// 情况2：只有一个子节点
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}

		// 情况3：有两个子节点
		// 找到右子树中的最小节点
		successor := findMin(node.Right)
		// 用后继节点的值替换当前节点
		node.Value = successor.Value
		// 删除后继节点
		node.Right = deleteRecursive(node.Right, successor.Value)
	}

	return node
}

// findMin 查找最小节点
func findMin(node *TreeNode) *TreeNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Update 更新节点值
func (bt *BinaryTree) Update(oldValue, newValue int) bool {
	// 首先删除旧值
	bt.Delete(oldValue)
	// 然后插入新值
	bt.Insert(newValue)
	// 检查新值是否存在
	return bt.Search(newValue) != nil
}

// InorderTraversal 中序遍历
func (bt *BinaryTree) InorderTraversal() {
	inorderTraversalRecursive(bt.Root)
	fmt.Println()
}

// inorderTraversalRecursive 递归中序遍历
func inorderTraversalRecursive(node *TreeNode) {
	if node != nil {
		inorderTraversalRecursive(node.Left)
		fmt.Printf("%d ", node.Value)
		inorderTraversalRecursive(node.Right)
	}
}

// PreorderTraversal 前序遍历
func (bt *BinaryTree) PreorderTraversal() {
	preorderTraversalRecursive(bt.Root)
	fmt.Println()
}

// preorderTraversalRecursive 递归前序遍历
func preorderTraversalRecursive(node *TreeNode) {
	if node != nil {
		fmt.Printf("%d ", node.Value)
		preorderTraversalRecursive(node.Left)
		preorderTraversalRecursive(node.Right)
	}
}

// PostorderTraversal 后序遍历
func (bt *BinaryTree) PostorderTraversal() {
	postorderTraversalRecursive(bt.Root)
	fmt.Println()
}

// postorderTraversalRecursive 递归后序遍历
func postorderTraversalRecursive(node *TreeNode) {
	if node != nil {
		postorderTraversalRecursive(node.Left)
		postorderTraversalRecursive(node.Right)
		fmt.Printf("%d ", node.Value)
	}
}

// LevelOrderTraversal 层序遍历
func (bt *BinaryTree) LevelOrderTraversal() {
	if bt.Root == nil {
		return
	}

	queue := []*TreeNode{bt.Root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", node.Value)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	fmt.Println()
}

// Height 计算二叉树高度
func (bt *BinaryTree) Height() int {
	return calculateHeight(bt.Root)
}

// calculateHeight 递归计算高度
func calculateHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := calculateHeight(node.Left)
	rightHeight := calculateHeight(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// Size 计算二叉树节点数量
func (bt *BinaryTree) Size() int {
	return calculateSize(bt.Root)
}

// calculateSize 递归计算节点数量
func calculateSize(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return calculateSize(node.Left) + calculateSize(node.Right) + 1
}

// IsEmpty 检查二叉树是否为空
func (bt *BinaryTree) IsEmpty() bool {
	return bt.Root == nil
}

// Clear 清空二叉树
func (bt *BinaryTree) Clear() {
	bt.Root = nil
}

// Print 打印二叉树结构
func (bt *BinaryTree) Print() {
	printTree(bt.Root, 0)
}

// printTree 递归打印树结构
func printTree(node *TreeNode, level int) {
	if node != nil {
		printTree(node.Right, level+1)
		for i := 0; i < level; i++ {
			fmt.Print("    ")
		}
		fmt.Println(node.Value)
		printTree(node.Left, level+1)
	}
}

// 二叉树的中序遍历递归写法
func InorderTraversal(node *TreeNode) []int {
	result := make([]int, 0)
	var inorder func(node *TreeNode)

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		result = append(result, node.Value)
		inorder(node.Right)
	}

	return result
}

// 二叉树的中序遍历，非递归写法
// 先遍历左边节点
// 再遍历右边节点
// 利用strack 结构存放当前数据
func InorderTraversalStack(node *TreeNode) []int {
	result := make([]int, 0)
	strackNode := make([]*TreeNode, 0)
	curNode := node

	for curNode != nil || len(strackNode) > 0 {

		for curNode != nil {
			strackNode = append(strackNode, curNode)
			curNode = curNode.Left
		}

		curNode = strackNode[len(strackNode)-1]
		strackNode = strackNode[:len(strackNode)-1]
		result = append(result, curNode.Value)

		curNode = curNode.Right
	}

	return result
}

// 二叉树的广度遍历BFS
func LevelOrderTraversal(root *TreeNode) []int {

	result := make([]int, 0)
	if root == nil {
		return result
	}

	queue := make([]TreeNode, 0)
	queue = append(queue, *root)

	if len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		result = append(result, cur.Value)

		if cur.Left != nil {
			queue = append(queue, *cur.Left)
		}

		if cur.Right != nil {
			queue = append(queue, *cur.Right)
		}
	}

	return result
}

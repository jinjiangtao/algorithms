package leetcode

//中序遍历
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

//中序遍历递归法
//中序遍历主要是通过迭代法遍历
func InorderTraversal(root *TreeNode) []int {

	res := make([]int, 0)
	var inoder func(node *TreeNode)

	inoder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inoder(node.Left)
		res = append(res, node.Value)
		inoder(node.Right)
	}

	inoder(root)

	return res
}

//中序遍历迭代法
//遍历迭代法实现
//中序通过栈的辅助结构实现
func InorderTraversalDiedai(root *TreeNode) []int {

	res := make([]int, 0)
	strack := make([]*TreeNode, 0)
	cur := root

	for cur != nil || len(strack) > 0 {
		for cur != nil {
			strack = append(strack, cur)
			cur = cur.Left
		}

		cur = strack[len(strack)-1]
		strack = strack[0 : len(strack)-1]

		res = append(res, cur.Value)
		cur = cur.Right
	}

	return res
}

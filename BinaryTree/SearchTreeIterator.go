package BinaryTree

/*
实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。
调用 next() 将返回二叉搜索树中的下一个最小的数。
 */


// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 成员：
// nums: 【递增】的切片
// root: 根节点
type BSTIterator struct {
	nums []int
	root *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	nums := make([]int, 0)
	// 得到二叉树中序遍历的切片形式
	inorder(root, &nums)
	return BSTIterator{
		nums: nums,
		root: root,
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	// 数组第一个为最小的
	smallest := this.nums[0]
	// 删除第一个元素
	this.nums = this.nums[1:]
	return smallest
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if len(this.nums) > 0 {
		return true
	}
	return false
}

// 中序遍历：左中右
func inorder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	// 左
	inorder(root.Left, nums)
	// 中
	*nums = append(*nums, root.Val)
	// 右
	inorder(root.Right, nums)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
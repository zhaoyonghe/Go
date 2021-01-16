type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    r := maxes(nums)
    return &TreeNode{
        nums[r],
        constructMaximumBinaryTree(nums[0:r]),
        constructMaximumBinaryTree(nums[r+1:]),
    }
}

func maxes(nums []int) int {
    res := 0
    for i := range nums {
        if nums[i] > nums[res] {
            res = i
        }
    }
    return res
}

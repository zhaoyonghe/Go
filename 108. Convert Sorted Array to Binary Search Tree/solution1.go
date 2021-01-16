type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    if len(nums) == 1 {
        return &TreeNode{
            Val: nums[0],
        }
    }
    n := len(nums)
    return &TreeNode{
        nums[n/2],
        sortedArrayToBST(nums[:n/2]),
        sortedArrayToBST(nums[n/2+1:]),
    }
}
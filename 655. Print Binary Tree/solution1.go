type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
    if root == nil {
        return nil
    }
    h := height(root)
    span := powerOfTwo(h) - 1

    var tmp [][]string
    for i := 0; i < h; i++ {
        var row []string
        for j := 0; j < span; j++ {
            row = append(row, "")
        }
        tmp = append(tmp, row)
    }

    fill(root, 0, 0, span, tmp)
    return tmp
}

func fill(root *TreeNode, level, start, end int, template [][]string) {
    if root == nil {
        return
    }
    mid := (end + start) / 2
    num := strconv.Itoa(root.Val) 
    template[level][mid] = num
    fill(root.Left, level+1, start, mid, template)
    fill(root.Right, level+1, mid+1, end, template)
}

func height(root *TreeNode) int {
    if root == nil {
        return 0
    }
    le := height(root.Left)
    ri := height(root.Right)
    if le > ri {
        return le + 1
    }
    return ri + 1
}

func powerOfTwo(k int) int {
    res := 1
    tmp := 2
    for k != 0 {
        if k % 2 != 0 {
            res *= tmp
        }
        tmp *= tmp
        k = k >> 1
    }
    return res
}

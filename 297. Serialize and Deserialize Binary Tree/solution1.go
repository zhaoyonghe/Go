type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return "#"
    }
    s := strconv.Itoa(root.Val)
    le := this.serialize(root.Left)
    ri := this.serialize(root.Right)

    return s + " " + le + " " + ri
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    fields := strings.Fields(data)
    i := 0
    return helper(fields, &i)
}

func helper(fields []string, i *int) *TreeNode {
    if fields[*i] == "#" {
        *i = *i + 1
        return nil
    }
    val, _ := strconv.Atoi(fields[*i])
    *i = *i + 1
    le := helper(fields, i)
    ri := helper(fields, i)
    return &TreeNode{val, le, ri}
}
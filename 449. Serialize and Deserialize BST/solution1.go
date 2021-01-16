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
        return ""
    }
    s := strconv.Itoa(root.Val)
    return s + " " + this.serialize(root.Left) + " " + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    ds := strings.Fields(data)
    var nodes []int
    for _, d := range ds {
        i, _ := strconv.Atoi(d)
        nodes = append(nodes, i)
    }
    i := 0
    return helper(nodes, math.MinInt32, math.MaxInt32, &i)
}

func helper(nodes []int, low, up int, i *int) *TreeNode {
    if len(nodes) == *i || !(low < nodes[*i] && nodes[*i] < up) {
        return nil
    }

    val := nodes[*i]
    res := &TreeNode{Val:val}

    *i++
    res.Left = helper(nodes, low, val, i)
    res.Right = helper(nodes, val, up, i)

    return res
}

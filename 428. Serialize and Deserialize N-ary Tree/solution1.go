type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type Codec struct {
    
}

func Constructor() *Codec {
    return &Codec{}
}

func (this *Codec) serialize(root *Node) string {
    if root == nil {
        return "#"
    }
    s := strconv.Itoa(root.Val) + "," + strconv.Itoa(len(root.Children)) + " "
    for _, c := range root.Children {
        s += this.serialize(c)
    }
    return s
}

func (this *Codec) deserialize(data string) *Node {
    if data == "#" {
        return nil
    }
    fields := strings.Fields(data)
    i := 0
    return helper(fields, &i)
}

func helper(fields []string, i *int) *Node {
    nums := strings.Split(fields[*i], ",")
    val, _ := strconv.Atoi(nums[0])
    childNum, _ := strconv.Atoi(nums[1])
    *i = *i + 1
    var children []*Node
    for childNum > 0 {
        childNum--
        children = append(children, helper(fields, i))
    }
    return &Node{val,children}
}
type List struct {
    Arr []int
    i int
}

func MakeList(size int) *List {
    return &List{
            Arr: make([]int, size), 
            i: size - 1,
        }
}

func (this *List) Add(num int) {
    this.Arr[this.i] = num
    this.i = this.i - 1
}

func findOrder(numCourses int, prerequisites [][]int) []int {
    var g [][]int = make([][]int, numCourses)
    var status []int = make([]int, numCourses)

    for _, pre := range prerequisites {
        g[pre[1]] = append(g[pre[1]], pre[0])
        status[pre[0]] = 0
        status[pre[1]] = 0
    }

    var res *List = MakeList(numCourses)

    for k, v := range status {
        if v == 0 && !dfs(g, status, k, res) {
            return []int{}
        }
    }

    return res.Arr
}

func dfs(g [][]int, status []int, start int, res *List) bool {
    status[start] = 1

    for _, nei := range g[start] {
        var ts int = status[nei]
        if ts == 1 {
            return false
        }
        if ts == 0 {
            if !dfs(g, status, nei, res) {
                return false
            }
        }
    }

    res.Add(start)
    status[start] = 2

    return true
}
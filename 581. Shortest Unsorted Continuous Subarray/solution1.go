type stack []int

func newStack() *stack {
    a := stack([]int{})
    return &a
}

func (st *stack) push(a int) {
    *st = append(*st, a)
}

func (st *stack) pop() {
    *st = (*st)[0:len(*st) - 1]
}

func (st *stack) top() int {
    return (*st)[len(*st) - 1]
}

func (st *stack) empty() bool {
    return len(*st) == 0
}

func (st *stack) clear() {
    *st = []int{}
}

func min(a, b int) int {
    if a > b {
        return b
    } else {
        return a
    }
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func findUnsortedSubarray(nums []int) int {
    n := len(nums)
    if n < 2 {
        return 0
    }

    st := newStack()
    le := n - 1
    ri := 0

    for i := 0; i < n; i++ {
        if st.empty() || nums[i] >= nums[st.top()] {
            st.push(i)
        } else {
            le = min(le, st.top())
            st.pop()
            i--
        }
    }

    st.clear()

    for i := n - 1; i >= 0; i-- {
        if st.empty() || nums[i] <= nums[st.top()] {
            st.push(i)
        } else {
            ri = max(ri, st.top())
            st.pop()
            i++
        }
    }

    if le > ri {
        return 0
    } else {
        return ri - le + 1
    }
}





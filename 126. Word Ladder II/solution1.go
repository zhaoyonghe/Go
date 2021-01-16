func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    if beginWord == endWord {
        return [][]string{
            []string{beginWord},
        }
    }

    words := make(map[string]int)
    for i, w := range wordList {
        words[w] = i
    }
    var beginIndex int
    if i, has := words[beginWord]; !has {
        words[beginWord] = len(wordList)
        beginIndex = len(wordList)
        wordList = append(wordList, beginWord)
    } else {
        beginIndex = i
    }
    
    var endIndex int
    if i, has := words[endWord]; !has {
        return [][]string{}
    } else {
        endIndex = i
    }

    q := []int{beginIndex}
    visited := make([]bool, len(wordList))
    visited[beginIndex] = true
    parents := make([][]int, len(wordList))
    res := 0
    found := false
    for len(q) != 0 {
        var tmpvis []int
        res++
        sz := len(q)
        for i := 0; i < sz; i++ {
            nexts := getNexts(wordList[q[i]], words)
            for _, next := range nexts {
                if next == endIndex {
                    found = true
                }
                if !visited[next] {
                    q = append(q, next)
                    tmpvis = append(tmpvis, next)
                    parents[next] = append(parents[next], q[i])
                }
            }
        }
        if found {
            break
        }
        q = q[sz:]
        for _, val := range tmpvis {
            visited[val] = true
        }
    }
    if !found {
        return [][]string{}
    }

    //fmt.Printf("%v, %v\n", wordList, parents)

    paths := [][]string{}
    get(parents, endIndex, beginIndex, wordList, &[]string{}, &paths)

    return paths
}

func getNexts(cur string, m map[string]int) []int {
    var res []int
    bs := []rune(cur)
    for i := 0; i < len(cur); i++ {
        before := bs[i]
        for replace := 'a'; replace <= 'z'; replace++ {
            bs[i] = replace
            tmp := string(bs)
            i, has := m[tmp]
            if tmp != cur && has {
                res = append(res, i)
            } 
        }
        bs[i] = before
    }
    return res
}

func get(parents [][]int, node int, target int, wordList []string, cur *[]string, res *[][]string) {
    if node == target {
        *cur = append(*cur, wordList[node])
        var tmp []string
        for i := len(*cur) - 1; i >= 0; i-- {
            tmp = append(tmp, (*cur)[i])
        }
        *res = append(*res, tmp)
        *cur = (*cur)[0: len(*cur) - 1]
        return
    }
    dup := make([]bool, len(wordList))
    for _, pa := range parents[node] {
        if dup[pa] {
            continue
        }
        dup[pa] = true
        *cur = append(*cur, wordList[node])
        get(parents, pa, target, wordList, cur, res)
        *cur = (*cur)[0: len(*cur) - 1]        
    }
}

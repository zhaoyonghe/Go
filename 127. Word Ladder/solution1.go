func ladderLength(beginWord string, endWord string, wordList []string) int {
    if beginWord == endWord {
        return 1
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
        return 0
    } else {
        endIndex = i
    }

    q := []int{beginIndex}
    visited := make([]bool, len(wordList))
    visited[beginIndex] = true
    res := 0
    for len(q) != 0 {
        res++
        sz := len(q)
        for i := 0; i < sz; i++ {
            nexts := getNexts(wordList[q[i]], words)
            for _, next := range nexts {
                if next == endIndex {
                    return res + 1
                }
                if !visited[next] {
                    q = append(q, next)
                    visited[next] = true
                }
            }
        }
        q = q[sz:]
    }
    return 0
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
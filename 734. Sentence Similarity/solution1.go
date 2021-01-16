func areSentencesSimilar(words1 []string, words2 []string, pairs [][]string) bool {
    if len(words1) != len(words2) {
        return false
    }

    n := len(words1)
    
    m := make(map[string]map[string]bool)
    for _, pair := range pairs {
        if _, ok := m[pair[0]]; ok {
            m[pair[0]][pair[1]] = true
        } else {
            mm := make(map[string]bool)
            mm[pair[1]] = true
            m[pair[0]] = mm
        }
        if _, ok := m[pair[1]]; ok {
            m[pair[1]][pair[0]] = true
        } else {
            mm := make(map[string]bool)
            mm[pair[0]] = true
            m[pair[1]] = mm
        }
    }

    for i := 0; i < n; i++ {
        if words1[i] != words2[i] {
            {
                mm, ok := m[words1[i]]
                if ok {
                    if _, ok := mm[words2[i]]; ok {
                        continue;
                    }
                }
            }
            {
                mm, ok := m[words2[i]]
                if ok {
                    if _, ok := mm[words1[i]]; ok {
                        continue;
                    }
                }
            }
            return false
        }
    }

    return true
}
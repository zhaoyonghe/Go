type Trie struct {
    exist bool
    next [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{}
}

func indexOf(c rune) int {
    return int(c - 'a')
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
    cur := this
    for _, c := range word {
        if cur.next[indexOf(c)] == nil {
            cur.next[indexOf(c)] = &Trie{}
        }
        cur = cur.next[indexOf(c)]
    }
    cur.exist = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    cur := this
    for _, c := range word {
        if cur.next[indexOf(c)] == nil {
            return false
        }
        cur = cur.next[indexOf(c)]
    }
    return cur.exist
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    cur := this
    for _, c := range prefix {
        if cur.next[indexOf(c)] == nil {
            return false
        }
        cur = cur.next[indexOf(c)]
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
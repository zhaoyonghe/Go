type Lock [4]uint8

func turn(lo Lock, n int, up bool) Lock {
    if up {
        lo[n] = (lo[n] + 1) % 10
    } else {
        lo[n] = (lo[n] + 9) % 10
    }
    return lo
}

func convert(s string) Lock {
    return Lock {uint8(s[0] - '0'), uint8(s[1] - '0'), uint8(s[2] - '0'), uint8(s[3] - '0')}
}

func openLock(deadends []string, target string) int {
    visisted := make(map[Lock]bool)
    for _, d := range deadends {
        visisted[convert(d)] = true
    }

    if _, ok := visisted[[4]uint8{0,0,0,0}]; ok {
        return -1
    }

    tgt := convert(target)

    q := []Lock{{0, 0, 0, 0}}
    res := 0
    for len(q) != 0 {
        sz := len(q)
        for i := 0; i < sz; i++ {
            for j := 0; j < 4; j++ {
                if q[i] == tgt {
                    return res
                }
                up := turn(q[i], j, true)
                if _, ok := visisted[up]; !ok {
                    q = append(q, up)
                    visisted[up] = true
                }
                down := turn(q[i], j, false)
                if _, ok := visisted[down]; !ok {
                    q = append(q, down)
                    visisted[down] = true
                }                
            }
        }
        // clear used
        q = q[sz:]
        res++
    }
    return -1
}
package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/


type node struct {
    id int
    isGateway bool
    links []int
}

type nodePair struct {
    id int
    len int
}

type kill struct {
    first int
    second int
}

type visited map[int]bool

type graph []node

type queue []nodePair

func (q *queue) shift () nodePair {
    temp := (*q)[0]
    *q = (*q)[1:]
    return temp
}

func (q queue) insert (id int, len int) queue {
    return append (q, nodePair{id, len})
}

func (q queue) insertFirst (id int, len int) queue {
    return append (queue{nodePair{id, len}}, q...)
}

func (q queue) empty () bool {
    return len(q) == 0
}

func (n node) findKills (g graph, v visited, k map[kill]bool) []kill{
    links := []kill{}
    q := queue{}
    q = q.insert(n.id, 0)
    for !q.empty(){
        currP := q.shift()
        currLen := currP.len
        currI := currP.id
        v[currI] = true
        for _, link := range g[currI].links {
            //fmt.Fprintf(os.Stderr, "%v - %v\n", currLen, killLen)
            if g[link].isGateway && !k[kill{currI, link}] {
                links = append(links, kill{currI, link})
            }else{
                if _, ok := v[link]; !ok {
                    v[link] = true
                    q = q.insert(link, currLen + 1)
                }
            }
        }
    }
    return links
}

func (n node) neighborG (g graph) bool{
    for _, val := range n.links {
        if g[val].isGateway{
            return true
        }
    }
    return false
}

func (n node) predictAI (g graph, numG int) []int {
    order := []int{}
    currN := n
    gvisited := visited{}
    for len(order) < numG {
        v := visited{}
        q := queue{}
        q = q.insert(currN.id, 0)
        for !q.empty(){
            currP := q.shift()
            currI := currP.id
            currLen := currP.len
            v[currI] = true
            for _, link := range g[currI].links {
                if g[link].isGateway && !gvisited[link]{
                    gvisited[link] = true
                    fmt.Fprintf(os.Stderr, "gateway %d - currI %d\n", link, currI)
                    order = append(order, link)
                    q = queue{}
                    currN = g[currI]
                    break
                }else{
                    if _, ok := v[link]; !ok {
                        v[link] = true
                        if g[link].neighborG(g){
                            q = q.insertFirst(link, currLen + 1)
                        }else{
                            q = q.insert(link, currLen + 1)
                        }
                    }
                }
            }
        }
    }
    return order
}

func getBestKill (kills []kill, indexes, order[]int) int {
    for _, oval := range order{
        for _, ival := range indexes{
            fmt.Fprintf(os.Stderr, "kill[ival] %d - oval %d\n", kills[ival],oval)
            if kills[ival].second == oval{
                return ival
            }
        }
    }
    return 0
}


func main() {
    // N: the total number of nodes in the level, including the gateways
    // L: the number of links
    // E: the number of exit gateways
    var N, L, E int
    fmt.Scan(&N, &L, &E)
    
    g := make(graph, N, N)
    for i, _ := range g {
        g[i] = node{i, false, []int{}}
    }
    
    
    for i := 0; i < L; i++ {
        // N1: N1 and N2 defines a link between these nodes
        var N1, N2 int
        fmt.Scan(&N1, &N2)
        g[N1].links = append(g[N1].links, N2)
        g[N2].links = append(g[N2].links, N1)
    }
    
    for i := 0; i < E; i++ {
        // EI: the index of a gateway node
        var EI int
        fmt.Scan(&EI)
        g[EI].isGateway = true
    }
    fmt.Fprintf(os.Stderr, "%v\n", g)
    
    killed := map[kill]bool{}
    first := true
    var order []int
    
    chk := make(chan []kill)
    chp := make(chan []int)
    
    //GAME LOOP
    for {
        visited := visited{}
        // SI: The index of the node on which the Skynet agent is positioned this turn
        var SI int
        fmt.Scan(&SI)
        
        
        go func(chk chan []kill) {
            chk <- g[SI].findKills(g, visited, killed)
        }(chk)
        if first {
            go func(chp chan []int) {
                chp <- g[SI].predictAI(g, E)
            }(chp)
            first = false
            order = <-chp
        }
        
        kills := <-chk
        
        curr := -1
        indexes := []int{}
        index := 0
        for i, kill := range kills{
            if kill.first == SI {
                index = i
                break
            }
            if curr == kill.first{
                indexes = append(indexes, i)
            }else{
                curr = kill.first
            }
        }
        
        if len(indexes) > 0 {
            index = getBestKill(kills, indexes, order)
            fmt.Fprintf(os.Stderr, "WTF\n%d\n", index)
        }
        
        fmt.Fprintf(os.Stderr, "%v - %v - kills %v - index %v\n%v\n", visited, killed, kills, indexes, order)
        
        kill := kills[index]
        killed[kill] = true
        
        
        // Example: 3 4 are the indices of the nodes you wish to sever the link between
        fmt.Printf("%d %d\n", kill.first, kill.second)
    }
}
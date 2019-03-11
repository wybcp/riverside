package graph

import (
	// "github/wybcp/riverside/data-structure/linked-list/simply-linked-list"
	"container/list"
	"fmt"
)

//Graph 无向图,adjacency list 邻接表实现, go自带的双向链表（可以采用单向链表实现）
type Graph struct {
	gList []*list.List
	// 顶点的个数
	n int
}

//init graphh according to capacity
func newGraph(n int) *Graph {
	g := &Graph{}
	g.n = n
	g.gList = make([]*list.List, n)
	for i := range g.gList {
		g.gList[i] = list.New()
	}
	return g
}

//insert as add edge，一条边存2次
func (g *Graph) addEdge(s int, t int) {
	g.gList[s].PushBack(t)
	g.gList[t].PushBack(s)
}

//BFS search path by 广度优先搜索（Breadth-First-Search）
func (g *Graph) BFS(s int, t int) {
	if s == t {
		return
	}

	//init prev
	prev := make([]int, g.n)
	for index := range prev {
		prev[index] = -1
	}

	//search by queue
	var queue []int
	visited := make([]bool, g.n)
	queue = append(queue, s)

	visited[s] = true
	isFound := false
	for len(queue) > 0 && !isFound {
		// 取值
		top := queue[0]
		// 出队列
		queue = queue[1:]

		linkedlist := g.gList[top]
		for e := linkedlist.Front(); e != nil; e = e.Next() {
			k := e.Value.(int)
			if !visited[k] {
				prev[k] = top
				if k == t {
					isFound = true
					break
				}
				queue = append(queue, k)
				visited[k] = true
			}
		}
	}

	if isFound {
		printPrev(prev, s, t)
	} else {
		fmt.Printf("no path found from %d to %d\n", s, t)
	}

}

//DFS search by DFS
func (g *Graph) DFS(s int, t int) {

	prev := make([]int, g.n)
	for i := range prev {
		prev[i] = -1
	}

	visited := make([]bool, g.n)
	visited[s] = true

	isFound := false
	g.recurse(s, t, prev, visited, isFound)

	printPrev(prev, s, t)
}

//recursivly find path
func (g *Graph) recurse(s int, t int, prev []int, visited []bool, isFound bool) {

	if isFound {
		return
	}

	visited[s] = true

	if s == t {
		// isFound = true
		return
	}

	linkedlist := g.gList[s]
	for e := linkedlist.Front(); e != nil; e = e.Next() {
		k := e.Value.(int)
		if !visited[k] {
			prev[k] = s
			g.recurse(k, t, prev, visited, false)
		}
	}

}

//print path recursively
func printPrev(prev []int, s int, t int) {
	if t == s || prev[t] == -1 {
		fmt.Printf("%d->", t)
	} else {
		printPrev(prev, s, prev[t])
		fmt.Printf("%d->", t)
	}
}

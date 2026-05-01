package main

import (
	"bufio"
	"fmt"
	"os"
)

type Edge struct {
	u, v int
}

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	return &DSU{parent: parent, size: size}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) UnionSet(u, v int) bool {
	root_a := d.Find(u)
	root_b := d.Find(v)

	if root_a == root_b {
		return false
	}

	if d.size[root_a] < d.size[root_b] {
		root_a, root_b = root_b, root_a
	}

	d.parent[root_b] = root_a
	d.size[root_a] += d.size[root_b]
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	edges := make([]Edge, m)

	for i := 0; i < m; i++ {
		fmt.Fscan(in, &edges[i].u, &edges[i].v)
		edges[i].u--
		edges[i].v--
	}

	var q int
	fmt.Fscan(in, &q)

	queries := make([]int, q)
	removed := make([]bool, m)

	for i := 0; i < q; i++ {
		fmt.Fscan(in, &queries[i])
		queries[i]--
		removed[queries[i]] = true
	}

	dsu := NewDSU(n)
	components := n

	for i := 0; i < m; i++ {
		if !removed[i] {
			if dsu.UnionSet(edges[i].u, edges[i].v) {
				components--
			}
		}
	}

	res := make([]int, q)

	for i := q - 1; i >= 0; i-- {
		res[i] = components

		e := edges[queries[i]]
		if dsu.UnionSet(e.u, e.v) {
			components--
		}
	}

	for _, x := range res {
		fmt.Fprint(out, x, " ")
	}

}

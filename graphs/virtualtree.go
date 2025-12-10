package graphs

/*
https://www.luogu.com.cn/record/252532463
*/

import "slices"

func VirtualTree(t *Tree, lca *LCA, a []int) [][2]int {
	slices.SortFunc(a, func(x, y int) int {
		return t.in[x] - t.in[y]
	})
	a = slices.Compact(a)
	n := len(a)
	for i := 0; i < n-1; i++ {
		a = append(a, lca.lca(a[i], a[i+1]))
	}
	slices.SortFunc(a, func(x, y int) int {
		return t.in[x] - t.in[y]
	})
	a = slices.Compact(a)
	var b [][2]int
	for i := 0; i < len(a)-1; i++ {
		b = append(b, [2]int{a[i+1], lca.lca(a[i], a[i+1])})
	}
	return b
}

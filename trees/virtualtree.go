package trees

import "slices"

func VirtualTree(a []int, lca func(int, int) int, dfn func(int) int) [][2]int {
	slices.SortFunc(a, func(x, y int) int { return dfn(x) - dfn(y) })
	m := len(a)
	for i := 0; i < m-1; i++ {
		a = append(a, lca(a[i], a[i+1]))
	}
	slices.SortFunc(a, func(x, y int) int { return dfn(x) - dfn(y) })
	a = slices.Compact(a)
	b := make([][2]int, len(a)-1)
	for i := range b {
		b[i] = [2]int{lca(a[i], a[i+1]), a[i+1]}
	}
	return b
}

/*
Sample usage:
func test() {
	hld := NewHLD(NewTree(0))
	_ = VirtualTree([]int{}, hld.Lca, hld.Dfn)
}
*/

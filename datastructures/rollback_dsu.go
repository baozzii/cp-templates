package datastructures

type RollbackDSU struct {
	fa, sz []int
	st     [][2]int
}

func NewRollbackDSU(n int) *RollbackDSU {
	fa := make([]int, n)
	sz := make([]int, n)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}
	return &RollbackDSU{fa, sz, make([][2]int, 0)}
}

func (dsu *RollbackDSU) Find(u int) int {
	if u == dsu.fa[u] {
		return u
	}
	return dsu.Find(dsu.fa[u])
}

func (dsu *RollbackDSU) Merge(u, v int) bool {
	u = dsu.Find(u)
	v = dsu.Find(v)
	if u == v {
		dsu.st = append(dsu.st, [2]int{-1, -1})
		return false
	}
	if dsu.sz[u] > dsu.sz[v] {
		u, v = v, u
	}
	dsu.sz[v] += dsu.sz[u]
	dsu.fa[u] = v
	dsu.st = append(dsu.st, [2]int{u, v})
	return true
}

func (dsu *RollbackDSU) Same(u, v int) bool {
	return dsu.Find(u) == dsu.Find(v)
}

func (dsu *RollbackDSU) Rollback() {
	u := dsu.st[len(dsu.st)-1][0]
	v := dsu.st[len(dsu.st)-1][1]
	if u == -1 {
		return
	}
	dsu.fa[u] = u
	dsu.sz[v] -= dsu.sz[u]
}

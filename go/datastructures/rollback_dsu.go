package datastructures

type RollbackDSU struct {
	fa, sz []int32
	st     [][2]int32
}

func NewRollbackDSU(n int) *RollbackDSU {
	fa := make([]int32, n)
	sz := make([]int32, n)
	for i := range fa {
		fa[i] = int32(i)
		sz[i] = 1
	}
	return &RollbackDSU{fa, sz, make([][2]int32, 0)}
}

func (dsu *RollbackDSU) Find(_u int) int {
	u := int32(_u)
	for u != dsu.fa[u] {
		u = dsu.fa[u]
	}
	return int(u)
}

func (dsu *RollbackDSU) Merge(u, v int) bool {
	u = dsu.Find(u)
	v = dsu.Find(v)
	if u == v {
		dsu.st = append(dsu.st, [2]int32{-1, -1})
		return false
	}
	if dsu.sz[u] > dsu.sz[v] {
		u, v = v, u
	}
	dsu.sz[v] += dsu.sz[u]
	dsu.fa[u] = int32(v)
	dsu.st = append(dsu.st, [2]int32{int32(u), int32(v)})
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

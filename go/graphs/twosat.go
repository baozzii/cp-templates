package graphs

import . "cp-templates/go/common"

type TwoSat struct {
	n int
	g *Graph[Void]
}

func NewTwoSat(n int) *TwoSat {
	return &TwoSat{n, NewGraph[Void](2 * n)}
}

func (t *TwoSat) AddClause(i int, f bool, j int, g bool) {
	t.g.AddEdge(2*i+Cond(f, 0, 1), 2*j+Cond(g, 1, 0))
	t.g.AddEdge(2*j+Cond(g, 0, 1), 2*i+Cond(f, 1, 0))
}

func (t *TwoSat) Satisfiable() bool {
	id := SccId(t.g)
	for i := range t.n {
		if id[i*2] == id[i*2+1] {
			return false
		}
	}
	return true
}

func (t *TwoSat) Answer() []bool {
	id := SccId(t.g)
	ans := make([]bool, t.n)
	for i := range t.n {
		ans[i] = id[i*2] < id[i*2+1]
	}
	return ans
}

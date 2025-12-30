package datastructures

import . "cp-templates/common"

type LichaoInfo[K Integer, T RealNumber] interface {
	Eval(K) T
}

type LichaoNode[K Integer, T RealNumber, M LichaoInfo[K, T]] struct {
	l, r *LichaoNode[K, T, M]
	f    M
}

type Lichao[K Integer, T RealNumber, M LichaoInfo[K, T]] struct {
	n     K
	t     *LichaoNode[K, T, M]
	def   M
	ismin bool
}

func NewLichao[K Integer, T RealNumber, M LichaoInfo[K, T]](n K, ismin bool, def M) *Lichao[K, T, M] {
	return &Lichao[K, T, M]{n: n, t: nil, def: def, ismin: ismin}
}

func (lc *Lichao[K, T, M]) cmp(a, b T) bool {
	if lc.ismin {
		return a < b
	}
	return a > b
}

func (lc *Lichao[K, T, M]) Insert(f M) {
	lc.InsertSeg(0, lc.n, f)
}

func (lc *Lichao[K, T, M]) InsertSeg(L, R K, f M) {
	var dfs func(p **LichaoNode[K, T, M], l, r K)
	dfs = func(p **LichaoNode[K, T, M], l, r K) {
		if l >= R || r <= L {
			return
		}
		if *p == nil {
			*p = &LichaoNode[K, T, M]{f: lc.def}
		}
		node := *p
		m := (l + r) / 2
		if L <= l && r <= R {
			if lc.cmp(f.Eval(m), node.f.Eval(m)) {
				node.f, f = f, node.f
			}
			if r-l == 1 {
				return
			}
			if lc.cmp(f.Eval(l), node.f.Eval(l)) {
				dfs(&node.l, l, m)
			} else {
				dfs(&node.r, m, r)
			}
		} else {
			if r-l == 1 {
				if lc.cmp(f.Eval(l), node.f.Eval(l)) {
					node.f = f
				}
				return
			}
			dfs(&node.l, l, m)
			dfs(&node.r, m, r)
		}
	}

	dfs(&lc.t, 0, lc.n)
}

func (lc *Lichao[K, T, M]) Query(x K) T {
	var dfs func(p *LichaoNode[K, T, M], l, r K) T
	dfs = func(p *LichaoNode[K, T, M], l, r K) T {
		if p == nil {
			return lc.def.Eval(x)
		}
		res := p.f.Eval(x)
		if r-l == 1 {
			return res
		}
		m := (l + r) / 2
		var sub T
		if x < m {
			sub = dfs(p.l, l, m)
		} else {
			sub = dfs(p.r, m, r)
		}
		if lc.cmp(sub, res) {
			return sub
		}
		return res
	}
	return dfs(lc.t, 0, lc.n)
}

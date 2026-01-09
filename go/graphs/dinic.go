package graphs

import . "cp-templates/go/common"

type MFGraph[Cap Integer] struct {
	n       int32
	posFrom []int32
	posIdx  []int32
	g       [][]mfEdge[Cap]
	level   []int32
	iter    []int32
	que     []int32
}

type mfEdge[Cap Integer] struct {
	to  int32
	rev int32
	cap Cap
}

type MFEdge[Cap Integer] struct {
	from int
	to   int
	cap  Cap
	flow Cap
}

func NewMFGraph[Cap Integer](n int) *MFGraph[Cap] {
	mf := &MFGraph[Cap]{n: int32(n)}
	mf.g = make([][]mfEdge[Cap], n)
	mf.level = make([]int32, n)
	mf.iter = make([]int32, n)
	mf.que = make([]int32, n)
	return mf
}

func (mf *MFGraph[Cap]) Reserve(m int) {
	if m <= 0 {
		return
	}
	if cap(mf.posFrom) < m {
		mf.posFrom = make([]int32, 0, m)
		mf.posIdx = make([]int32, 0, m)
	}
	n := int(mf.n)
	if n > 0 {
		avg := (2*m + n - 1) / n
		if avg > 0 {
			for i := 0; i < n; i++ {
				if cap(mf.g[i]) < avg {
					mf.g[i] = make([]mfEdge[Cap], 0, avg)
				}
			}
		}
	}
}

func (mf *MFGraph[Cap]) AddEdge(from, to int, cap Cap) int {
	fi := int32(from)
	ti := int32(to)

	id := len(mf.posFrom)
	mf.posFrom = append(mf.posFrom, fi)
	mf.posIdx = append(mf.posIdx, int32(len(mf.g[from])))

	fromID := int32(len(mf.g[from]))
	toID := int32(len(mf.g[to]))
	if from == to {
		toID++
	}
	mf.g[from] = append(mf.g[from], mfEdge[Cap]{to: ti, rev: toID, cap: cap})
	mf.g[to] = append(mf.g[to], mfEdge[Cap]{to: fi, rev: fromID, cap: 0})
	return id
}

func (mf *MFGraph[Cap]) GetEdge(i int) MFEdge[Cap] {
	from := int(mf.posFrom[i])
	idx := int(mf.posIdx[i])
	e := mf.g[from][idx]
	re := mf.g[int(e.to)][int(e.rev)]
	return MFEdge[Cap]{
		from: from,
		to:   int(e.to),
		cap:  e.cap + re.cap,
		flow: re.cap,
	}
}

func (mf *MFGraph[Cap]) Edges() []MFEdge[Cap] {
	m := len(mf.posFrom)
	res := make([]MFEdge[Cap], 0, m)
	for i := 0; i < m; i++ {
		res = append(res, mf.GetEdge(i))
	}
	return res
}

func (mf *MFGraph[Cap]) ChangeEdge(i int, newCap, newFlow Cap) {
	from := int(mf.posFrom[i])
	idx := int(mf.posIdx[i])
	e := mf.g[from][idx]
	re := mf.g[int(e.to)][int(e.rev)]
	e.cap = newCap - newFlow
	re.cap = newFlow
}

func (mf *MFGraph[Cap]) Flow(s, t int) Cap {
	return mf.FlowLimit(s, t, Limit[Cap]().Max())
}

func (mf *MFGraph[Cap]) FlowLimit(s, t int, flowLimit Cap) Cap {
	n := int(mf.n)
	level := mf.level
	iter := mf.iter
	q := mf.que

	bfs := func() {
		for i := 0; i < n; i++ {
			level[i] = -1
		}
		level[s] = 0
		head, tail := 0, 0
		q[tail] = int32(s)
		tail++

		for head < tail {
			v := int(q[head])
			head++
			lv := level[v]

			gv := mf.g[v]
			for i := 0; i < len(gv); i++ {
				e := gv[i]
				to := int(e.to)
				if e.cap == 0 || level[to] >= 0 {
					continue
				}
				level[to] = lv + 1
				if to == t {
					return
				}
				q[tail] = e.to
				tail++
			}
		}
	}

	var dfs func(v int32, up Cap) Cap
	dfs = func(v int32, up Cap) Cap {
		if v == int32(s) {
			return up
		}
		res := Cap(0)
		lv := level[v]
		gv := mf.g[v]
		for i := iter[v]; int(i) < len(gv); i++ {
			iter[v] = i
			e := &mf.g[v][i]
			to := e.to
			rev := &mf.g[to][e.rev]

			if lv <= level[to] || rev.cap == 0 {
				continue
			}
			need := up - res
			if need <= 0 {
				break
			}
			if rev.cap < need {
				need = rev.cap
			}

			d := dfs(to, need)
			if d <= 0 {
				continue
			}
			e.cap += d
			rev.cap -= d
			res += d
			if res == up {
				return res
			}
		}
		iter[v] = int32(len(gv))
		level[v] = mf.n
		return res
	}

	flow := Cap(0)
	for flow < flowLimit {
		bfs()
		if level[t] == -1 {
			break
		}
		for i := 0; i < n; i++ {
			iter[i] = 0
		}
		f := dfs(int32(t), flowLimit-flow)
		if f == 0 {
			break
		}
		flow += f
	}
	return flow
}

func (mf *MFGraph[Cap]) MinCut(s int) []bool {
	n := int(mf.n)
	vis := make([]bool, n)

	q := mf.que
	head, tail := 0, 0
	vis[s] = true
	q[tail] = int32(s)
	tail++

	for head < tail {
		v := int(q[head])
		head++
		gv := mf.g[v]
		for i := 0; i < len(gv); i++ {
			e := gv[i]
			if e.cap == 0 {
				continue
			}
			to := int(e.to)
			if !vis[to] {
				vis[to] = true
				q[tail] = e.to
				tail++
			}
		}
	}
	return vis
}

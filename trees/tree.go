package trees

type Tree struct {
	n       int32
	root    int32
	adj     [][]int32
	pa      []int32
	in, out []int32
	dep     []int32
	ord     []int32
}

func NewTree(n int) *Tree {
	return &Tree{int32(n), 0, make([][]int32, n), make([]int32, n), make([]int32, n), make([]int32, n), make([]int32, n), make([]int32, n)}
}

func (t *Tree) AddEdge(u, v int) {
	t.adj[u] = append(t.adj[u], int32(v))
	t.adj[v] = append(t.adj[v], int32(u))
}

func (t *Tree) FromEdges(ed [][]int) {
	for _, e := range ed {
		t.AddEdge(e[0], e[1])
	}
}

func (t *Tree) SetRoot(r int) {
	t.root = int32(r)
}

func (t *Tree) Build() {
	ts := int32(0)
	var dfs func(int32, int32)
	dfs = func(u, p int32) {
		t.pa[u] = p
		t.in[u] = ts
		t.ord[ts] = u
		ts++
		var ch []int32
		for _, v := range t.adj[u] {
			if v != p {
				ch = append(ch, v)
				t.dep[v] = t.dep[u] + 1
				dfs(v, u)
			}
		}
		t.adj[u] = append([]int32{}, ch...)
		t.out[u] = ts
	}
	dfs(t.root, -1)
}

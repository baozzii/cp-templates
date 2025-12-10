package graphs

type Tree struct {
	n       int
	adj     [][]int
	pa      []int
	in, out []int
	dep     []int
	ord     []int
}

func NewTree(n int) *Tree {
	return &Tree{n, make([][]int, n), make([]int, n), make([]int, n), make([]int, n), make([]int, n), make([]int, n)}
}

func (t *Tree) AddEdge(u, v int) {
	t.adj[u] = append(t.adj[u], v)
	t.adj[v] = append(t.adj[v], u)
}

func (t *Tree) Build(r int) {
	ts := 0
	var dfs func(int, int)
	dfs = func(u, p int) {
		t.pa[u] = p
		t.in[u] = ts
		t.ord[ts] = u
		ts++
		var ch []int
		for _, v := range t.adj[u] {
			if v != p {
				ch = append(ch, v)
				t.dep[v] = t.dep[u] + 1
				dfs(v, u)
			}
		}
		t.adj[u] = append([]int{}, ch...)
		t.out[u] = ts
	}
	dfs(r, -1)
}

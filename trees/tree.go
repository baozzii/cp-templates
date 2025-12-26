package trees

type Tree struct {
	n, root int
	adj     [][]int
	pa, dep []int
}

func NewTree(n int) *Tree {
	return &Tree{n, 0, make([][]int, n), make([]int, n), make([]int, n)}
}

func (t *Tree) AddEdge(u, v int) {
	t.adj[u] = append(t.adj[u], v)
	t.adj[v] = append(t.adj[v], u)
}

func (t *Tree) SetRoot(r int) {
	t.root = r
}

func (t *Tree) Build() {
	var dfs func(int, int)
	dfs = func(u, p int) {
		t.pa[u] = p
		for _, v := range t.adj[u] {
			if v != p {
				t.dep[v] = t.dep[u] + 1
				dfs(v, u)
			}
		}
		l := len(t.adj[u])
		for i := 0; i < l; i++ {
			if t.adj[u][i] == p {
				t.adj[u] = append(t.adj[u][:i], t.adj[u][i+1:]...)
				break
			}
		}
	}
	dfs(t.root, -1)
}

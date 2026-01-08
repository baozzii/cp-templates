package strings

type ACAMNode struct {
	fail int32
	son  []int32
}

type ACAM struct {
	tr      []ACAMNode
	sz, sig int32
	prep    bool
}

func NewACAM(sig int) *ACAM {
	return &ACAM{[]ACAMNode{{0, make([]int32, sig)}}, 0, int32(sig), false}
}

func (ac *ACAM) InsertInts(a []int) int {
	u := int32(0)
	for _, w := range a {
		if ac.tr[u].son[w] == 0 {
			ac.tr[u].son[w] = int32(len(ac.tr))
			ac.tr = append(ac.tr, ACAMNode{0, make([]int32, ac.sig)})
		}
		u = ac.tr[u].son[w]
	}
	return int(u)
}

func (ac *ACAM) Size() int {
	if !ac.prep {
		ac.Prepare()
	}
	return int(ac.sz)
}

func (ac *ACAM) Prepare() {
	ac.sz = int32(len(ac.tr))
	ac.prep = true
	var q []int32
	for c := range ac.sig {
		if ac.tr[0].son[c] != 0 {
			q = append(q, ac.tr[0].son[c])
		}
	}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for c := range ac.sig {
			if ac.tr[u].son[c] != 0 {
				ac.tr[ac.tr[u].son[c]].fail = ac.tr[ac.tr[u].fail].son[c]
				q = append(q, ac.tr[u].son[c])
			} else {
				ac.tr[u].son[c] = ac.tr[ac.tr[u].fail].son[c]
			}
		}
	}
}

func (ac *ACAM) InsertString(s string, offset int) int {
	a := make([]int, len(s))
	for i := range a {
		a[i] = int(s[i]) - offset
	}
	return ac.InsertInts(a)
}

func (ac *ACAM) Fail(u int) int {
	if !ac.prep {
		ac.Prepare()
	}
	return int(ac.tr[u].fail)
}

func (ac *ACAM) Next(u, c int) int {
	if !ac.prep {
		ac.Prepare()
	}
	return int(ac.tr[u].son[c])
}

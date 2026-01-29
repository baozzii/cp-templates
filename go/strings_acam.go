package templates

type acam_node struct {
	fail int32
	son  []int32
}

type acam struct {
	tr      []acam_node
	sz, sig int32
	prep    bool
}

func new_acam[T integer](sig T) *acam {
	return &acam{[]acam_node{{0, make([]int32, sig)}}, 0, int32(sig), false}
}

func (ac *acam) insert_ints(a []int) int {
	u := int32(0)
	for _, w := range a {
		if ac.tr[u].son[w] == 0 {
			ac.tr[u].son[w] = int32(len(ac.tr))
			ac.tr = append(ac.tr, acam_node{0, make([]int32, ac.sig)})
		}
		u = ac.tr[u].son[w]
	}
	return int(u)
}

func (ac *acam) size() int {
	if !ac.prep {
		ac.prepare()
	}
	return int(ac.sz)
}

func (ac *acam) prepare() {
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

func (ac *acam) insert_string(s string, offset int) int {
	a := make([]int, len(s))
	for i := range a {
		a[i] = int(s[i]) - offset
	}
	return ac.insert_ints(a)
}

func (ac *acam) fail(u int) int {
	if !ac.prep {
		ac.prepare()
	}
	return int(ac.tr[u].fail)
}

func (ac *acam) next(u, c int) int {
	if !ac.prep {
		ac.prepare()
	}
	return int(ac.tr[u].son[c])
}

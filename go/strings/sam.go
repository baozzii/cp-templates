package strings

/*
https://atcoder.jp/contests/abc433/submissions/71554910
*/

type SAMNode struct {
	len, link int32
	next      []int32
}

type SAM struct {
	sig, last int32
	tr        []SAMNode
}

func (s *SAM) new_node(len, link int32) {
	next := make([]int32, s.sig)
	for i := range next {
		next[i] = -1
	}
	node := SAMNode{len, link, next}
	s.tr = append(s.tr, node)
}

func NewSAM(sig int) *SAM {
	sam := SAM{}
	sam.last = 0
	sam.sig = int32(sig)
	sam.new_node(0, -1)
	return &sam
}

func (s *SAM) Extend(c int) int {
	cur := int32(len(s.tr))
	s.new_node(s.tr[s.last].len+1, -1)
	p := s.last
	for p != -1 && s.tr[p].next[c] == -1 {
		s.tr[p].next[c] = cur
		p = s.tr[p].link
	}
	if p == -1 {
		s.tr[cur].link = 0
	} else {
		q := s.tr[p].next[c]
		if s.tr[p].len+1 == s.tr[q].len {
			s.tr[cur].link = q
		} else {
			clone := int32(len(s.tr))
			s.new_node(s.tr[p].len+1, s.tr[q].link)
			copy(s.tr[len(s.tr)-1].next, s.tr[q].next)
			for p != -1 && s.tr[p].next[c] == q {
				s.tr[p].next[c] = clone
				p = s.tr[p].link
			}
			s.tr[q].link = clone
			s.tr[cur].link = clone
		}
	}
	s.last = cur
	return int(s.last)
}

func (s *SAM) Size() int {
	return len(s.tr)
}

func (s *SAM) Len(u int) int {
	return int(s.tr[u].len)
}

func (s *SAM) Link(u int) int {
	return int(s.tr[u].link)
}

func (s *SAM) Next(u, c int) int {
	return int(s.tr[u].next[c])
}

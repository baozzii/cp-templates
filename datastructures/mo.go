package datastructures

/*
https://www.luogu.com.cn/record/252059884
*/

import (
	"math"
	"slices"
)

type MoInfo[T any] interface {
	addr(int)
	addl(int)
	delr(int)
	dell(int)
	get(int, int) T
}

type MoSolver[T any] struct {
	n                      int
	addr, addl, delr, dell func(int)
	get                    func(int, int) T
	queries                [][3]int
}

func NewMoSolver[T any, M MoInfo[T]](n int, m M) *MoSolver[T] {
	return &MoSolver[T]{n, m.addr, m.addl, m.delr, m.dell, m.get, make([][3]int, 0)}
}

func (s *MoSolver[T]) AddQuery(l, r int) {
	s.queries = append(s.queries, [3]int{l, r, len(s.queries)})
}

func (s *MoSolver[T]) Solve() []T {
	B := s.n/(int(math.Sqrt(float64(len(s.queries)+1)))+1) + 1
	slices.SortFunc(s.queries, func(p1, p2 [3]int) int {
		if p1[0]/B == p2[0]/B {
			if (p1[0]/B)%2 == 1 {
				return p2[1] - p1[1]
			} else {
				return p1[1] - p2[1]
			}
		} else {
			return p1[0]/B - p2[0]/B
		}
	})
	ans := make([]T, len(s.queries))
	cl, cr := 0, 0
	addr := s.addr
	addl := s.addl
	delr := s.delr
	dell := s.dell
	get := s.get
	for _, qu := range s.queries {
		l, r, i := qu[0], qu[1], qu[2]
		for cr < r {
			addr(cr)
			cr++
		}
		for cl > l {
			cl--
			addl(cl)
		}
		for cr > r {
			cr--
			delr(cr)
		}
		for cl < l {
			dell(cl)
			cl++
		}
		ans[i] = get(l, r)
	}
	return ans
}

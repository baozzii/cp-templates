package math

/*
https://judge.yosupo.jp/submission/336318 PolyExp
https://judge.yosupo.jp/submission/336315 PolyInv
https://judge.yosupo.jp/submission/336316 PolyLog
*/

import . "codeforces-go/common"

var rev []int
var roots = []int{0, 1}

type Poly []int

func Dft(a Poly) {
	n := len(a)
	if n <= 1 {
		return
	}
	if len(rev) != n {
		k := Ctz(n) - 1
		for len(rev) < n {
			rev = append(rev, 0)
		}
		rev = rev[:n]
		for i := 0; i < n; i++ {
			rev[i] = (rev[i>>1] >> 1) | ((i & 1) << k)
		}
	}
	for i := 0; i < n; i++ {
		if rev[i] < i {
			a[i], a[rev[i]] = a[rev[i]], a[i]
		}
	}
	if len(roots) < n {
		k := Ctz(len(roots))
		for len(roots) < n {
			roots = append(roots, 0)
		}
		for ; (1 << k) < n; k++ {
			e := Pow(31, 1<<(Ctz(M-1)-k-1), M)
			for i := 1 << (k - 1); i < (1 << k); i++ {
				roots[i<<1] = roots[i]
				roots[i<<1|1] = Mul(roots[i], e)
			}
		}
	}
	for k := 1; k < n; k <<= 1 {
		for i := 0; i < n; i += 2 * k {
			for j := 0; j < k; j++ {
				u := a[i+j]
				v := Mul(a[i+j+k], roots[j+k])
				a[i+j] = Add(u, v)
				a[i+j+k] = Sub(u, v)
			}
		}
	}
}

func Idft(a Poly) {
	n := len(a)
	for i := 0; i < (n-1)/2; i++ {
		a[i+1], a[n-i-1] = a[n-i-1], a[i+1]
	}
	Dft(a)
	inv := Mul(Sub(1, M), Inv(n))
	for i := 0; i < n; i++ {
		a[i] = Mul(a[i], inv)
	}
}

func PolyAdd(a, b Poly) Poly {
	res := make(Poly, max(len(a), len(b)))
	for i := range a {
		res[i] = Add(res[i], a[i])
	}
	for i := range b {
		res[i] = Add(res[i], b[i])
	}
	return res
}

func PolySub(a, b Poly) Poly {
	res := make(Poly, max(len(a), len(b)))
	for i := range a {
		res[i] = Add(res[i], a[i])
	}
	for i := range b {
		res[i] = Sub(res[i], b[i])
	}
	return res
}

func PolyMul(_a, _b Poly) Poly {
	a := make(Poly, len(_a))
	b := make(Poly, len(_b))
	copy(a, _a)
	copy(b, _b)
	n := len(a) + len(b) - 1
	m := 1
	for m < n {
		m <<= 1
	}
	for len(a) < m {
		a = append(a, 0)
	}
	for len(b) < m {
		b = append(b, 0)
	}
	Dft(a)
	Dft(b)
	res := make(Poly, m)
	for i := range res {
		res[i] = Mul(a[i], b[i])
	}
	Idft(res)
	res = res[:n]
	return res
}

func PolyTrunc(_a Poly, m int) Poly {
	a := make(Poly, len(_a))
	copy(a, _a)
	if len(a) > m {
		a = a[:m]
	} else {
		for len(a) < m {
			a = append(a, 0)
		}
	}
	return a
}

func PolyDeriv(_a Poly) Poly {
	if len(_a) == 0 {
		return Poly{}
	}
	f := make(Poly, len(_a)-1)
	for i := range f {
		f[i] = Mul(i+1, _a[i+1])
	}
	return f
}

func PolyInteg(_a Poly) Poly {
	f := make(Poly, len(_a)+1)
	for i := range _a {
		f[i+1] = Mul(_a[i], Inv(i+1))
	}
	return f
}

func PolyInv(_a Poly, m int) Poly {
	f := Poly{Inv(_a[0])}
	k := 1
	for k < m {
		k <<= 1
		f = PolyTrunc(PolyMul(f, PolySub(Poly{2}, PolyMul(PolyTrunc(_a, k), f))), k)
	}
	return PolyTrunc(f, m)
}

func PolyLog(_a Poly, m int) Poly {
	return PolyTrunc(PolyInteg(PolyMul(PolyDeriv(_a), PolyInv(_a, m))), m)
}

func PolyExp(_a Poly, m int) Poly {
	f := Poly{1}
	k := 1
	for k < m {
		k <<= 1
		g := Poly{1}
		g = PolySub(g, PolyLog(f, k))
		g = PolyAdd(g, PolyTrunc(_a, k))
		g = PolyMul(f, g)
		f = PolyTrunc(g, k)
	}
	return PolyTrunc(f, m)
}

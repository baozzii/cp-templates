package math

/*
https://judge.yosupo.jp/submission/334632
*/

type Comb struct {
	fac, ifac []int
	n         int
}

func NewComb() *Comb {
	return &Comb{[]int{1}, []int{1}, 0}
}

func (c *Comb) grow(n int) {
	if n <= c.n {
		return
	}
	for len(c.fac) <= n {
		c.fac = append(c.fac, Mul(c.fac[len(c.fac)-1], len(c.fac)))
		c.ifac = append(c.ifac, 0)
	}
	c.ifac[n] = Inv(c.fac[n])
	for i := n - 1; c.ifac[i] == 0; i-- {
		c.ifac[i] = Mul(c.ifac[i+1], i+1)
	}
	c.n = n
}

func (c *Comb) Fac(n int) int {
	c.grow(n)
	return c.fac[n]
}

func (c *Comb) Ifac(n int) int {
	c.grow(n)
	return c.ifac[n]
}

func (c *Comb) Binom(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	return Mul(c.Fac(n), Mul(c.Ifac(k), c.Ifac(n-k)))
}

func (c *Comb) Catalan(n int) int {
	return Div(c.Binom(2*n, n), n+1)
}

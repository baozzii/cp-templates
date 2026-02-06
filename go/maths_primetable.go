package templates

const (
	PRIMETABLE_MU = 1 << iota
	PRIMETABLE_PHI
	PRIMETABLE_MIN_PRIME_FACTOR
	PRIMETABLE_MAX_PRIME_FACTOR
	PRIMETABLE_FACTORS
)

type primetable struct {
	__isprime                             []bool
	__mu, __phi, __minp, __maxp, __primes []int32
	__factors                             [][]int
}

func (t *primetable) isprime(n int) bool {
	return t.__isprime[n]
}

func (t *primetable) mu(n int) int {
	return int(t.__mu[n])
}

func (t *primetable) phi(n int) int {
	return int(t.__phi[n])
}

func (t *primetable) minp(n int) int {
	return int(t.__minp[n])
}

func (t *primetable) maxp(n int) int {
	return int(t.__maxp[n])
}

func (t *primetable) kth_prime(k int) int {
	return int(t.__primes[k])
}

func (t *primetable) factors(n int) []int {
	return t.__factors[n]
}

func new_primetable(n int, f int) *primetable {
	pt := &primetable{}
	minp := make([]int32, n+1)
	primes := make([]int32, 0)
	var mu, phi, maxp []int32
	var factors [][]int
	isprime := make([]bool, n+1)
	nmu := (f & PRIMETABLE_MU) != 0
	nphi := (f & PRIMETABLE_PHI) != 0
	nminp := (f & PRIMETABLE_MIN_PRIME_FACTOR) != 0
	nmaxp := (f & PRIMETABLE_MAX_PRIME_FACTOR) != 0
	nfac := (f & PRIMETABLE_FACTORS) != 0
	if nmu {
		mu = make([]int32, n+1)
		mu[1] = 1
	}
	if nphi {
		phi = make([]int32, n+1)
		phi[1] = 1
	}
	if nmaxp {
		maxp = make([]int32, n+1)
	}
	if nfac {
		factors = make([][]int, n+1)
	}
	for i := 2; i <= n; i++ {
		if minp[i] == 0 {
			minp[i] = int32(i)
			primes = append(primes, int32(i))
			if nmu {
				mu[i] = -1
			}
			if nphi {
				phi[i] = int32(i - 1)
			}
			if nmaxp {
				maxp[i] = int32(i)
			}
		}
		for _, _p := range primes {
			p := int(_p)
			v := i * p
			if v > n {
				break
			}
			minp[v] = _p

			if i%p == 0 {
				if nphi {
					phi[v] = phi[i] * _p
				}
				if nmu {
					mu[v] = 0
				}
				if nmaxp {
					maxp[v] = max(maxp[i], _p)
				}
				break
			} else {
				if nphi {
					phi[v] = phi[i] * (_p - 1)
				}
				if nmu {
					mu[v] = -mu[i]
				}
				if nmaxp {
					maxp[v] = max(maxp[i], _p)
				}
			}
		}
	}
	if nfac {
		for x := 1; x <= n; x++ {
			for y := x; y <= n; y += x {
				factors[y] = append(factors[y], x)
			}
		}
	}
	for i := 2; i <= n; i++ {
		isprime[i] = minp[i] == int32(i)
	}
	pt.__primes = primes
	pt.__isprime = isprime
	if nmu {
		pt.__mu = mu
	}
	if nphi {
		pt.__phi = phi
	}
	if nminp {
		pt.__minp = minp
	}
	if nmaxp {
		pt.__maxp = maxp
	}
	if nfac {
		pt.__factors = factors
	}
	return pt
}

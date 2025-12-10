package math

type PrimeTable struct {
	mpf []int
	mu  []int
	phi []int
	pr  []int
	n   int
}

func NewPrimeTable(n int) *PrimeTable {
	mpf := make([]int, n+1)
	mu := make([]int, n+1)
	phi := make([]int, n+1)
	primes := make([]int, 0)
	mu[1] = 1
	phi[1] = 1
	if n >= 1 {
		mpf[1] = 1
	}
	for i := 2; i <= n; i++ {
		if mpf[i] == 0 {
			mpf[i] = i
			primes = append(primes, i)
			mu[i] = -1
			phi[i] = i - 1
		}
		for _, p := range primes {
			if p*i > n {
				break
			}
			mpf[p*i] = p
			if i%p == 0 {
				mu[p*i] = 0
				phi[p*i] = phi[i] * p
				break
			} else {
				mu[p*i] = -mu[i]
				phi[p*i] = phi[i] * (p - 1)
			}
		}
	}
	return &PrimeTable{mpf: mpf, mu: mu, phi: phi, pr: primes, n: n}
}

func (pt *PrimeTable) IsPrime(x int) bool {
	return pt.mpf[x] == x && x >= 2
}

func (pt *PrimeTable) Factorize(x int) []int {
	if x <= 1 {
		return nil
	}
	factors := make([]int, 0)
	for x > 1 {
		factors = append(factors, pt.mpf[x])
		x /= pt.mpf[x]
	}
	return factors
}

func (pt *PrimeTable) Get(i int) int {
	return pt.pr[i]
}

func (pt *PrimeTable) Phi(x int) int {
	return pt.phi[x]
}

func (pt *PrimeTable) Mu(x int) int {
	return pt.mu[x]
}

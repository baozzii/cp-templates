package math

type PrimeTable struct {
	mpf []int32
	mu  []int8
	phi []int32
	pr  []int32
	n   int
}

func NewPrimeTable(n int) *PrimeTable {
	mpf := make([]int32, n+1)
	mu := make([]int8, n+1)
	phi := make([]int32, n+1)
	primes := make([]int32, 0)
	mu[1] = 1
	phi[1] = 1
	if n >= 1 {
		mpf[1] = 1
	}
	for i := int32(2); i <= int32(n); i++ {
		if mpf[i] == 0 {
			mpf[i] = i
			primes = append(primes, i)
			mu[i] = -1
			phi[i] = i - 1
		}
		for _, p := range primes {
			if p*i > int32(n) {
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
	return pt.mpf[x] == int32(x) && x >= 2
}

func (pt *PrimeTable) Factorize(x int) []int {
	if x <= 1 {
		return nil
	}
	factors := make([]int, 0)
	for x > 1 {
		factors = append(factors, int(pt.mpf[x]))
		x /= int(pt.mpf[x])
	}
	return factors
}

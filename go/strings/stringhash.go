package strings

import . "cp-templates/go/common"

const M1 = 998244353
const M2 = uint(1e9 + 9)
const B1 = 17603351
const B2 = 87853631

var pw1, pw2 []uint

type StringHashResult struct {
	hash1, hash2 uint
	len          int
}

func StringHashConcat(s1, s2 StringHashResult) StringHashResult {
	return StringHashResult{(s1.hash1 + s2.hash1*pw1[s1.len]) % M1, (s1.hash2 + s2.hash2*pw2[s1.len]) % M2, s1.len + s2.len}
}

func StringHashEquals(s1, s2 StringHashResult) bool {
	return s1.hash1 == s2.hash1 && s1.hash2 == s2.hash2 && s1.len == s2.len
}

type StringHash struct {
	n        int
	hs1, hs2 []uint
}

func grow(n int) {
	if len(pw1) >= n {
		return
	}
	if len(pw1) == 0 {
		pw1 = append(pw1, 1)
		pw2 = append(pw2, 1)
	}
	for len(pw1) < n {
		pw1 = append(pw1, pw1[len(pw1)-1]*B1%M1)
		pw2 = append(pw2, pw2[len(pw2)-1]*B2%M2)
	}
}

func NewStringHashWithInts[T Integer, E ~[]T](s E) *StringHash {
	n := len(s)
	hs1 := make([]uint, n+1)
	hs2 := make([]uint, n+1)
	grow(n + 1)
	for i := range s {
		pw1[i+1] = pw1[i] * B1 % M1
		pw2[i+1] = pw2[i] * B2 % M2
		hs1[i+1] = (hs1[i]*B1 + uint(s[i])) % M1
		hs2[i+1] = (hs2[i]*B2 + uint(s[i])) % M2
	}
	return &StringHash{n, hs1, hs2}
}

func NewStringHashWithString[T ~string](s T) *StringHash {
	n := len(s)
	hs1 := make([]uint, n+1)
	hs2 := make([]uint, n+1)
	grow(n + 1)
	for i := range s {
		pw1[i+1] = pw1[i] * B1 % M1
		pw2[i+1] = pw2[i] * B2 % M2
		hs1[i+1] = (hs1[i]*B1 + uint(s[i])) % M1
		hs2[i+1] = (hs2[i]*B2 + uint(s[i])) % M2
	}
	return &StringHash{n, hs1, hs2}
}

func (s *StringHash) Get(l, r int) StringHashResult {
	hash1 := (s.hs1[r] - s.hs1[l]*pw1[r-l]%M1 + M1) % M1
	hash2 := (s.hs2[r] - s.hs2[l]*pw2[r-l]%M2 + M2) % M2
	return StringHashResult{hash1, hash2, r - l}
}

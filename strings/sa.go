package strings

import (
	"index/suffixarray"
	"unsafe"
)

type SuffixArray struct {
	sa []int32
	rk []int32
	ht []int32
}

func NewSuffixArrayFromBytes(s []byte) *SuffixArray {
	sa := (*struct {
		_  []byte
		sa []int32
	})(unsafe.Pointer(suffixarray.New([]byte(s)))).sa
	rk := make([]int32, len(sa))
	for i := range sa {
		rk[sa[i]] = int32(i)
	}
	ht := make([]int32, len(sa))
	h := 0
	for i, r := range rk {
		if h > 0 {
			h--
		}
		if r > 0 {
			for j := int(sa[r-1]); i+h < len(s) && j+h < len(s) && s[i+h] == s[j+h]; h++ {
			}
		}
		ht[r] = int32(h)
	}
	return &SuffixArray{sa, rk, ht}
}

func NewSuffixArrayFromInts(a []int32) *SuffixArray {
	n := len(a)
	_s := make([]byte, 0, n*4)
	for _, v := range a {
		_s = append(_s, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
	}
	_sa := (*struct {
		_  []byte
		sa []int32
	})(unsafe.Pointer(suffixarray.New([]byte(_s)))).sa
	sa := make([]int32, 0, n)
	for _, p := range _sa {
		if p&3 == 0 {
			sa = append(sa, p>>2)
		}
	}
	rk := make([]int32, len(sa))
	for i := range sa {
		rk[sa[i]] = int32(i)
	}
	ht := make([]int32, len(sa))
	h := 0
	for i, r := range rk {
		if h > 0 {
			h--
		}
		if r > 0 {
			for j := int(sa[r-1]); i+h < len(a) && j+h < len(a) && a[i+h] == a[j+h]; h++ {
			}
		}
		ht[r] = int32(h)
	}
	return &SuffixArray{sa, rk, ht}
}

func (sa *SuffixArray) Get(i int) int {
	return int(sa.sa[i])
}

func (sa *SuffixArray) Rank(i int) int {
	return int(sa.rk[i])
}

func (sa *SuffixArray) Height(i int) int {
	return int(sa.ht[i])
}

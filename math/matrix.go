package math

/*
#include <stdint.h>
#include <stddef.h>
#include <stdlib.h>

uint32_t *mat_add(size_t n, uint32_t M, const uint32_t *a, const uint32_t *b) {
	uint32_t *c = malloc(n * sizeof(uint32_t));
	for (size_t i = 0; i < n; i++) c[i] = (a[i] + b[i]) % M;
	return c;
}
uint32_t *mat_mul(size_t m, size_t n, size_t l, uint32_t M,
                  const uint32_t *a, const uint32_t *b) {
    uint32_t *c = (uint32_t*)calloc(m * l, sizeof(uint32_t));
    for (size_t k = 0; k < n; k++)
        for (size_t i = 0; i < m; i++)
            for (size_t j = 0; j < l; j++) {
                uint64_t add = (uint64_t)a[i * n + k] * b[k * l + j] % M;
                c[i * l + j] = (c[i * l + j] + (uint32_t)add) % M;
            }
    return c;
}
*/
import "C"
import "unsafe"

type Matrix struct {
	m, n uint32
	a    []uint32
}

func (t *Matrix) Get(i, j int) int {
	return int(t.a[i*int(t.n)+j])
}

func (t *Matrix) Set(i, j, v int) {
	t.a[i*int(t.n)+j] = uint32(v)
}

func NewZeroMatrix(m, n int) Matrix {
	return Matrix{uint32(m), uint32(n), make([]uint32, m*n)}
}

func NewIdentityMatrix(n int) Matrix {
	t := Matrix{uint32(n), uint32(n), make([]uint32, n*n)}
	for i := range n {
		t.a[i*n+i] = 1
	}
	return t
}

func MatAdd(a, b Matrix) Matrix {
	if a.m != b.m || a.n != b.n {
		panic(-1)
	}
	p := C.mat_add(C.size_t(a.m*a.n), C.uint32_t(M),
		(*C.uint32_t)(unsafe.Pointer(&a.a[0])),
		(*C.uint32_t)(unsafe.Pointer(&b.a[0])),
	)
	defer C.free(unsafe.Pointer(p))
	c := NewZeroMatrix(int(a.m), int(a.n))
	copy(c.a, unsafe.Slice((*uint32)(unsafe.Pointer(p)), a.m*a.n))
	return c
}

func MatMul(a, b Matrix) Matrix {
	if a.n != b.m {
		panic(-1)
	}
	if a.m == 0 || a.n == 0 || b.n == 0 {
		return NewZeroMatrix(int(a.m), int(b.n))
	}
	p := C.mat_mul(
		C.size_t(a.m),
		C.size_t(a.n),
		C.size_t(b.n),
		C.uint32_t(M),
		(*C.uint32_t)(unsafe.Pointer(&a.a[0])),
		(*C.uint32_t)(unsafe.Pointer(&b.a[0])),
	)
	defer C.free(unsafe.Pointer(p))
	c := NewZeroMatrix(int(a.m), int(b.n))
	copy(c.a, unsafe.Slice((*uint32)(unsafe.Pointer(p)), int(a.m)*int(b.n)))
	return c
}

func MatPow(a Matrix, n int) Matrix {
	if a.n != a.m {
		panic(-1)
	}
	c := NewIdentityMatrix(int(a.n))
	for ; n > 0; a, n = MatMul(a, a), n>>1 {
		if n&1 == 1 {
			c = MatMul(a, c)
		}
	}
	return c
}

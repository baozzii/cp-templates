package templates

/*
#cgo CFLAGS: -O3

#include <stdint.h>

typedef uint32_t u32;
typedef uint64_t u64;
typedef u32* matrix;

const u32 M = 998244353;

void matrix_mul(u32 m, u32 n, u32 l, matrix a, matrix b, matrix c) {
    for (u32 i = 0; i < m; i++) for (u32 j = 0; j < n; j++) {
        u32 x = a[i * n + j];
        for (u32 k = 0; k < l; k++) {
            c[i * l + k] = (u32)(((u64)c[i * l + k] + (u64)x * (u64)b[j * l + k]) % M);
        }
    }
}
*/
import "C"

const MATRIX_THRESHOLD = 10000

type mint_matrix vector[vector[mint]]

func new_identity_mint_matrix(n int) mint_matrix {
	a := vec2[mint](n, n)
	for i := range n {
		a[i][i] = 1
	}
	return mint_matrix(a)
}

func new_zero_mint_matrix(m, n int) mint_matrix {
	return mint_matrix(vec2[mint](m, n))
}

func (a mint_matrix) row() int {
	return len(a)
}

func (a mint_matrix) col() int {
	return len(a[0])
}

func (a mint_matrix) __matrix_mul_go(b mint_matrix) mint_matrix {
	m := a.row()
	n := a.col()
	l := b.col()
	c := new_zero_mint_matrix(m, l)
	for i := range m {
		for j := range n {
			x := a[i][j]
			for k := range l {
				c[i][k] = c[i][k].add(x.mul(b[j][k]))
			}
		}
	}
	return c
}

func (a mint_matrix) __matrix_mul_c(b mint_matrix) mint_matrix {
	m := a.row()
	n := a.col()
	l := b.col()
	_c := make([]C.uint32_t, m*l)
	_a := make([]C.uint32_t, m*n)
	_b := make([]C.uint32_t, n*l)
	for i := range m {
		for j := range n {
			_a[i*n+j] = C.uint32_t(a[i][j])
		}
	}
	for i := range n {
		for j := range l {
			_b[i*l+j] = C.uint32_t(b[i][j])
		}
	}
	C.matrix_mul(C.uint32_t(m), C.uint32_t(n), C.uint32_t(l), &_a[0], &_b[0], &_c[0])
	c := new_zero_mint_matrix(m, l)
	for i := range m {
		for j := range l {
			c[i][j] = mint(_c[i*l+j])
		}
	}
	return c
}

func (a mint_matrix) matrix_mul(b mint_matrix) mint_matrix {
	m := a.row()
	n := a.col()
	l := b.col()
	if m*n*l <= MATRIX_THRESHOLD {
		return a.__matrix_mul_go(b)
	} else {
		return a.__matrix_mul_c(b)
	}
}

func (a mint_matrix) add(b mint_matrix) mint_matrix {
	c := new_zero_mint_matrix(a.row(), a.col())
	for i := range a.row() {
		for j := range a.col() {
			c[i][j] = a[i][j].add(b[i][j])
		}
	}
	return c
}

func (a mint_matrix) sub(b mint_matrix) mint_matrix {
	c := new_zero_mint_matrix(a.row(), a.col())
	for i := range a.row() {
		for j := range a.col() {
			c[i][j] = a[i][j].sub(b[i][j])
		}
	}
	return c
}

func (a mint_matrix) scalar_mul(x mint) mint_matrix {
	c := new_zero_mint_matrix(a.row(), a.col())
	for i := range a.row() {
		for j := range a.col() {
			c[i][j] = a[i][j].mul(x)
		}
	}
	return c
}

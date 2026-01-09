package maths

/*
#include <stddef.h>
#include <stdint.h>
#include <stdbool.h>
#include <string.h>
#include <stdlib.h>
#include <stdio.h>

typedef uint32_t u32;
typedef uint64_t u64;
typedef int32_t i32;
typedef int64_t i64;

enum { R = 23 };
static const u32 M = 998244353, G = 3;
static bool prepared = false;

static u32 root[R + 1], iroot[R + 1], rate2[R - 1], irate2[R - 1], rate3[R - 2], irate3[R - 2];

static u32 norm_mod(u64 x) {
    return x % M;
}

static u32 add_mod(u32 x, u32 y) {
    return x + y >= M ? x + y - M : x + y;
}

static u32 sub_mod(u32 x, u32 y) {
   return x >= y ? x - y : x + M - y;
}

static u32 mul_mod(u32 x, u32 y) {
    return (u64)x * y % M;
}

static u32 pow_mod(u32 x, u64 n) {
    u32 r = 1;
    for (; n; n >>= 1, x = mul_mod(x, x)) if (n & 1) r = mul_mod(r, x);
    return r;
}

static u32 inv_mod(u32 x) {
    return pow_mod(x, M - 2);
}

static u32 countr_zero(u32 n) {
    if (n == 0) return 0;
    u32 r = 0;
    for (; ~n & 1; n >>= 1, r++);
    return r;
}

static u32 bit_ceil(u32 n) {
    u32 m = 1;
    while (m < n) m <<= 1;
    return m;
}

static void prepare() {
    if (prepared) return;
    prepared = true;
    root[R] = pow_mod(G, (M - 1) >> R);
    iroot[R] = inv_mod(root[R]);
    for (int i = R - 1; i >= 0; i--) {
        root[i] = mul_mod(root[i + 1], root[i + 1]);
        iroot[i] = mul_mod(iroot[i + 1], iroot[i + 1]);
    }
    {
        u32 prod = 1, iprod = 1;
        for (int i = 0; i <= R - 2; i++) {
            rate2[i] = mul_mod(root[i + 2], prod);
            irate2[i] = mul_mod(iroot[i + 2], iprod);
            prod = mul_mod(prod, iroot[i + 2]);
            iprod = mul_mod(iprod, root[i + 2]);
        }
    }
    {
        u32 prod = 1, iprod = 1;
        for (int i = 0; i <= R - 3; i++) {
            rate3[i] = mul_mod(root[i + 3], prod);
            irate3[i] = mul_mod(iroot[i + 3], iprod);
            prod = mul_mod(prod, iroot[i + 3]);
            iprod = mul_mod(iprod, root[i + 3]);
        }
    }
}

static void butterfly(u32 *a, u32 n) {
    u32 h = countr_zero(n);
    prepare();

    u32 len = 0;
    while (len < h) {
        if (h - len == 1) {
            u32 p = 1 << (h - len - 1);
            u32 rot = 1;
            for (u32 s = 0; s < (1 << len); s++) {
                u32 offset = s << (h - len);
                for (u32 i = 0; i < p; i++) {
                    u32 l = a[i + offset];
                    u32 r = mul_mod(a[i + offset + p], rot);
                    a[i + offset] = add_mod(l, r);
                    a[i + offset + p] = sub_mod(l, r);
                }
                if (s + 1 != (1 << len)) rot = mul_mod(rot, rate2[countr_zero(~s)]);
            }
            len++;
        } else {
            u32 p = 1 << (h - len - 2);
            u32 rot = 1, imag = root[2];
            for (u32 s = 0; s < (1 << len); s++) {
                u32 rot2 = mul_mod(rot, rot);
                u32 rot3 = mul_mod(rot2, rot);
                u32 offset = s << (h - len);
                for (u32 i = 0; i < p; i++) {
                    u64 mod2 = 1ULL * M * M;
                    u64 a0 = (u64)a[i + offset];
                    u64 a1 = (u64)a[i + offset + p] * rot;
                    u64 a2 = (u64)a[i + offset + 2 * p] * rot2;
                    u64 a3 = (u64)a[i + offset + 3 * p] * rot3;
                    u64 a1na3imag = 1ULL * norm_mod(a1 + mod2 - a3) * imag;
                    u64 na2 = mod2 - a2;
                    a[i + offset] = norm_mod(a0 + a2 + a1 + a3);
                    a[i + offset + 1 * p] = norm_mod(a0 + a2 + (2 * mod2 - (a1 + a3)));
                    a[i + offset + 2 * p] = norm_mod(a0 + na2 + a1na3imag);
                    a[i + offset + 3 * p] = norm_mod(a0 + na2 + (mod2 - a1na3imag));
                }
                if (s + 1 != (1 << len)) rot = mul_mod(rot, rate3[countr_zero(~s)]);
            }
            len += 2;
        }
    }
}

static void butterfly_inv(u32 *a, u32 n) {
    int h = countr_zero(n);
    prepare();

    u32 len = h;
    while (len) {
        if (len == 1) {
            u32 p = 1 << (h - len);
            u32 irot = 1;
            for (u32 s = 0; s < (1 << (len - 1)); s++) {
                u32 offset = s << (h - len + 1);
                for (u32 i = 0; i < p; i++) {
                    u32 l = a[i + offset];
                    u32 r = a[i + offset + p];
                    a[i + offset] = add_mod(l, r);
                    a[i + offset + p] = mul_mod(sub_mod(l, r), irot);
                }
                if (s + 1 != (1 << (len - 1))) irot = mul_mod(irot, irate2[countr_zero(~s)]);
            }
            len--;
        } else {
            u32 p = 1 << (h - len);
            u32 irot = 1, iimag = iroot[2];
            for (u32 s = 0; s < (1 << (len - 2)); s++) {
                u32 irot2 = mul_mod(irot, irot);
                u32 irot3 = mul_mod(irot2, irot);
                u32 offset = s << (h - len + 2);
                for (u32 i = 0; i < p; i++) {
                    u64 a0 = a[i + offset + 0 * p];
                    u64 a1 = a[i + offset + 1 * p];
                    u64 a2 = a[i + offset + 2 * p];
                    u64 a3 = a[i + offset + 3 * p];
                    u32 a2na3iimag = mul_mod(sub_mod(a2, a3), iimag);

                    a[i + offset] = norm_mod(a0 + a1 + a2 + a3);
                    a[i + offset + 1 * p] = mul_mod((a0 + (M - a1) + a2na3iimag), irot);
                    a[i + offset + 2 * p] = mul_mod((a0 + a1 + (M - a2) + (M - a3)), irot2);
                    a[i + offset + 3 * p] = mul_mod((a0 + (M - a1) + (M - a2na3iimag)), irot3);
                }
                if (s + 1 != (1 << (len - 2))) irot = mul_mod(irot, irate3[countr_zero(~s)]);
            }
            len -= 2;
        }
    }
}

static u32 *convolution_naive(u32 *a, u32 n, u32 *b, u32 m) {
    u32 *ans = (u32 *)calloc(n + m - 1, sizeof(u32));
    if (n < m) {
        for (u32 j = 0; j < m; j++) {
            for (u32 i = 0; i < n; i++) {
                ans[i + j] = add_mod(ans[i + j], mul_mod(a[i], b[j]));
            }
        }
    } else {
        for (u32 i = 0; i < n; i++) {
            for (u32 j = 0; j < m; j++) {
                ans[i + j] = add_mod(ans[i + j], mul_mod(a[i], b[j]));
            }
        }
    }
    return ans;
}

static u32 *convolution_fft(u32 *a, u32 n, u32 *b, u32 m) {
    u32 z = bit_ceil(n + m - 1);
    u32 *_a = (u32 *)calloc(z, sizeof(u32));
    memcpy(_a, a, n * sizeof(u32));
    u32 *_b = (u32 *)calloc(z, sizeof(u32));
    memcpy(_b, b, m * sizeof(u32));
    butterfly(_a, z);
    butterfly(_b, z);
    for (u32 i = 0; i < z; i++) _a[i] = mul_mod(_a[i], _b[i]);
    butterfly_inv(_a, z);
    u32 iz = inv_mod(z);
    for (u32 i = 0; i < n + m - 1; i++) _a[i] = mul_mod(_a[i], iz);
    u32 *r = (u32 *)malloc((n + m - 1) * sizeof(u32));
    memcpy(r, _a, (n + m - 1) * sizeof(u32));
    free(_a);
    free(_b);
    return r;
}

static u32 *poly_mul(u32 *a, u32 n, u32 *b, u32 m) {
    if (n <= 60 || m <= 60) return convolution_naive(a, n, b, m);
    else return convolution_fft(a, n, b, m);
}

static u32 *poly_add(u32 *a, u32 n, u32 *b, u32 m) {
    u32 *c = (u32 *)calloc((n > m ? n : m), sizeof(u32));
    for (u32 i = 0; i < n; i++) c[i] = add_mod(c[i], a[i]);
    for (u32 i = 0; i < m; i++) c[i] = add_mod(c[i], b[i]);
    return c;
}

static u32 *poly_sub(u32 *a, u32 n, u32 *b, u32 m) {
    u32 *c = (u32 *)calloc((n > m ? n : m), sizeof(u32));
    for (u32 i = 0; i < n; i++) c[i] = add_mod(c[i], a[i]);
    for (u32 i = 0; i < m; i++) c[i] = sub_mod(c[i], b[i]);
    return c;
}

static u32 *poly_trunc(u32 *a, u32 n, u32 m) {
    u32 *c = (u32 *)calloc(m, sizeof(u32));
    for (u32 i = 0; i < (n < m ? n : m); i++) c[i] = a[i];
    return c;
}

static u32 *poly_deriv(u32 *a, u32 n) {
    u32 *f = (u32 *)calloc((n > 1 ? n - 1 : 1), sizeof(u32));
    if (n <= 1) return f;
    for (u32 i = 0; i < n - 1; i++) f[i] = mul_mod(i + 1, a[i + 1]);
    return f;
}

static u32 *poly_integ(u32 *a, u32 n) {
    u32 *f = (u32 *)calloc(n + 1, sizeof(u32));
    for (u32 i = 0; i < n; i++) f[i + 1] = mul_mod(a[i], inv_mod(i + 1));
    return f;
}

static u32 *poly_inv(u32 *a, u32 n, u32 m) {
    u32 k = 1;
    u32 *f = (u32 *)calloc(1, sizeof(u32));
    f[0] = inv_mod(a[0]);
    while (k < m) {
        u32 nk = k << 1;
        if (nk > m) nk = m;
        u32 *ak = poly_trunc(a, n, nk);
        u32 *t1_full = poly_mul(ak, nk, f, k);
        u32 *t1 = poly_trunc(t1_full, nk + k - 1, nk);
        u32 *two = (u32 *)calloc(nk, sizeof(u32));
        two[0] = 2;
        u32 *t2 = poly_sub(two, nk, t1, nk);
        u32 *t3_full = poly_mul(f, k, t2, nk);
        u32 *fnew = poly_trunc(t3_full, k + nk - 1, nk);
        free(ak);
        free(t1_full);
        free(t1);
        free(two);
        free(t2);
        free(t3_full);
        free(f);
        f = fnew;
        k = nk;
    }
    if (k != m) {
        u32 *res = poly_trunc(f, k, m);
        free(f);
        return res;
    }
    return f;
}

static u32 *poly_log(u32 *a, u32 n, u32 m) {
    u32 dn = (n > 1 ? n - 1 : 1);
    u32 *da = poly_deriv(a, n);
    u32 *inva = poly_inv(a, n, m);
    u32 *prod = poly_mul(da, dn, inva, m);
    u32 prodn = dn + m - 1;
    u32 *integ = poly_integ(prod, prodn);
    u32 *res = poly_trunc(integ, prodn + 1, m);
    free(da);
    free(inva);
    free(prod);
    free(integ);
    return res;
}

static u32 *poly_exp(u32 *a, u32 n, u32 m) {
    u32 k = 1;
    u32 *f = (u32 *)calloc(1, sizeof(u32));
    f[0] = 1;
    while (k < m) {
        u32 nk = k << 1;
        if (nk > m) nk = m;
        u32 *logf = poly_log(f, k, nk);
        u32 *one = (u32 *)calloc(nk, sizeof(u32));
        one[0] = 1;
        u32 *t = poly_sub(one, nk, logf, nk);
        u32 *ak = poly_trunc(a, n, nk);
        u32 *t2 = poly_add(t, nk, ak, nk);
        u32 *t3_full = poly_mul(f, k, t2, nk);
        u32 *fnew = poly_trunc(t3_full, k + nk - 1, nk);
        free(logf);
        free(one);
        free(t);
        free(ak);
        free(t2);
        free(t3_full);
        free(f);
        f = fnew;
        k = nk;
    }
    if (k != m) {
        u32 *res = poly_trunc(f, k, m);
        free(f);
        return res;
    }
    return f;
}
*/
import "C"
import (
	. "cp-templates/go/common"
	"runtime"
	"unsafe"
)

type Poly[T Integer] []T

func to_c_array[T Integer](a Poly[T]) (*C.uint32_t, []C.uint32_t) {
	_a := make([]C.uint32_t, len(a))
	for i := range a {
		_a[i] = C.uint32_t(uint32(a[i]))
	}
	return (*C.uint32_t)(unsafe.Pointer(&_a[0])), _a
}

func to_go_poly[T Integer](a *C.uint32_t, n int) Poly[T] {
	defer C.free(unsafe.Pointer(a))
	b := unsafe.Slice((*C.uint32_t)(unsafe.Pointer(a)), n)
	_a := make(Poly[T], n)
	for i := range _a {
		_a[i] = T(b[i])
	}
	return _a
}

func PolyMul[T Integer](a, b Poly[T]) Poly[T] {
	n, m := len(a), len(b)
	if n == 0 || m == 0 {
		return Poly[T]{}
	}
	ca, _a := to_c_array(a)
	cb, _b := to_c_array(b)
	res := to_go_poly[T](C.poly_mul(ca, C.uint32_t(n), cb, C.uint32_t(m)), n+m-1)
	runtime.KeepAlive(_a)
	runtime.KeepAlive(_b)
	return res
}

func (a *Poly[T]) Mul(b Poly[T]) Poly[T] {
	*a = PolyMul(*a, b)
	return *a
}

func PolyAdd[T Integer](a, b Poly[T]) Poly[T] {
	n, m := len(a), len(b)
	ca, _a := to_c_array(a)
	cb, _b := to_c_array(b)
	res := to_go_poly[T](C.poly_add(ca, C.uint32_t(n), cb, C.uint32_t(m)), n+m-1)
	runtime.KeepAlive(_a)
	runtime.KeepAlive(_b)
	return res
}

func (a *Poly[T]) Add(b Poly[T]) Poly[T] {
	*a = PolyAdd(*a, b)
	return *a
}

func PolySub[T Integer](a, b Poly[T]) Poly[T] {
	n, m := len(a), len(b)
	ca, _a := to_c_array(a)
	cb, _b := to_c_array(b)
	res := to_go_poly[T](C.poly_sub(ca, C.uint32_t(n), cb, C.uint32_t(m)), n+m-1)
	runtime.KeepAlive(_a)
	runtime.KeepAlive(_b)
	return res
}

func (a *Poly[T]) Sub(b Poly[T]) Poly[T] {
	*a = PolySub(*a, b)
	return *a
}

func PolyTrunc[T Integer](a Poly[T], m int) Poly[T] {
	n := len(a)
	ca, _a := to_c_array(a)
	res := to_go_poly[T](C.poly_trunc(ca, C.uint32_t(n), C.uint32_t(m)), m)
	runtime.KeepAlive(_a)
	return res
}

func (a *Poly[T]) Trunc(m int) Poly[T] {
	*a = PolyTrunc(*a, m)
	return *a
}

func PolyDeriv[T Integer](a Poly[T]) Poly[T] {
	n := len(a)
	ca, _a := to_c_array(a)
	res := to_go_poly[T](C.poly_deriv(ca, C.uint32_t(n)), max(1, n-1))
	runtime.KeepAlive(_a)
	return res
}

func (a *Poly[T]) Deriv() Poly[T] {
	*a = PolyDeriv(*a)
	return *a
}

func PolyInteg[T Integer](a Poly[T]) Poly[T] {
	n := len(a)
	ca, _a := to_c_array(a)
	res := to_go_poly[T](C.poly_integ(ca, C.uint32_t(n)), n+1)
	runtime.KeepAlive(_a)
	return res
}

func (a *Poly[T]) Integ() Poly[T] {
	*a = PolyInteg(*a)
	return *a
}

func PolyInv[T Integer](a Poly[T], m int) Poly[T] {
	n := len(a)
	ca, _a := to_c_array(a)
	res := to_go_poly[T](C.poly_inv(ca, C.uint32_t(n), C.uint32_t(m)), m)
	runtime.KeepAlive(_a)
	return res
}

func (a *Poly[T]) Inv(m int) Poly[T] {
	*a = PolyInv(*a, m)
	return *a
}

func PolyLog[T Integer](a Poly[T], m int) Poly[T] {
	n := len(a)
	ca, _a := to_c_array(a)
	res := to_go_poly[T](C.poly_log(ca, C.uint32_t(n), C.uint32_t(m)), m)
	runtime.KeepAlive(_a)
	return res
}

func (a *Poly[T]) Log(m int) Poly[T] {
	*a = PolyLog(*a, m)
	return *a
}

func PolyExp[T Integer](a Poly[T], m int) Poly[T] {
	n := len(a)
	ca, _a := to_c_array(a)
	res := to_go_poly[T](C.poly_exp(ca, C.uint32_t(n), C.uint32_t(m)), m)
	runtime.KeepAlive(_a)
	return res
}

func (a *Poly[T]) Exp(m int) Poly[T] {
	*a = PolyExp(*a, m)
	return *a
}

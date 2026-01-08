package math

/*
#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>

typedef uint32_t u32;
typedef uint64_t u64;
typedef long long i64;

static const u32 M = 998244353u;
static const u32 PRIMITIVE_ROOT = 3u;

typedef struct {
    u32 mod;
    u64 im;
} barrett;

static inline barrett barrett_init(u32 mod) {
    barrett b;
    b.mod = mod;
    b.im = (u64)((~(u64)0) / mod);
    return b;
}

static inline u32 barrett_reduce_u64(const barrett* b, u64 x) {
    u64 q = (u64)(((unsigned __int128)x * b->im) >> 64);
    u64 r = x - q * b->mod;
    if (r >= b->mod) r -= b->mod;
    if (r >= b->mod) r -= b->mod;
    return (u32)r;
}

static inline u32 mod_add(u32 a, u32 b) {
    u32 s = a + b;
    if (s >= M) s -= M;
    return s;
}

static inline u32 mod_sub(u32 a, u32 b) {
    return (a >= b) ? (a - b) : (a + M - b);
}

static const barrett BR = { M, (u64)((~(u64)0) / M) };

static inline u32 mod_mul(u32 a, u32 b) {
    u64 z = (u64)a * b;
    return barrett_reduce_u64(&BR, z);
}

static u32 mod_pow(u32 a, u64 e) {
    u32 r = 1;
    while (e) {
        if (e & 1) r = mod_mul(r, a);
        a = mod_mul(a, a);
        e >>= 1;
    }
    return r;
}
static inline u32 mod_inv(u32 a) {
    return mod_pow(a, (u64)M - 2);
}

static inline int ctz_u32(u32 x) {
    return __builtin_ctz(x);
}

static inline u32 bit_ceil_u32(u32 x) {
    if (x <= 1) return 1;
    return 1u << (32 - __builtin_clz(x - 1));
}

enum { RANK2 = 23 };

typedef struct {
    u32 root[RANK2 + 1];
    u32 iroot[RANK2 + 1];
    u32 rate2[RANK2 - 1];  // size = rank2-2+1 = 22
    u32 irate2[RANK2 - 1];
    u32 rate3[RANK2 - 2];  // size = rank2-3+1 = 21
    u32 irate3[RANK2 - 2];
    int inited;
} fft_info;

static fft_info INFO;

static void fft_init_once(void) {
    if (INFO.inited) return;
    INFO.inited = 1;

    INFO.root[RANK2]  = mod_pow(PRIMITIVE_ROOT, ((u64)M - 1) >> RANK2);
    INFO.iroot[RANK2] = mod_inv(INFO.root[RANK2]);

    for (int i = RANK2 - 1; i >= 0; --i) {
        INFO.root[i]  = mod_mul(INFO.root[i + 1],  INFO.root[i + 1]);
        INFO.iroot[i] = mod_mul(INFO.iroot[i + 1], INFO.iroot[i + 1]);
    }

    {
        u32 prod = 1, iprod = 1;
        for (int i = 0; i <= RANK2 - 2; i++) {
            INFO.rate2[i]  = mod_mul(INFO.root[i + 2],  prod);
            INFO.irate2[i] = mod_mul(INFO.iroot[i + 2], iprod);
            prod  = mod_mul(prod,  INFO.iroot[i + 2]);
            iprod = mod_mul(iprod, INFO.root[i + 2]);
        }
    }
    {
        u32 prod = 1, iprod = 1;
        for (int i = 0; i <= RANK2 - 3; i++) {
            INFO.rate3[i]  = mod_mul(INFO.root[i + 3],  prod);
            INFO.irate3[i] = mod_mul(INFO.iroot[i + 3], iprod);
            prod  = mod_mul(prod,  INFO.iroot[i + 3]);
            iprod = mod_mul(iprod, INFO.root[i + 3]);
        }
    }
}

static void butterfly(u32* a, int n) {
    fft_init_once();
    assert((n & (n - 1)) == 0);
    int h = ctz_u32((u32)n);

    int len = 0;
    while (len < h) {
        if (h - len == 1) {
            int p = 1 << (h - len - 1);
            u32 rot = 1;
            for (int s = 0; s < (1 << len); s++) {
                int offset = s << (h - len);
                for (int i = 0; i < p; i++) {
                    u32 l = a[i + offset];
                    u32 r = mod_mul(a[i + offset + p], rot);
                    a[i + offset] = mod_add(l, r);
                    a[i + offset + p] = mod_sub(l, r);
                }
                if (s + 1 != (1 << len)) {
                    int idx = ctz_u32(~(u32)s);
                    rot = mod_mul(rot, INFO.rate2[idx]);
                }
            }
            len++;
        } else {
            int p = 1 << (h - len - 2);
            u32 rot = 1;
            u32 imag = INFO.root[2]
            for (int s = 0; s < (1 << len); s++) {
                u32 rot2 = mod_mul(rot, rot);
                u32 rot3 = mod_mul(rot2, rot);
                int offset = s << (h - len);
                for (int i = 0; i < p; i++) {
                    u64 a0 = a[i + offset];
                    u64 a1 = (u64)mod_mul(a[i + offset + p], rot);
                    u64 a2 = (u64)mod_mul(a[i + offset + 2 * p], rot2);
                    u64 a3 = (u64)mod_mul(a[i + offset + 3 * p], rot3);
                    u32 t13 = mod_sub((u32)a1, (u32)a3);
                    u32 a1na3imag = mod_mul(t13, imag);
                    u32 na2 = (a2 == 0) ? 0 : (u32)(M - (u32)a2);

                    // a0+a2+a1+a3, etc.
                    a[i + offset] = barrett_reduce_u64(&BR, a0 + a2 + a1 + a3);

                    // a0+a2-(a1+a3)
                    u32 s02 = mod_add((u32)a0, (u32)a2);
                    u32 s13 = mod_add((u32)a1, (u32)a3);
                    a[i + offset + 1 * p] = mod_sub(s02, s13);

                    // a0 - a2 + (a1-a3)*imag
                    u32 d02 = mod_add((u32)a0, na2);
                    a[i + offset + 2 * p] = mod_add(d02, a1na3imag);
                    a[i + offset + 3 * p] = mod_sub(d02, a1na3imag);
                }
                if (s + 1 != (1 << len)) {
                    int idx = ctz_u32(~(u32)s);
                    rot = mod_mul(rot, INFO.rate3[idx]);
                }
            }
            len += 2;
        }
    }
}

static void butterfly_inv(u32* a, int n) {
    fft_init_once();
    assert((n & (n - 1)) == 0);
    int h = ctz_u32((u32)n);

    int len = h;
    while (len) {
        if (len == 1) {
            int p = 1 << (h - len);
            u32 irot = 1;
            for (int s = 0; s < (1 << (len - 1)); s++) {
                int offset = s << (h - len + 1);
                for (int i = 0; i < p; i++) {
                    u32 l = a[i + offset];
                    u32 r = a[i + offset + p];
                    a[i + offset] = mod_add(l, r);
                    u32 diff = mod_sub(l, r);
                    a[i + offset + p] = mod_mul(diff, irot);
                }
                if (s + 1 != (1 << (len - 1))) {
                    int idx = ctz_u32(~(u32)s);
                    irot = mod_mul(irot, INFO.irate2[idx]);
                }
            }
            len--;
        } else {
            int p = 1 << (h - len);
            u32 irot = 1;
            u32 iimag = INFO.iroot[2];
            for (int s = 0; s < (1 << (len - 2)); s++) {
                u32 irot2 = mod_mul(irot, irot);
                u32 irot3 = mod_mul(irot2, irot);
                int offset = s << (h - len + 2);
                for (int i = 0; i < p; i++) {
                    u32 a0 = a[i + offset + 0 * p];
                    u32 a1 = a[i + offset + 1 * p];
                    u32 a2 = a[i + offset + 2 * p];
                    u32 a3 = a[i + offset + 3 * p];

                    // a2na3iimag = (a2 - a3) * iimag
                    u32 t23 = mod_sub(a2, a3);
                    u32 a2na3iimag = mod_mul(t23, iimag);

                    a[i + offset] = barrett_reduce_u64(&BR, (u64)a0 + a1 + a2 + a3);

                    // (a0 - a1 + a2na3iimag) * irot
                    u32 t1 = mod_add(mod_sub(a0, a1), a2na3iimag);
                    a[i + offset + 1 * p] = mod_mul(t1, irot);

                    // (a0 + a1 - a2 - a3) * irot2
                    u32 t2 = mod_sub(mod_add(a0, a1), mod_add(a2, a3));
                    a[i + offset + 2 * p] = mod_mul(t2, irot2);

                    // (a0 - a1 - a2na3iimag) * irot3
                    u32 t3 = mod_sub(mod_sub(a0, a1), a2na3iimag);
                    a[i + offset + 3 * p] = mod_mul(t3, irot3);
                }
                if (s + 1 != (1 << (len - 2))) {
                    int idx = ctz_u32(~(u32)s);
                    irot = mod_mul(irot, INFO.irate3[idx]);
                }
            }
            len -= 2;
        }
    }
}

static u32* convolution_naive(const u32* a, int n, const u32* b, int m, int* out_len) {
    if (n == 0 || m == 0) { *out_len = 0; return NULL; }
    int L = n + m - 1;
    u32* c = (u32*)calloc((size_t)L, sizeof(u32));

    if (n < m) {
        for (int j = 0; j < m; j++) {
            u32 bj = b[j];
            for (int i = 0; i < n; i++) {
                c[i + j] = mod_add(c[i + j], mod_mul(a[i], bj));
            }
        }
    } else {
        for (int i = 0; i < n; i++) {
            u32 ai = a[i];
            for (int j = 0; j < m; j++) {
                c[i + j] = mod_add(c[i + j], mod_mul(ai, b[j]));
            }
        }
    }
    *out_len = L;
    return c;
}

static u32* convolution_fft(const u32* a, int n, const u32* b, int m, int* out_len) {
    if (n == 0 || m == 0) { *out_len = 0; return NULL; }
    int need = n + m - 1;
    u32 z = bit_ceil_u32((u32)need);

    assert(((u64)M - 1) % z == 0);
    assert(z <= (1u << RANK2));

    u32* fa = (u32*)calloc((size_t)z, sizeof(u32));
    u32* fb = (u32*)calloc((size_t)z, sizeof(u32));
    if (!fa || !fb) { free(fa); free(fb); return NULL; }

    memcpy(fa, a, (size_t)n * sizeof(u32));
    memcpy(fb, b, (size_t)m * sizeof(u32));

    butterfly(fa, (int)z);
    butterfly(fb, (int)z);

    for (u32 i = 0; i < z; i++) fa[i] = mod_mul(fa[i], fb[i]);

    butterfly_inv(fa, (int)z);

    u32 iz = mod_inv(z);
    u32* c = (u32*)malloc((size_t)need * sizeof(u32));
    if (!c) { free(fa); free(fb); return NULL; }

    for (int i = 0; i < need; i++) c[i] = mod_mul(fa[i], iz);

    free(fa);
    free(fb);
    *out_len = need;
    return c;
}

u32* convolution_mod(const u32* a, int n, const u32* b, int m, int* out_len) {
    if (!a || !b || n < 0 || m < 0) { if (out_len) *out_len = 0; return NULL; }
    if (n == 0 || m == 0) { *out_len = 0; return NULL; }

    int z = n + m - 1;
    assert(z <= (1 << RANK2));

    if (n < m) { int t = n; n = m; m = t; const u32* tmp = a; a = b; b = tmp; } // optional swap
    if (m <= 60) return convolution_naive(a, n, b, m, out_len);
    return convolution_fft(a, n, b, m, out_len);
}
*/

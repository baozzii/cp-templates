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
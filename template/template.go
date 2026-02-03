// https://github.com/baozzii/cp-templates

package main

import (
	"cmp"
	"fmt"
	"io"
	"math"
	"math/bits"
	"os"
	"reflect"
	"slices"
	"strconv"
	"unsafe"
)

func ctz[T integer](x T) int {
	return bits.TrailingZeros(uint(x))
}

func clz[T integer](x T) int {
	return bits.LeadingZeros(uint(x))
}

func popcount[T integer](x T) int {
	return bits.OnesCount(uint(x))
}

func lowbit[T integer](x T) T {
	y := uint(x)
	return T(y & -y)
}

func highbit[T integer](x T) T {
	if x == 0 {
		return x
	}
	return T(1) << (63 - clz(uint(x)))
}

type void struct{}

type numeric_limits[T integer] struct{}

func limit[T integer]() numeric_limits[T] {
	return struct{}{}
}

func (numeric_limits[T]) max() T {
	var z T
	if (^z) < 0 {
		b := uint(unsafe.Sizeof(z) * 8)
		u := uint(1)<<(b-1) - 1
		return T(u)
	} else {
		return ^z
	}
}

func (numeric_limits[T]) min() T {
	var z T
	if (^z) < 0 {
		b := uint(unsafe.Sizeof(z) * 8)
		u := uint(1) << (b - 1)
		return T(-int(u))
	}
	return 0
}

func to_string[T any](e T) string {
	return fmt.Sprintf("%v", e)
}

func to_int[T integer](s string) T {
	x, _ := strconv.Atoi(s)
	return T(x)
}

func ckmax[T cmp.Ordered](x *T, y T) {
	*x = max(*x, y)
}

func ckmin[T cmp.Ordered](x *T, y T) {
	*x = min(*x, y)
}

func sum[T complex, E ~[]T](v E) T {
	var s T
	for _, w := range v {
		s += w
	}
	return s
}

func presum[T complex, E ~[]T](v E) E {
	p := make(E, len(v)+1)
	for i, w := range v {
		p[i+1] = p[i] + w
	}
	return p
}

func count[T comparable, E ~[]T](v E, e T) int {
	cnt := 0
	for _, w := range v {
		if w == e {
			cnt++
		}
	}
	return cnt
}

func cond[T any](cond bool, x, y T) T {
	if cond {
		return x
	}
	return y
}

func abs[T real](x T) T {
	if x < T(0) {
		return -x
	}
	return x
}

func gcd[T integer](x, y T) T {
	if x < 0 || y < 0 {
		return gcd(abs(x), abs(y))
	}
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func lcm[T integer](x, y T) T {
	return x / gcd(x, y) * y
}

func pow[S, T integer](x S, n T, m S) S {
	r := S(1) % m
	for ; n > 0; n, x = n>>1, x*x%m {
		if n%2 == 1 {
			r = r * x % m
		}
	}
	return r
}

func exgcd[T integer](a, b T) (T, T, T) {
	x0, y0 := T(1), T(0)
	x1, y1 := T(0), T(1)
	for b != 0 {
		q := a / b
		a, b = b, a-q*b
		x0, x1 = x1, x0-q*x1
		y0, y1 = y1, y0-q*y1
	}
	return a, x0, y0
}

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type real interface {
	integer |
		~float32 | ~float64
}

type complex interface {
	real |
		~complex64 | ~complex128
}

type umap[T comparable, K any] map[T]K

func (s *umap[T, K]) erase(x T) {
	delete(*s, x)
}

func (s *umap[T, K]) contains(x T) bool {
	_, ok := (*s)[x]
	return ok
}

func (s *umap[T, K]) size() int {
	return len(*s)
}

func (s *umap[T, K]) empty() bool {
	return s.size() == 0
}

func (s *umap[T, K]) clear() {
	clear(*s)
}

func (s *umap[T, K]) keys() vector[T] {
	b := make(vector[T], 0, s.size())
	for v := range *s {
		b.push_back(v)
	}
	return b
}

func (s *umap[T, K]) values() vector[K] {
	b := make(vector[K], 0, s.size())
	for _, v := range *s {
		b.push_back(v)
	}
	return b
}

type uset[T comparable] map[T]struct{}

func (s *uset[T]) insert(x T) {
	(*s)[x] = struct{}{}
}

func (s *uset[T]) erase(x T) {
	delete(*s, x)
}

func (s *uset[T]) contains(x T) bool {
	_, ok := (*s)[x]
	return ok
}

func (s *uset[T]) size() int {
	return len(*s)
}

func (s *uset[T]) empty() bool {
	return s.size() == 0
}

func (s *uset[T]) clear() {
	clear(*s)
}

func (s *uset[T]) keys() vector[T] {
	b := make(vector[T], 0, s.size())
	for v := range *s {
		b.push_back(v)
	}
	return b
}

func (s *uset[T]) copy() uset[T] {
	t := make(uset[T])
	t.union(*s)
	return t
}

func (s *uset[T]) union(t uset[T]) {
	for v := range t {
		s.insert(v)
	}
}

func (s *uset[T]) intersect(t uset[T]) {
	ns := make(uset[T])
	for v := range t {
		if s.contains(v) {
			ns.insert(v)
		}
	}
	*s = ns
}

func union[T comparable](s, t uset[T]) uset[T] {
	ns := s.copy()
	ns.union(t)
	return ns
}

func intersect[T comparable](s, t uset[T]) uset[T] {
	ns := make(uset[T])
	for v := range s {
		if t.contains(v) {
			ns.insert(v)
		}
	}
	return ns
}

type vector[T any] []T

func (v *vector[T]) size() int {
	return len(*v)
}

func (v *vector[T]) empty() bool {
	return len(*v) == 0
}

func (v *vector[T]) clear() {
	*v = (*v)[:0]
}

func (v *vector[T]) push_back(x T) {
	*v = append(*v, x)
}

func (v *vector[T]) pop_back() {
	*v = (*v)[:len(*v)-1]
}

func (v *vector[T]) resize(x int) {
	if v.size() > x {
		*v = (*v)[:x]
	} else {
		*v = append(*v, make([]T, x-v.size())...)
	}
}

func (v *vector[T]) back() *T {
	return &((*v)[v.size()-1])
}

func (v *vector[T]) front() *T {
	return &((*v)[0])
}

func (v *vector[T]) erase(x, y int) {
	*v = slices.Delete(*v, x, y)
}

func (v *vector[T]) insert(x int, w ...T) {
	*v = slices.Insert(*v, x, w...)
}

func (v *vector[T]) reverse() {
	n := v.size()
	for i := 0; i < n/2; i++ {
		(*v)[i], (*v)[n-i-1] = (*v)[n-i-1], (*v)[i]
	}
}

func (v *vector[T]) fill(w T) {
	for i := range *v {
		(*v)[i] = w
	}
}

func (v *vector[T]) copy() vector[T] {
	return slices.Clone(*v)
}

func (v *vector[T]) slice(i, j int) vector[T] {
	w := (*v)[i:j]
	return w.copy()
}

func (v *vector[T]) scan(read func(...any)) {
	for _, w := range *v {
		read(w)
	}
}

func (v *vector[T]) concat(w vector[T]) {
	for _, c := range w {
		v.push_back(c)
	}
}

func concat[T any](a, b vector[T]) vector[T] {
	na := a.copy()
	na.concat(b)
	return na
}

func vec1[T any, S integer](n ...S) vector[T] {
	if len(n) == 0 {
		return make(vector[T], 0)
	} else {
		return make(vector[T], n[0])
	}
}

func vec2[T any, S integer](n ...S) vector[vector[T]] {
	if len(n) == 0 {
		return make(vector[vector[T]], 0)
	} else {
		res := make(vector[vector[T]], n[0])
		for i := range res {
			res[i] = vec1[T](n[1:]...)
		}
		return res
	}
}

func vec3[T any, S integer](n ...S) vector[vector[vector[T]]] {
	if len(n) == 0 {
		return make(vector[vector[vector[T]]], 0)
	} else {
		res := make(vector[vector[vector[T]]], n[0])
		for i := range res {
			res[i] = vec2[T](n[1:]...)
		}
		return res
	}
}

func to_umap[T comparable, E ~[]T](v E) umap[T, int] {
	cnt := make(umap[T, int])
	for _, w := range v {
		cnt[w]++
	}
	return cnt
}

func to_uset[T comparable, E ~[]T](v E) uset[T] {
	cnt := make(uset[T])
	for _, w := range v {
		cnt.insert(w)
	}
	return cnt
}

type scannable interface {
	scan(func(...any))
}

type writable interface {
	format() []byte
}

type fastio struct {
	in         io.Reader
	out        io.Writer
	rbuf, wbuf []byte
	i, n       int
	fpc        int
}

func new_fastio(in io.Reader, out io.Writer) *fastio {
	return &fastio{in, out, make([]byte, 4096), make([]byte, 0), 0, 0, -1}
}

func new_std_fastio() *fastio {
	return new_fastio(os.Stdin, os.Stdout)
}

func (fastio *fastio) __read_byte() byte {
	if fastio.i == fastio.n {
		fastio.n, _ = fastio.in.Read(fastio.rbuf)
		if fastio.n == 0 {
			return 0
		}
		fastio.i = 0
	}
	b := fastio.rbuf[fastio.i]
	fastio.i++
	return b
}

func (fastio *fastio) read(ptrs ...any) {
	var rd func(v reflect.Value)
	rd = func(v reflect.Value) {
		for i := 0; i < v.Len(); i++ {
			elem := v.Index(i)
			if elem.Kind() == reflect.Slice {
				rd(elem)
			} else {
				fastio.read(elem.Addr().Interface())
			}
		}
	}
	read_unsigned := func() uint {
		var x uint
		b := fastio.__read_byte()
		for ; '0' > b || b > '9'; b = fastio.__read_byte() {
			if b == 0 {
				return x
			}
		}
		for ; '0' <= b && b <= '9'; b = fastio.__read_byte() {
			x = x*10 + uint(b&15)
		}
		return x
	}
	read_signed := func() int {
		var y uint
		var x int
		neg := false
		b := fastio.__read_byte()
		for ; '0' > b || b > '9'; b = fastio.__read_byte() {
			if b == 0 {
				return x
			}
			if b == '-' {
				neg = true
			}
		}
		for ; '0' <= b && b <= '9'; b = fastio.__read_byte() {
			y = y*10 + uint(b&15)
		}
		if neg {
			if y == math.MaxInt+1 {
				x = math.MinInt
			} else {
				x = -int(y)
			}
		} else {
			x = int(y)
		}
		return x
	}
	read_float := func() float64 {
		b := fastio.__read_byte()
		var s []byte
		for ; b == ' ' || b == '\n' || b == '\r' || b == '\t'; b = fastio.__read_byte() {
		}
		for ; !(b == ' ' || b == '\n' || b == '\r' || b == '\t' || b == 0); b = fastio.__read_byte() {
			s = append(s, b)
		}
		w, _ := strconv.ParseFloat(string(s), 64)
		return w
	}
	for _, p := range ptrs {
		switch v := any(p).(type) {
		case *uint:
			*v = read_unsigned()
		case *uint8:
			*v = uint8(read_unsigned())
		case *uint16:
			*v = uint16(read_unsigned())
		case *uint32:
			*v = uint32(read_unsigned())
		case *uint64:
			*v = uint64(read_unsigned())
		case *int:
			*v = read_signed()
		case *int8:
			*v = int8(read_signed())
		case *int16:
			*v = int16(read_signed())
		case *int32:
			*v = int32(read_signed())
		case *int64:
			*v = int64(read_signed())
		case *float32:
			*v = float32(read_float())
		case *float64:
			*v = read_float()
		case *string:
			{
				b := fastio.__read_byte()
				var s []byte
				for ; b == ' ' || b == '\n' || b == '\r'; b = fastio.__read_byte() {
				}
				for ; !(b == ' ' || b == '\n' || b == '\r' || b == 0); b = fastio.__read_byte() {
					s = append(s, b)
				}
				*v = string(s)
			}
		default:
			if v, ok := p.(scannable); ok {
				v.scan(fastio.read)
			} else {
				rv := reflect.ValueOf(p)
				if rv.Kind() == reflect.Ptr && (rv.Elem().Kind() == reflect.Slice || rv.Elem().Kind() == reflect.Array) {
					rd(rv.Elem())
				}
			}
		}
	}
}

func (fastio *fastio) write(a ...any) {
	unsigned_to_string := func(v uint64) []byte {
		var s []byte
		if v == 0 {
			return []byte{'0'}
		}
		for v > 0 {
			s = append(s, '0'+byte(v%10))
			v /= 10
		}
		for i := 0; i < len(s)/2; i++ {
			s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
		}
		return s
	}
	signed_to_string := func(v int64) []byte {
		if v == 0 {
			return []byte{'0'}
		}
		if v == math.MinInt64 {
			return []byte("-9223372036854775808")
		}
		neg := v < 0
		if neg {
			v = -v
		}
		var s []byte
		for v > 0 {
			s = append(s, '0'+byte(v%10))
			v /= 10
		}
		if neg {
			s = append(s, '-')
		}
		for i := 0; i < len(s)/2; i++ {
			s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
		}
		return s
	}

	for i, p := range a {
		if i != 0 {
			fastio.wbuf = append(fastio.wbuf, ' ')
		}
		switch v := p.(type) {
		case uint:
			fastio.wbuf = append(fastio.wbuf, unsigned_to_string(uint64(v))...)
		case uint8:
			fastio.wbuf = append(fastio.wbuf, unsigned_to_string(uint64(v))...)
		case uint16:
			fastio.wbuf = append(fastio.wbuf, unsigned_to_string(uint64(v))...)
		case uint32:
			fastio.wbuf = append(fastio.wbuf, unsigned_to_string(uint64(v))...)
		case uint64:
			fastio.wbuf = append(fastio.wbuf, unsigned_to_string(v)...)

		case int:
			fastio.wbuf = append(fastio.wbuf, signed_to_string(int64(v))...)
		case int8:
			fastio.wbuf = append(fastio.wbuf, signed_to_string(int64(v))...)
		case int16:
			fastio.wbuf = append(fastio.wbuf, signed_to_string(int64(v))...)
		case int32:
			fastio.wbuf = append(fastio.wbuf, signed_to_string(int64(v))...)
		case int64:
			fastio.wbuf = append(fastio.wbuf, signed_to_string(v)...)

		case float32:
			fastio.wbuf = append(fastio.wbuf, []byte(strconv.FormatFloat(float64(v), 'f', fastio.fpc, 64))...)
		case float64:
			fastio.wbuf = append(fastio.wbuf, []byte(strconv.FormatFloat(v, 'f', fastio.fpc, 64))...)
		case string:
			fastio.wbuf = append(fastio.wbuf, v...)
		default:
			if v, ok := p.(writable); ok {
				fastio.wbuf = append(fastio.wbuf, v.format()...)
			} else {
				rv := reflect.ValueOf(p)
				if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
					if rv.Type().Elem().Kind() == reflect.Slice {
						for j := 0; j < rv.Len(); j++ {
							if j+1 == rv.Len() {
								fastio.write(rv.Index(j).Interface())
							} else {
								fastio.writeln(rv.Index(j).Interface())
							}
						}
					} else {
						for j := 0; j < rv.Len(); j++ {
							if j != 0 {
								fastio.wbuf = append(fastio.wbuf, ' ')
							}
							fastio.write(rv.Index(j).Interface())
						}
					}
				}
			}
		}
	}
}

func (fastio *fastio) set_precision(x int) {
	fastio.fpc = x
}

func (fastio *fastio) writeln(a ...any) {
	fastio.write(a...)
	fastio.write("\n")
}

func (fastio *fastio) flush() {
	fastio.out.Write(fastio.wbuf)
	fastio.wbuf = fastio.wbuf[:0]
}

var (
	stdio   = new_std_fastio()
	read    = stdio.read
	write   = stdio.write
	writeln = stdio.writeln
	flush   = stdio.flush
)

func solve() {

}

func main() {
	t := 1
	for read(&t); t > 0; t-- {
		solve()
	}
	flush()
}

// https://github.com/baozzii/cp-templates

package main

import (
	"cmp"
	"fmt"
	. "io"
	"math"
	"math/bits"
	"os"
	"reflect"
	"slices"
	"strconv"
	"unsafe"
)

type IO struct {
	in         Reader
	out        Writer
	rbuf, wbuf []byte
	i, n       int
	fpc        int
}

func NewIO(in Reader, out Writer) *IO {
	return &IO{in, out, make([]byte, 4096), make([]byte, 0), 0, 0, -1}
}

func NewStdIO() *IO {
	return NewIO(os.Stdin, os.Stdout)
}

func (io *IO) read_byte() byte {
	if io.i == io.n {
		io.n, _ = io.in.Read(io.rbuf)
		if io.n == 0 {
			return 0
		}
		io.i = 0
	}
	b := io.rbuf[io.i]
	io.i++
	return b
}

func (io *IO) Read(ptrs ...any) {
	var rd func(v reflect.Value)
	rd = func(v reflect.Value) {
		for i := 0; i < v.Len(); i++ {
			elem := v.Index(i)
			if elem.Kind() == reflect.Slice {
				rd(elem)
			} else {
				io.Read(elem.Addr().Interface())
			}
		}
	}
	for _, p := range ptrs {
		switch v := any(p).(type) {
		case *uint:
			{
				var x uint
				b := io.read_byte()
				for ; '0' > b || b > '9'; b = io.read_byte() {
					if b == 0 {
						return
					}
				}
				for ; '0' <= b && b <= '9'; b = io.read_byte() {
					x = x*10 + uint(b&15)
				}
				*v = x
			}
		case *uint8:
			{
				var x uint8
				b := io.read_byte()
				for ; '0' > b || b > '9'; b = io.read_byte() {
					if b == 0 {
						return
					}
				}
				for ; '0' <= b && b <= '9'; b = io.read_byte() {
					x = x*10 + uint8(b&15)
				}
				*v = x
			}
		case *uint16:
			{
				var x uint16
				b := io.read_byte()
				for ; '0' > b || b > '9'; b = io.read_byte() {
					if b == 0 {
						return
					}
				}
				for ; '0' <= b && b <= '9'; b = io.read_byte() {
					x = x*10 + uint16(b&15)
				}
				*v = x
			}
		case *uint32:
			{
				var x uint32
				b := io.read_byte()
				for ; '0' > b || b > '9'; b = io.read_byte() {
					if b == 0 {
						return
					}
				}
				for ; '0' <= b && b <= '9'; b = io.read_byte() {
					x = x*10 + uint32(b&15)
				}
				*v = x
			}
		case *uint64:
			{
				var x uint64
				b := io.read_byte()
				for ; '0' > b || b > '9'; b = io.read_byte() {
					if b == 0 {
						return
					}
				}
				for ; '0' <= b && b <= '9'; b = io.read_byte() {
					x = x*10 + uint64(b&15)
				}
				*v = x
			}
		case *int:
			{
				neg := false
				b := io.read_byte()
				for ; '0' > b || b > '9'; b = io.read_byte() {
					if b == 0 {
						return
					}
					if b == '-' {
						neg = true
					}
				}
				var y uint
				var x int
				for ; '0' <= b && b <= '9'; b = io.read_byte() {
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
				*v = x
			}
		case *int8:
			{
				neg := false
				b := io.read_byte()
				for ; '0' > b || b > '9'; b = io.read_byte() {
					if b == 0 {
						return
					}
					if b == '-' {
						neg = true
					}
				}
				var x int8
				for ; '0' <= b && b <= '9'; b = io.read_byte() {
					x = x*10 + int8(b&15)
				}
				if neg {
					x = -x
				}
				*v = x
			}
		case *int16:
			{
				neg := false
				b := io.read_byte()
				for ; '0' > b || b > '9'; b = io.read_byte() {
					if b == 0 {
						return
					}
					if b == '-' {
						neg = true
					}
				}
				var x int16
				for ; '0' <= b && b <= '9'; b = io.read_byte() {
					x = x*10 + int16(b&15)
				}
				if neg {
					x = -x
				}
				*v = x
			}
		case *int32:
			{
				neg := false
				b := io.read_byte()
				for ; '0' > b || b > '9'; b = io.read_byte() {
					if b == 0 {
						return
					}
					if b == '-' {
						neg = true
					}
				}
				var x int32
				for ; '0' <= b && b <= '9'; b = io.read_byte() {
					x = x*10 + int32(b&15)
				}
				if neg {
					x = -x
				}
				*v = x
			}
		case *int64:
			{
				neg := false
				b := io.read_byte()
				for ; '0' > b || b > '9'; b = io.read_byte() {
					if b == 0 {
						return
					}
					if b == '-' {
						neg = true
					}
				}
				var y uint
				var x int64
				for ; '0' <= b && b <= '9'; b = io.read_byte() {
					y = y*10 + uint(b&15)
				}
				if neg {
					if y == math.MaxInt64+1 {
						x = math.MinInt64
					} else {
						x = -int64(y)
					}
				} else {
					x = int64(y)
				}
				*v = x
			}
		case *float32:
			{
				b := io.read_byte()
				var s []byte
				for ; b == ' ' || b == '\n' || b == '\r' || b == '\t'; b = io.read_byte() {
				}
				for ; !(b == ' ' || b == '\n' || b == '\r' || b == '\t' || b == 0); b = io.read_byte() {
					s = append(s, b)
				}
				w, _ := strconv.ParseFloat(string(s), 32)
				*v = float32(w)
			}
		case *float64:
			{
				b := io.read_byte()
				var s []byte
				for ; b == ' ' || b == '\n' || b == '\r' || b == '\t'; b = io.read_byte() {
				}
				for ; !(b == ' ' || b == '\n' || b == '\r' || b == '\t' || b == 0); b = io.read_byte() {
					s = append(s, b)
				}
				w, _ := strconv.ParseFloat(string(s), 64)
				*v = w
			}
		case *string:
			{
				b := io.read_byte()
				var s []byte
				for ; b == ' ' || b == '\n' || b == '\r'; b = io.read_byte() {
				}
				for ; !(b == ' ' || b == '\n' || b == '\r' || b == 0); b = io.read_byte() {
					s = append(s, b)
				}
				*v = string(s)
			}
		default:
			rv := reflect.ValueOf(p)
			if rv.Kind() == reflect.Ptr && (rv.Elem().Kind() == reflect.Slice || rv.Elem().Kind() == reflect.Array) {
				rd(rv.Elem())
			}
		}
	}
}

func (io *IO) Write(a ...any) {
	uitos := func(v uint64) []byte {
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
	itos := func(v int64) []byte {
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
			io.wbuf = append(io.wbuf, ' ')
		}
		switch v := p.(type) {
		case uint:
			io.wbuf = append(io.wbuf, uitos(uint64(v))...)
		case uint8:
			io.wbuf = append(io.wbuf, uitos(uint64(v))...)
		case uint16:
			io.wbuf = append(io.wbuf, uitos(uint64(v))...)
		case uint32:
			io.wbuf = append(io.wbuf, uitos(uint64(v))...)
		case uint64:
			io.wbuf = append(io.wbuf, uitos(v)...)

		case int:
			io.wbuf = append(io.wbuf, itos(int64(v))...)
		case int8:
			io.wbuf = append(io.wbuf, itos(int64(v))...)
		case int16:
			io.wbuf = append(io.wbuf, itos(int64(v))...)
		case int32:
			io.wbuf = append(io.wbuf, itos(int64(v))...)
		case int64:
			io.wbuf = append(io.wbuf, itos(v)...)

		case float32:
			io.wbuf = append(io.wbuf, []byte(strconv.FormatFloat(float64(v), 'f', io.fpc, 64))...)
		case float64:
			io.wbuf = append(io.wbuf, []byte(strconv.FormatFloat(v, 'f', io.fpc, 64))...)
		case string:
			io.wbuf = append(io.wbuf, v...)
		default:
			rv := reflect.ValueOf(p)
			if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
				if rv.Type().Elem().Kind() == reflect.Slice {
					for j := 0; j < rv.Len(); j++ {
						if j+1 == rv.Len() {
							io.Write(rv.Index(j).Interface())
						} else {
							io.Writeln(rv.Index(j).Interface())
						}
					}
				} else {
					for j := 0; j < rv.Len(); j++ {
						if j != 0 {
							io.wbuf = append(io.wbuf, ' ')
						}
						io.Write(rv.Index(j).Interface())
					}
				}
			}
		}
	}
}

func (io *IO) SetPrecision(x int) {
	io.fpc = x
}

func (io *IO) Writeln(a ...any) {
	io.Write(a...)
	io.Write("\n")
}

func (io *IO) Flush() {
	io.out.Write(io.wbuf)
	io.wbuf = io.wbuf[:0]
}

func (io *IO) Close() {
	io.Flush()
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type RealNumber interface {
	Integer |
		~float32 | ~float64
}

type ComplexNumber interface {
	RealNumber |
		~complex64 | ~complex128
}

type Void struct{}

type NumericLimit[T Integer] struct{}

func Limit[T Integer]() NumericLimit[T] {
	return struct{}{}
}

func (NumericLimit[T]) Max() T {
	var z T
	if (^z) < 0 {
		b := uint(unsafe.Sizeof(z) * 8)
		u := uint(1)<<(b-1) - 1
		return T(u)
	} else {
		return ^z
	}
}

func (NumericLimit[T]) Min() T {
	var z T
	if (^z) < 0 {
		b := uint(unsafe.Sizeof(z) * 8)
		u := uint(1) << (b - 1)
		return T(-int(u))
	}
	return 0
}

func ToString[T any](e T) string {
	return fmt.Sprintf("%v", e)
}

func ToInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

func Chmax[T cmp.Ordered](x *T, y T) {
	*x = max(*x, y)
}

func Chmin[T cmp.Ordered](x *T, y T) {
	*x = min(*x, y)
}

func Sum[T ComplexNumber, E ~[]T](v E) T {
	var s T
	for _, w := range v {
		s += w
	}
	return s
}

func PreSum[T ComplexNumber, E ~[]T](v E) E {
	p := make(E, len(v)+1)
	for i, w := range v {
		p[i+1] = p[i] + w
	}
	return p
}

func Count[T comparable, E ~[]T](v E, e T) int {
	cnt := 0
	for _, w := range v {
		if w == e {
			cnt++
		}
	}
	return cnt
}

func Iota[T Integer, E ~[]T](v E, e T) {
	for i := range v {
		v[i] = e + T(i)
	}
}

func Cond[T any](cond bool, x, y T) T {
	if cond {
		return x
	}
	return y
}

func Ctz[T Integer](x T) int {
	return bits.TrailingZeros(uint(x))
}

func Clz[T Integer](x T) int {
	return bits.LeadingZeros(uint(x))
}

func Popcount[T Integer](x T) int {
	return bits.OnesCount(uint(x))
}

func Lowbit[T Integer](x T) T {
	y := uint(x)
	return T(y & -y)
}

func Highbit[T Integer](x T) T {
	if x == 0 {
		return x
	}
	return T(1) << (63 - Clz(uint(x)))
}

func Abs[T RealNumber](x T) T {
	if x < T(0) {
		return -x
	}
	return x
}

func Gcd[T Integer](x, y T) T {
	if x < 0 || y < 0 {
		return Gcd(Abs(x), Abs(y))
	}
	if y == 0 {
		return x
	}
	return Gcd(y, x%y)
}

func Lcm[T Integer](x, y T) T {
	return x / Gcd(x, y) * y
}

func Pow[S, T Integer](x S, n T, m S) S {
	r := S(1)
	for ; n > 0; n, x = n>>1, x*x%m {
		if n%2 == 1 {
			r = r * x % m
		}
	}
	return r
}

func Exgcd[T Integer](a, b T) (T, T, T) {
	if b == 0 {
		return a, 1, 0
	}
	d, x2, y2 := Exgcd(b, a%b)
	return d, y2, x2 - (a/b)*y2
}

type Vec[T any] []T

func (v *Vec[T]) Size() int {
	return len(*v)
}

func (v *Vec[T]) Empty() bool {
	return len(*v) == 0
}

func (v *Vec[T]) Clear() {
	*v = (*v)[:0]
}

func (v *Vec[T]) PushBack(x T) {
	*v = append(*v, x)
}

func (v *Vec[T]) PopBack() {
	*v = (*v)[:len(*v)-1]
}

func (v *Vec[T]) Resize(x int) {
	if v.Size() > x {
		*v = (*v)[:x]
	} else {
		*v = append(*v, make([]T, x-v.Size())...)
	}
}

func (v *Vec[T]) Back() *T {
	return &((*v)[v.Size()-1])
}

func (v *Vec[T]) Front() *T {
	return &((*v)[0])
}

func (v *Vec[T]) Erase(x, y int) {
	*v = slices.Delete(*v, x, y)
}

func (v *Vec[T]) Insert(x int, w ...T) {
	*v = slices.Insert(*v, x, w...)
}

func (v *Vec[T]) Reverse() {
	n := v.Size()
	for i := 0; i < n/2; i++ {
		(*v)[i], (*v)[n-i-1] = (*v)[n-i-1], (*v)[i]
	}
}

func (v *Vec[T]) Fill(w T) {
	for i := range *v {
		(*v)[i] = w
	}
}

func ToVec1[T any, E ~[]T](v E) Vec[T] {
	return Vec[T](v)
}

func ToSlice1[T any](v Vec[T]) []T {
	return []T(v)
}

func ToVec2[T any, E ~[][]T](v E) Vec[Vec[T]] {
	res := make(Vec[Vec[T]], len(v))
	for i := range v {
		res[i] = ToVec1(v[i])
	}
	return res
}

func ToSlice2[T any](v Vec[Vec[T]]) [][]T {
	res := make([][]T, len(v))
	for i := range v {
		res[i] = ToSlice1(v[i])
	}
	return res
}

func ToVec3[T any, E ~[][][]T](v E) Vec[Vec[Vec[T]]] {
	res := make(Vec[Vec[Vec[T]]], len(v))
	for i := range v {
		res[i] = ToVec2(v[i])
	}
	return res
}

func ToSlice3[T any](v Vec[Vec[Vec[T]]]) [][][]T {
	res := make([][][]T, len(v))
	for i := range v {
		res[i] = ToSlice2(v[i])
	}
	return res
}

func Vec1[T any](n ...int) Vec[T] {
	if len(n) == 0 {
		return make(Vec[T], 0)
	}
	return make(Vec[T], n[0])
}

func Vec2[T any](n ...int) Vec[Vec[T]] {
	if len(n) == 0 {
		return make(Vec[Vec[T]], 0)
	}
	v := make(Vec[Vec[T]], n[0])
	for i := range v {
		v[i] = Vec1[T](n[1:]...)
	}
	return v
}

func Vec3[T any](n ...int) Vec[Vec[Vec[T]]] {
	if len(n) == 0 {
		return make(Vec[Vec[Vec[T]]], 0)
	}
	v := make(Vec[Vec[Vec[T]]], n[0])
	for i := range v {
		v[i] = Vec2[T](n[1:]...)
	}
	return v
}

func (v *Vec[T]) Copy() Vec[T] {
	return slices.Clone(*v)
}

func (v *Vec[T]) Slice(i, j int) Vec[T] {
	w := (*v)[i:j]
	return w.Copy()
}

type Uset[T comparable] map[T]struct{}

func (s *Uset[T]) Insert(x T) {
	(*s)[x] = struct{}{}
}

func (s *Uset[T]) Erase(x T) {
	delete(*s, x)
}

func (s *Uset[T]) Contains(x T) bool {
	_, ok := (*s)[x]
	return ok
}

func (s *Uset[T]) Size() int {
	return len(*s)
}

func (s *Uset[T]) Empty() bool {
	return s.Size() == 0
}

func (s *Uset[T]) Clear() {
	clear(*s)
}

func (s *Uset[T]) Keys() Vec[T] {
	b := make(Vec[T], 0, s.Size())
	for v := range *s {
		b.PushBack(v)
	}
	return b
}

func (s *Uset[T]) FromVec(v Vec[T]) {
	for _, w := range v {
		s.Insert(w)
	}
}

type Umap[T comparable, K any] map[T]K

func (s *Umap[T, K]) Erase(x T) {
	delete(*s, x)
}

func (s *Umap[T, K]) Contains(x T) bool {
	_, ok := (*s)[x]
	return ok
}

func (s *Umap[T, K]) Size() int {
	return len(*s)
}

func (s *Umap[T, K]) Empty() bool {
	return s.Size() == 0
}

func (s *Umap[T, K]) Clear() {
	clear(*s)
}

func (s *Umap[T, K]) Keys() Vec[T] {
	b := make(Vec[T], 0, s.Size())
	for v := range *s {
		b.PushBack(v)
	}
	return b
}

func (s *Umap[T, K]) Values() Vec[K] {
	b := make(Vec[K], 0, s.Size())
	for _, v := range *s {
		b.PushBack(v)
	}
	return b
}

func Counter[T comparable, E ~[]T](v E) Umap[T, int] {
	cnt := make(Umap[T, int])
	for _, w := range v {
		cnt[w]++
	}
	return cnt
}

var io = NewStdIO()

func solve() {

}

func main() {
	t := 1
	for io.Read(&t); t > 0; t-- {
		solve()
	}
	io.Close()
}

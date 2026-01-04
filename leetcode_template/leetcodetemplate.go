// https://github.com/baozzii/cp-templates

package main

import (
	"bufio"
	"cmp"
	"encoding/json"
	"fmt"
	"math/bits"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"unsafe"
)

type LeetcodeIO struct {
	paramtypes []reflect.Type
	paramcount int
	fntype     reflect.Type
	fn         reflect.Value
	in         *bufio.Reader
	out        *bufio.Writer
}

func RegisterFunc(f any) *LeetcodeIO {
	fn := reflect.ValueOf(f)
	fntype := reflect.TypeOf(f)
	if fntype == nil || fntype.Kind() != reflect.Func {
		panic(-1)
	}
	pcount := fntype.NumIn()
	paramtypes := make([]reflect.Type, pcount)
	for i := 0; i < pcount; i++ {
		paramtypes[i] = fntype.In(i)
	}
	return &LeetcodeIO{paramtypes, pcount, fntype, fn, bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)}
}

func (io *LeetcodeIO) parse(s string, r reflect.Type) reflect.Value {
	if r.Kind() == reflect.String {
		var v string
		if err := json.Unmarshal([]byte(s), &v); err != nil {
			panic(err)
		}
		return reflect.ValueOf(v).Convert(r)
	}
	switch r.Kind() {
	case reflect.Bool:
		v, err := strconv.ParseBool(s)
		if err != nil {
			panic(err)
		}
		x := reflect.New(r).Elem()
		x.SetBool(v)
		return x
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		x := reflect.New(r).Elem()
		x.SetInt(v)
		return x
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		v, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			panic(err)
		}
		x := reflect.New(r).Elem()
		x.SetUint(v)
		return x
	case reflect.Float32, reflect.Float64:
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic(err)
		}
		x := reflect.New(r).Elem()
		x.SetFloat(v)
		return x
	}
	ptr := reflect.New(r)
	if err := json.Unmarshal([]byte(s), ptr.Interface()); err != nil {
		panic(err)
	}
	return ptr.Elem()
}

func (io *LeetcodeIO) format(r reflect.Value) string {
	b, _ := json.Marshal(r.Interface())
	return string(b)
}

func (io *LeetcodeIO) Run() {
	defer io.out.Flush()
	for {
		args := make([]reflect.Value, io.paramcount)
		for i := 0; i < io.paramcount; i++ {
			s, err := io.in.ReadString('\n')
			if err != nil && len(s) == 0 {
				return
			}
			args[i] = io.parse(strings.TrimSpace(s), io.paramtypes[i])
		}
		res := io.execute(args)
		fmt.Fprintln(io.out, io.format(res[0]))
	}
}

func (io *LeetcodeIO) execute(args []reflect.Value) []reflect.Value {
	return io.fn.Call(args)
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

func main() {
	RegisterFunc(nil).Run()
}

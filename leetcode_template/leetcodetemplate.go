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

type leetcodeio struct {
	paramtypes []reflect.Type
	paramcount int
	fntype     reflect.Type
	fn         reflect.Value
	in         *bufio.Reader
	out        *bufio.Writer
}

func register(f any) *leetcodeio {
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
	return &leetcodeio{paramtypes, pcount, fntype, fn, bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)}
}

func (io *leetcodeio) __parse(s string, r reflect.Type) reflect.Value {
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

func (io *leetcodeio) __format(r reflect.Value) string {
	b, _ := json.Marshal(r.Interface())
	return string(b)
}

func (io *leetcodeio) run() {
	defer io.out.Flush()
	for {
		args := make([]reflect.Value, io.paramcount)
		for i := 0; i < io.paramcount; i++ {
			s, err := io.in.ReadString('\n')
			if err != nil && len(s) == 0 {
				return
			}
			args[i] = io.__parse(strings.TrimSpace(s), io.paramtypes[i])
		}
		res := io.__execute(args)
		fmt.Fprintln(io.out, io.__format(res[0]))
	}
}

func (io *leetcodeio) __execute(args []reflect.Value) []reflect.Value {
	return io.fn.Call(args)
}

func solve() {

}

func main() {
	register(nil).run()
}

package templates

import (
	"io"
	"math"
	"os"
	"reflect"
	"strconv"
)

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

func (fastio *fastio) close() {
	fastio.flush()
}

var stdio = new_std_fastio()
var read = stdio.read
var write = stdio.write
var writeln = stdio.writeln
var flush = stdio.flush

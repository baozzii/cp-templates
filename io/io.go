package io

import (
	. "io"
	"os"
	"reflect"
	"strconv"
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
		io.n, _ = os.Stdin.Read(io.rbuf)
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
				var x int
				for ; '0' <= b && b <= '9'; b = io.read_byte() {
					x = x*10 + int(b&15)
				}
				if neg {
					x = -x
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
				var x int64
				for ; '0' <= b && b <= '9'; b = io.read_byte() {
					x = x*10 + int64(b&15)
				}
				if neg {
					x = -x
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
			if rv.Kind() == reflect.Slice {
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

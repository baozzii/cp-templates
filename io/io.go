package io

import (
	. "io"
	"os"
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
	for i, p := range a {
		var s []byte
		switch v := p.(type) {
		case uint, uint8, uint16, uint32, uint64:
			{
				if v == 0 {
					s = append(s, '0')
				} else {
					for v.(uint) > 0 {
						s = append(s, '0'+byte(v.(uint)%10))
						v = v.(uint) / 10
					}
					for j := 0; j < len(s)/2; j++ {
						s[j], s[len(s)-j-1] = s[len(s)-j-1], s[j]
					}
				}
			}
		case int, int8, int16, int32, int64:
			{
				if v == 0 {
					s = append(s, '0')
				} else {
					neg := false
					if v.(int) < 0 {
						neg = true
						v = -(v.(int))
					}
					for v.(int) > 0 {
						s = append(s, '0'+byte(v.(int)%10))
						v = v.(int) / 10
					}
					if neg {
						s = append(s, '-')
					}
					for j := 0; j < len(s)/2; j++ {
						s[j], s[len(s)-j-1] = s[len(s)-j-1], s[j]
					}
				}
			}
		case float32, float64:
			{
				s = []byte(strconv.FormatFloat(v.(float64), 'f', io.fpc, 64))
			}
		case string:
			{
				s = []byte(v)
			}
		}
		io.wbuf = append(io.wbuf, s...)
		if i != len(a)-1 {
			io.wbuf = append(io.wbuf, ' ')
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

func (io *IO) Close() {
	io.out.Write(io.wbuf)
}

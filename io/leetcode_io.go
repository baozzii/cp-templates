package io

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
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

/*
Usage:
Program will read from stdin, one line per parameter.
Supports multiple test cases.
Program will write to stdout, one line per test case.

func main() {
	RegisterFunc(_FuncName_).Run()
}
*/

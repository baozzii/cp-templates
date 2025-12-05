package template

import . "codeforces-go/io"

var io = NewStdIO()

func solve() {

}

func main() {
	var t int
	for io.Read(&t); t > 0; t-- {
		solve()
	}
	io.Close()
}

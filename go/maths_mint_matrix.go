package templates

type mint_matrix vector[vector[mint]]

func new_identity_mint_matrix(n int) mint_matrix {
	a := vec2[mint](n, n)
	for i := range n {
		a[i][i] = 1
	}
	return mint_matrix(a)
}

func new_zero_mint_matrix(m, n int) mint_matrix {
	return mint_matrix(vec2[mint](m, n))
}

func (a mint_matrix) row() int {
	return len(a)
}

func (a mint_matrix) col() int {
	return len(a[0])
}

func (a mint_matrix) matrix_mul(b mint_matrix) mint_matrix {
	m := a.row()
	n := a.col()
	l := b.col()
	c := new_zero_mint_matrix(m, l)
	for i := range m {
		for j := range n {
			x := a[i][j]
			for k := range l {
				c[i][k] = c[i][k].add(x.mul(b[j][k]))
			}
		}
	}
	return c
}

func (a mint_matrix) add(b mint_matrix) mint_matrix {
	c := new_zero_mint_matrix(a.row(), a.col())
	for i := range a.row() {
		for j := range a.col() {
			c[i][j] = a[i][j].add(b[i][j])
		}
	}
	return c
}

func (a mint_matrix) sub(b mint_matrix) mint_matrix {
	c := new_zero_mint_matrix(a.row(), a.col())
	for i := range a.row() {
		for j := range a.col() {
			c[i][j] = a[i][j].sub(b[i][j])
		}
	}
	return c
}

func (a mint_matrix) scalar_mul(x mint) mint_matrix {
	c := new_zero_mint_matrix(a.row(), a.col())
	for i := range a.row() {
		for j := range a.col() {
			c[i][j] = a[i][j].mul(x)
		}
	}
	return c
}

package templates

type matrix[T complex] vector[vector[T]]

func new_identity_matrix[T complex](n int) matrix[T] {
	a := vec2[T](n, n)
	for i := range n {
		a[i][i] = 1
	}
	return matrix[T](a)
}

func new_zero_matrix[T complex](m, n int) matrix[T] {
	return matrix[T](vec2[T](m, n))
}

func (a matrix[T]) row() int {
	return len(a)
}

func (a matrix[T]) col() int {
	return len(a[0])
}

func (a matrix[T]) matrix_mul(b matrix[T]) matrix[T] {
	m := a.row()
	n := a.col()
	l := b.col()
	c := new_zero_matrix[T](m, l)
	for i := range m {
		for j := range n {
			x := a[i][j]
			for k := range l {
				c[i][k] += x * b[j][k]
			}
		}
	}
	return c
}

func (a matrix[T]) add(b matrix[T]) matrix[T] {
	c := new_zero_matrix[T](a.row(), a.col())
	for i := range a.row() {
		for j := range a.col() {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return a
}

func (a matrix[T]) sub(b matrix[T]) matrix[T] {
	c := new_zero_matrix[T](a.row(), a.col())
	for i := range a.row() {
		for j := range a.col() {
			c[i][j] = a[i][j] - b[i][j]
		}
	}
	return a
}

func (a matrix[T]) scalar_mul(x T) matrix[T] {
	c := new_zero_matrix[T](a.row(), a.col())
	for i := range a.row() {
		for j := range a.col() {
			c[i][j] = a[i][j] * x
		}
	}
	return a
}

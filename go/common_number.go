package templates

type unsigned_integer interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type signed_integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type integer interface {
	signed_integer | unsigned_integer
}

type real interface {
	integer |
		~float32 | ~float64
}

type complex interface {
	real |
		~complex64 | ~complex128
}

package templates

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

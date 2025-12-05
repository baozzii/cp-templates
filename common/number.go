package common

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type RealNumber interface {
	Integer |
		~float32 | ~float64
}

type ComplexNumber interface {
	RealNumber |
		~complex64 | ~complex128
}

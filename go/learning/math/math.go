package math

// signed is a constraint that permits any signed integer type.
// int8, int16 are not included because they are rarely used.
type signed interface{ ~int | ~int32 | ~int64 }

// unsigned is a constraint that permits any unsigned integer type.
// uint8, uint16, uintptr are not included because they are rarely used.
type unsigned interface{ ~uint | ~uint32 | ~uint64 }

// integer is a constraint that permits any integer type.
type integer interface{ signed | unsigned }

// float is a constraint that permits any floating-point type.
type float interface{ ~float32 | ~float64 }

// actual is a constraint that permits any complex numeric type.
type actual interface{ integer | float }

// imaginary is a constraint that permits any complex numeric type.
type imaginary interface{ ~complex64 | ~complex128 }

// ordered is a constraint that permits any ordered type: any type
type ordered interface{ integer | float | ~string }

// addable is a constraint that permits any ordered type: any type
type addable interface {
	integer | float | imaginary | ~string
}

// Pow returns x**n, the base-x exponential of n.
func Pow[T actual](x T, n int) T {
	y := T(1)
	for n > 0 {
		if n&1 == 1 {
			y *= x
		}
		x *= x
		n >>= 1
	}
	return y
}

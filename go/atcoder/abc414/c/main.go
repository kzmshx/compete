package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func IsPalindrome[T Int](a []T) bool {
	n := len(a)
	for i := 0; i < n/2; i++ {
		if a[i] != a[n-1-i] {
			return false
		}
	}
	return true
}

func IsPalindromeInt[T Int](n T, base T) bool {
	return IsPalindrome(Digits(n, base))
}

func Solve(r *Reader, w *Writer) {
	a, n := r.Int(), r.Int()
	sum := 0

	createPalindrome := func(half int, nd int, base int) int {
		halfDigits := Digits(half, base)

		ndNeed := (nd + 1) / 2
		for len(halfDigits) < ndNeed {
			halfDigits = append(halfDigits, 0)
		}

		var palindrome []int
		for i := len(halfDigits) - 1; i >= 0; i-- {
			palindrome = append(palindrome, halfDigits[i])
		}
		start := 0
		if nd%2 == 1 {
			start = 1
		}
		for i := start; i < len(halfDigits); i++ {
			palindrome = append(palindrome, halfDigits[i])
		}
		result := 0
		m := 1
		for i := len(palindrome) - 1; i >= 0; i-- {
			result += palindrome[i] * m
			m *= base
		}
		return result
	}

	for i := 1; i < a && i <= n; i++ {
		sum += i
	}

	for nd := 2; Pow(a, nd-1) <= n; nd++ {
		ndhalf := (nd + 1) / 2
		start, end := Pow(a, ndhalf-1), Pow(a, ndhalf)
		for i := start; i < end; i++ {
			num := createPalindrome(i, nd, a)
			if num > n {
				break
			}
			if IsPalindromeInt(num, 10) {
				sum += num
			}
		}
	}

	w.Println(sum)
}

func main() {
	r, w := NewReader(os.Stdin, 1*1024*1024), NewWriter(os.Stdout)
	defer w.Flush()
	Solve(r, w)
}

type Sgn interface{ ~int | ~int32 | ~int64 }
type Uns interface{ ~uint | ~uint32 | ~uint64 }
type Int interface{ Sgn | Uns }
type Float interface{ ~float32 | ~float64 }
type Num interface{ Int | Float }
type Ord interface{ Int | Float | ~string }

func Unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
func Atoi(s string) int { return Unwrap(strconv.Atoi(s)) }

type Reader struct{ sc *bufio.Scanner }

func NewReader(r io.Reader, size int) *Reader {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, size), size)
	sc.Split(bufio.ScanWords)
	return &Reader{sc}
}
func (r *Reader) Int() int    { return Atoi(r.Str()) }
func (r *Reader) Str() string { r.sc.Scan(); return r.sc.Text() }

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer { return &Writer{bufio.NewWriter(w)} }
func (w *Writer) Println(a ...any)  { fmt.Fprintln(w.bf, a...) }
func (w *Writer) Flush()            { w.bf.Flush() }

func Digits[T Int](n T, base T) (r []T) {
	for n > 0 {
		r = append(r, n%base)
		n /= base
	}
	return r
}

func Pow[T Num](x T, n int) T {
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

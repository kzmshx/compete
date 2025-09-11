
| key  | value                                               |
| ---- | --------------------------------------------------- |
| url  | <https://atcoder.jp/contests/abc414/tasks/abc414_c> |
| tags | #palindrome                                         |

## 考察

- この問題の解答は $2^{63}$ 未満なので i64 整数で表現可能
- $N$ を $A$ 進数に変換する計算量は $\log_A{N}$
- Brute-force だと $O(10^{12} \times \log_A{N})$
- 10進数で回分になる数を効率よく求める必要がある
- 回分の求め方
	- 10進数 `n` の桁数 `i | 1 <= i <= digits(n)`
	- 列挙: 1 ~ (digits(n) + 1) / 2 までの桁の数を 0 ~ 9 で試す、その数が N より大きいかどうかチェック

## 実装

```go
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

func IsPalindromeString(s string) bool {
	return IsPalindrome([]rune(s))
}

func Solve(r *Reader, w *Writer) {
	a, n := r.Int(), r.Int()
	c := 0
	
	createPalindrome := func (half int, nd int, base int) int {
		halfDigits := Digits(half, base)
		p, m := 0, 1
		for i := 0; i < len(halfDigits); i++ {
			p += halfDigits[i] * m
			m *= base
		}
		mirrorStart := len(halfDigits)
		if nd % 2 == 1 {
			mirrosStart--
		}
		for i := mirrorStart - 1; i >= 0; i-- {
			p += halfDigits[i] * m
			m *= base
		}
		return p
	}
	
	for i := 1; i < a && i <= n; i ++ {
		c++
	}
	
	for nd := 2; Pow(a, nd-1) <= n; nd++ {
		start, end := Pow(a, (nd-1)/2), Pow(a, nd/2)
		for i := start; i < end; i++ {
			num := createPalindrome(i, nd, a)
			if num > n {
				break
			}
			if isPalindromeInt(num, 10) {
				c++
			}
		}
	}
	
	w.Println(c)
}
```
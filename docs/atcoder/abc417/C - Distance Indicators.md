
| key  | value                                             |
| ---- | ------------------------------------------------- |
| url  | https://atcoder.jp/contests/abc417/tasks/abc417_c |
| tags | #combination                                      |
| lose | #lose/2025/09/05                                  |

## 考察

- 計算量
	- $1 \le N \le 2 \times 10^5$ -> $O(N^2)$ は NG、可能なら $O(N)$ で解きたい
	- 数字列なので先頭から順に処理して解答にたどり着けるなら $O(N)$ になる
- DP か？
	- DP であるなら配列の値に答えが積まれていくようなロジックになる
	- $A_1$ を $C$ とするとき、その後の要素において $j - i = A_i + A_j$ が満たされるとは $j - 1 = $C + A_j$ が満たされるということ
	- もう一つ変数を確定すると判定結果が決まる
	- $1 \le k \le N$, $2 \le j \le N$
	- `dp[k][j]`: k 個目までの数列を使って $j - i = A_i + A_j$ を満たす組の個数
	- んん、違った
- ペア数を数える
	- 条件式変形
		- `j - i = A[i] + A[j]`
		- `j - A[j] = i + A[i]`
		- こうすると添字が `i` の側と `j` の側に分けられる
		- そうなると `i` について `i + A[i]` の位置を記録、`j` について `j - A[j]` の位置を記録、両者が一致する数をカウントする、というロジックにできる

## 実装

### Brute-force: $O(N^2)$

```go
func Solve(r *Reader, w *Writer) {
	n := r.Int()
	a := r.Ints(n)
	
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if j - i == a[i] + a[j] {
				count++
			}
		}
	}
	
	w.Println(count)
}
```

### 式変形

```go
	c1, c2 := make(map[int]int), make(map[int]int)
	for i, n := 0, r.Int(); i < n; i++ {
		a := r.Int()
		if v := i + 1 - a; v >= 0 {
			c1[v]++
		}
		if v := i + 1 + a; v < n {
			c2[v]++
		}
	}

	ans := 0
	for k, c := range c1 {
		ans += c * c2[k]
	}

	w.Println(ans)
```

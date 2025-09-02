
| key  | value                                             |
| ---- | ------------------------------------------------- |
| url  | https://atcoder.jp/contests/abc418/tasks/abc418_c |
| tags | #binary_search                                             |

## 考察

- 勝利する（= $x$ 個の中から $b$ 個同じフレーバーを選ぶ）には、ディーラーがどのように $x$ を選んでもそこに同じフレーバーが $b$ 個以上含まれる必要がある。
- $\max(A_1, A_2, ..., A_N) < b$ の場合、そもそも同じフレーバーが $b$ ないので、勝利は不可能
- $b \le A_i$ を満たすフレーバーであれば勝利は可能
- 勝利するためには $b$ 個以下のフレーバーがすべて $x$ 個の中に含まれる必要がある
- つまり、
- 勝利可能条件は $max(A_1, A_2, ..., A_N) >= b$
- 累積和と二分探索を使ってみる

## 実装

```go
func Solve(r *Reader, w *Writer) {
	n, q := r.Int(), r.Int()
	a := r.Ints(n)
	sort.Ints(a)

	asum := make([]int, n+1)
	asum[0] = 0
	for i, v := range a {
		asum[i+1] = asum[i] + v
	}

	for i := 0; i < q; i++ {
		b := r.Int()
		if j := BinarySearch(0, len(a), func(i int) bool { return a[i] >= b }); j < len(a) {
			w.Println(asum[j] + (b-1)*(len(a)-j) + 1)
		} else {
			w.Println(-1)
		}
	}
}
```
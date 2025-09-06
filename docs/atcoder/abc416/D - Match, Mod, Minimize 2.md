
| key  | value                                             |
| ---- | ------------------------------------------------- |
| url  | https://atcoder.jp/contests/abc416/tasks/abc416_d |
| tags | #greedy #two_pointer                              |

## 考察

- 貪欲法（Greedy Algorithm）
	- 貪欲法は、局所最適解が全体最適を損なわないことを前提に、順番に局所最適解を選び続けることによって、全探索よりも効率的に解を得る方法
	- = `k` までの選択が最適ならば、`k+1` 番目の貪欲選択も全体最適を保持する
	- = `k` までの局所最適解を抜いても、`k+1` 番目の最適解は残る
	- それが使える状況の証明が重要
- モジュラー演算の性質
	- $(a + b) \mod{M}$
		- $a \ge M - b$ の場合: $(a + b) \mod{M} = a + b - M$
		- $a \lt M - b$ の場合: $(a + b) \mod{M} = a + b$
- これを利用すると
	- `target = M - B[i]` として
		- (1) `A[j] >= target` の場合 `(A[j] + B[i]) mod M = A[j] + B[i] - M`
			- -> `A[j]` が小さいほど余りが小さくなる
			- -> `target` 以上の未使用最小値を選ぶのが最適
		- (2) `A[j] < target` の場合 `(A[j] + B[i]) mod M = A[j] + B[i]`
			- -> `A[j]` が小さいほど余りが小さくなる
			- -> しかし大きい値は将来の `B[k]` で (1) のパターンになる可能性
			- -> 未使用の最小値を選び、大きい値は将来のために温存

## 実装

### Greedy: $O(N^2)$

- この方法だと配列部分で $O(N^2)$ になってしまう

```go
func Solve(r *Reader, w *Writer) {
	t := r.Int()
	for i := 0; i < t; i++ {
		n, m := r.Int(), r.Int()
		a := r.Ints(n)
		sort.Ints(a)
		
		ans := 0
		for i := 0; i < n; i++ {
			b := r.Int()
			if j := LowerBound(a, m-b); j == n {
				ans, a = ans + a[0], a[1:]
			} else {
				ans, a = ans + a[j], append(a[0:j], a[j+1:])
			}
		}
		w.Println(ans)
	}
}
```

### Greedy + Two-Pointer

- 計算を見直す
	- $A_i + B_i \ge M$ を満たす $i$ の個数を $C$ とすると $\sum_{i=1}^{N} ((A_i + B_i) \mod M) = \sum_{i=1}^{N} (A_i + B_i) - CM$
	- -> $C$ を最大化するように A を並び替えれば良い

```go
func Solve(r *Reader, w *Writer) {
	for i, t := 0, r.Int(); i < t; i++ {
		n, m := r.Int(), r.Int()
		a, asum, b, bsum := make([]int, n), 0, make([]int, n), 0
		for i := 0; i < n; i++ {
			a[i] = r.Int()
			asum += a[i]
		}
		for i := 0; i < n; i++ {
			b[i] = r.Int()
			bsum += b[i]
		}
		sort.Sort(sort.Reverse(sort.IntSlice(a)))
		sort.Sort(sort.IntSlice(b))

		c, j := 0, 0
		for i := 0; i < n; i++ {
			for j < n && a[i]+b[j] < m {
				j++
			}
			if j < n {
				c++
				j++
			}
		}

		w.Println(asum + bsum - m*c)
	}
}
```
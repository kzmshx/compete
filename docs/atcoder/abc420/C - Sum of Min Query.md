
| key  | value                                             |
| ---- | ------------------------------------------------- |
| url  | https://atcoder.jp/contests/abc420/tasks/abc420_c |
| tags |                                                   |
| lose |                                                   |

## 考察

- 計算量
	- $Q$ を $O(N)$ で捌くのは OK
	- $QN$ はそのままだと $O(10^{10})$ なので NG
	- $Q$ を処理すると同時に $N$ 分を最適化できると効率よい
- $A$ と $B$ の Min を最初から取っておいて、そこに $Q$ を適用するだけでいいのでは

## 実装

```
n, q := r.Int(), r.Int()
a, b, sum := r.Ints(n), make([]int, n), 0
for i := 0; i < n; i++ {
	b[i] = r.Int()
	sum += Min(a[i], b[i])
}
for i := 0; i < q; i++ {
	if c, x, v := r.String(), r.Int(), r.Int(); c == "A" {
		sum += Min(v, b[i]) - Min(a[i], b[i])
	} else if c == "B" {
		sum += Min(a[i], v) - Min(a[i], b[i])
	} else {
		panic("unexpected query")
	}
	w.Println(sum)
}
```
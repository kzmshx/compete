
| key  | value                                             |
| ---- | ------------------------------------------------- |
| url  | https://atcoder.jp/contests/abc419/tasks/abc419_d |
| tags | #diff_array #bit                                      |

## 考察

- $S$ , $T$ は同じ長さ $N$
- $1 \le N \le 5 \times 10^{5}$ , $1 \le M \le 2 \times 10^{5}$
- 素直に考えるなら $1$ から $M$ まで、 `S[Li,Ri]` と `T[Li, Ri]` を swap していく → これは TLE
- 実際の文字列スワップ操作を素直にやると $M$ 分が残るので TLE
- $O(M)$ を $O(1)$ に近づければよい、文字列スワップ操作をビット演算に置き換えると $O(1)$ で済むのでは
- $S$ であるインデックスを $0$ とし、 $T$ であるインデックスを $1$ として表現すれば、ビット演算が使えそう、ただビットは 64 個なので $5 \times 10^5$ に対応できない
- 範囲の反転の情報を効率的に更新して保持できるデータ構造が必要

## 実装

### Brute-force

```
var n, m int
var s, t string
operations [][2]int

for _, op = range operations {
	new_s = s[0:op[0]-1] + t[op[0]-1:op[1]-1] + s[op[1]-1:]
	new_t = t[0:op[0]-1] + s[op[0]-1:op[1]-1] + t[op[1]-1:]
	s = new_s
	t = new_t
}

print(new_s)
```

- これは TLE になる
- これは $O(NM) = O(5 \times 10^5 \times 2 \times 10^5) = O(10^{11})$ になるから

### ビット演算

- $1 \le N \le 64$ なら使えたはず、今回は N が長いので無理

### 差分配列（DiffArray）

```go
type DiffArray[T Addable] struct {
	diff []T
}

func NewDiffArray[T Addable](n int) *DiffArray[T] {
	return &DiffArray[T]{ diff: make([]T, n+1) }
}

func (da *DiffArray[T]) Add(l, r int, val T) {
	da.diff[l] += val
	da.diff[r] -= val
}

func (da *DiffArray[T]) Increment(l, r int) {
	da.Add(l, r, 1)
}

func (da *DiffArray[T]) Build() []T {
	n := len(d.diff) - 1
	result := make([]T, n)
	var sum T
	for i := 0; i < n; i++ {
		sum += d.diff[i]
		result[i] = sum
	}
	return result
}
```

```
var n, m int
var s, t string
var ops [][2]int

diff = diff_array(size: n)
for _, op = range ops {
	diff.increment(l: op[l], r: op[r]+1)
}

diff_built = diff.build()
var new_s string
for i, v = range diff_built {
	if v % 2 == 0 {
		new_s[i] = s[i]
	} else {
		new_s[i] = t[i]
	}
}
print(new_s)
```
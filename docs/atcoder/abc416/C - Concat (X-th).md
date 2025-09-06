
| key  | value                                             |
| ---- | ------------------------------------------------- |
| url  | https://atcoder.jp/contests/abc416/tasks/abc416_c |
| tags | #brute_force                                      |
| lose |                                                   |

## 考察

- $1 \le N \le 10, 1 \le K \le 5$
- $N$ 個の文字列 $S_1, ..., S_N$ の中から $K$ 個選んで連結してできる文字列 $f(A_1, ..., A_K)$ を辞書順に並べる、その中で小さい方から $X$ 番目の文字列を求める
- 同じ $i$  番目の文字列 $S_i$ を何度使ってもいい設定なのでパターン数は $N^K$ 個ある
- $X$ は最大 $10^5$ になる

## 実装

### Brute-force: $O(N + N^K + N \log{N} + N) = (N(N^{K-1} + \log{N} + 2))$

- $N^K$ が $10^{5}$ オーダーなので全探索でも間に合う

```go
func enumerate(baseStrs []string, strs []string, i int) []string {
	if i == 0 {
		return strs
	}
	newStrs := make([]string, len(strs)*len(baseStrs))
	for i, s := range strs {
		for j, b := range baseStrs {
			newRet[i*len(s)+j] = s + b
		}
	}
	return newRet
}

func Solve(r *Reader, w *Writer) {
	n, k, x := r.Int(), r.Int(), r.Int()
	// O(N)
	s := r.Strings(n)
	// O(N^K)
	strs := enumerate(s, s, k-1)
	// O(N*log(N))
	sort.Strings(strs)
	w.Println(strs[x-1])
}
```
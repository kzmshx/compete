
| key  | value                                             |
| ---- | ------------------------------------------------- |
| url  | https://atcoder.jp/contests/abc421/tasks/abc421_c |
| tags |                                                   |

## 考察

- 「すべての A が奇数インデックスになるコスト」と「すべての A が偶数インデックスになるコスト」、どっちが小さいか問題

## 実装

```go
func Solve(r *Reader, w *Writer) {
	_, s := r.Int(), r.String()
	t := 0
	e, o := 0, 0
	for i, c := range s {
		if c == 'B' {
			continue
		}
		e, o = e+Abs(i-(t+1)), o+Abs(i-t)
		t += 2
	}
	w.Println(Min(e, o))
}
```
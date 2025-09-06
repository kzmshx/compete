
| key  | value                                             |
| ---- | ------------------------------------------------- |
| url  | https://atcoder.jp/contests/abc418/tasks/abc418_b |
| tags | #brute_force                                      |
| lose |                                                   |

## 考察

- 全探索 $O(N^2)$
- $1 \le |S| \le 100$ なので無問題
- 素直に解けばすぐにできるし間に合うけど、自分は逐一最適で綺麗な解法を目指しがちな気がする、それだとコンテストでは他の問題に時間を割きづらいので注意

## 実装

```go
func Solve(r *Reader, w *Writer) {
	s := r.String()
	n := len(s)
	maxRate := 0.0
	
	for i := 0; i < n; i++ {
		for j := i + 3; j <= n; j++ {
			if sub := s[i:j]; sub[0] == 't' && sub[len(sub)-1] == 't' {
				tCount := 0
				for _, c := range sub {
					if c == 't' {
						tCount++
					}
				}
				rate := float64(tCount-2) / float64(len(sub)-2)
				ChooseMax(&maxRate, rate)
			}
		}
	}
	
	w.PrintlnFloat64(maxRate)
}
```
## URL

[C - King's Summit](https://atcoder.jp/contests/abc419/tasks/abc419_c)

## タグ

#二分探索 #チェビシェフ距離 #座標範囲の交差 #LowerBound

## 考察

- $10^9 \times 10^9$ 列とあるので全探索は不可
- 8近傍マスに移動できるということは、初期座標を $(x, y)$ として時刻 $t$ で到達できる範囲は $[x-t, x+t], [y-t, y+t]$ の正方形になる
- N 個の点からチェビシェフ距離 a の座標で N 個の座標が重なるとする、この a を最小化する
- 与えられた座標範囲（$10^9 \times 10^9$ マス）において、各座標からその他の座標へのチェビシェフ距離は最小 $1$、最大 $10^9 - 1$
- N 個の各点を中心とする距離 t の範囲がすべて重なるかどうかの計算量は O(N)、これは各点がもつ範囲を Intersect することを N 回繰り返せば良いから、
- N の制約は、$1 < N < 2 \times 10^5$
- $1$ から $10^9 - 1$  までを全て試すのは $10^{9}$ オーダーなので TLE、二分探索なら $log(10^{9})$ で済む
- 全体の計算量は $2 \times 10^{5} \times log(10^{9})$
- 二分探索で行ける理由は、各点の移動可能範囲は時刻の増加によって単調増加し、いずれかの時刻以降は常に N 個の座標が重なる点が存在するから
- 二分探索、もっといえば LowerBound

## 擬似コード

```
fn binary_search(l int, r int, f fn (int) bool) int {
	for l < r {
		m = (l + r) / 2
		if f(m) {
			r = m
		} else {
			l = m + 1	
		}
	}
	return l
}

fn intersect_1d(a [2]int, b [2]int) [[2]int, bool] {
	l = max(a[0], b[0])
	r = min(a[1], b[1])
	return [[l, r], min <= max]
}

fn intersect_2d(a [2][2]int, b [2][2]int) [[2][2]int, bool] {
	[row, ok_row] = intersect_1d(a[0], b[0])
	[col, ok_col] = intersect_1d(a[1], b[1])
	return [[row, col], ok_row && ok_col]
}

fn main() {
	var n int
	var coords [][2]int
	binary_search(1, 10**9, fn (i int) bool {
		[cur, ok] = [[[1, 10**9], [1, 10**9]], true]
		for _, c = range coords {
			[cur, ok] = intersect_2d(cur, [
				[max(1, c[0]-i), min(10**9, c[0]+i)],
				[max(1, c[1]-i), min(10**9, c[1]+i)],
			])
			if !ok {
				return false
			}
		}
		return true
	})
}
```

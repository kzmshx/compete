
| key  | value                                             |
| ---- | ------------------------------------------------- |
| url  | https://atcoder.jp/contests/abc419/tasks/abc419_e |
| tags | #動的計画法 #剰余                                        |

## 考察

- 長さ $L$ の連続部分配列の和をすべて $M$ で割り切れるようにする
	- $S_i = A_0 + A_1 + ... + A_{i-1}$ を累積和とするとき
	- 長さ $L$  の部分配列 $A_i, A_{i+1}, ..., A_{i+L-1}$ の和は $S_{i+L} - S_i$
	- この問題の条件は、妥当な任意の $i$ に対して $(S_{i+L} - S_i) \equiv 0 \pmod M$
	- 隣接する2つの部分配列を考えてみる
		- $a_i = A_i, A_{i+1}, ..., A_{i+L-1}$
		- $a_{i+1} = A_{i+1}, A_{i+2}, ..., A_{i+L}$
	- 両方の我が $M$ で割り切れるということは
		- $A_i + A_{i+1} + ... + A_{i+L-1} \equiv 0 \pmod M$
		- $A_{i+1} + A_{i+2} + ... + A_{i+L} \equiv 0 \pmod M$
	- 両者を引くと
		- $A_{i+L} - A_i \equiv 0 \pmod M$
	- よって $L$ 個離れた要素は同じ余りを持つ必要がある
	- この法則を使って配列を $L$ 個間隔の要素のグループに分ける
		- $G_0 = A_0, A_{0+L}, A_{0+2L}, ...$
		- $G_1 = A_1, A_{1+L}, A_{1+2L}, ...$
		- $G_2 = A_2, A_{2+L}, A_{2+2L}, ...$
		- こうすると各グループ内の要素はすべて合同（ここでは同じ値）にする必要がある
	- こうするとこの問題は、各グループで現在の値を目標値まで増やすコストを求める問題になる
	- 目標値は $0$ ~ $M-1$

## 実装

### $L$ 個飛ばしのグループごとにコストを最小化

```
var n, m, l int
var a []int

total_cost = 0
for g = 0; g < l; g++ {
	var vals []int
	for i = g; i < n; i += l {
		vals.push(vals, a[i])
	}
	if len(vals) {
		continue
	}
	group_cost = INF
	for target = 0; target < m; target++ {
		cost = 0
		for _, v = range vals {
			cost += (target - v + m) % m
		}
		group_cost = Min(group_cost, cost)
	}
	total_cost += group_cost
}
print(total_cost)
```

- この方法では各グループは最適化されるが、各部分配列の要素 $L$ 個の総和が $M$ の倍数になることは保証されない
- あるグループの選んだ目標余り値によって他のグループの目標余り値を変えなければならない、というところに目が行くと、DP のイメージに辿り着けそう

### $L$ 個飛ばしのグループごとに目標余り値ごとのコストを計算したうえで DP

```
var n, m, l int
var a []int

groups = [l][]int{}
for i = 0; i < n; i++ {
	groups[i%l] = append(groups[i%l], a[i])
}

costs = [l][]int{}
for gi = 0; gi < l; gi++ {
	costs[gi] = [m]int{}
	for t = 0; t < m; t++ {
		cost = 0
		for _, v = range groups[gi] {
			cost += (target - v + m) % m
		}
		costs[g][target] = cost
	}
}

// dp[i][s] = 最初の i 個のグループの目標値を決めて、その和が s (mod M) のときの最小コスト
dp = [l+1][]int{}
for i = 0; i <= l; i++ {
	dp[i] = [m]int{}
	for j := 0; j < m; j++ {
		dp[i][j] = MAX_INT
	}
}
dp[0][0] = 0

for gi = 0; gi < l; gi++ {
	for prev_sum = 0; prev_sum < m; prev_sum++ {
		if dp[gi][prev_sum] == MAX_INT {
			continue
		}
		for t = 0; t < m; t++ {
			new_sum = (prev_sum + t) % m
			new_cost = dp[g][prev_sum] + costs[g][t]
			ChooseMin(&dp[g+1][new_sum], new_cost)
		}
	}	
}

print(dp[l][0])
```
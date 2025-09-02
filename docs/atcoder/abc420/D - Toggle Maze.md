
| key  | value                                             |
| ---- | ------------------------------------------------- |
| url  | https://atcoder.jp/contests/abc420/tasks/abc420_d |
| tags | #bfs                                              |

## 考察

- 計算量
	- これはどう考えるべきか分からない
	- 全マスに1回ずつ移動したとして $500 \times 500 = 2.5 \times 10^4$ ？
- dp だと思うんだけどなぜか？
	- これは最小化する問題
	- 移動回数 $t$ として、 $t$ のときにゴールマスにいることができるならばその操作回数の最小値、できないなら $-1$ を返す
	- どこでアルゴリズムをストップできる？
	- `S` から移動を始めて、各点に移動できる最小回数を、メモ化しながら探索を進めていく
	- BFS すると、探索とメモ化を効率的に両立できそう
	- `S` の点をコスト 0 として始める、その他の点は全部 -1、すべてのマスについてそこを去る時点で S からそこに至るコストが最小化されていれば、あとは探索終了時点の G のコストを出力すれば良いことになるか
	- 探索中、段階ごとにスイッチマスの通過数を保持しておく必要がある、というかスイッチマスの通過数が偶数／奇数別で最小コストを持っておく必要がありそう

## 実装

```go
func Solve(r *Reader, w *Writer) {
	H, W := r.Int(), r.Int()
	g := make([][]byte, H)
	var start, goal [2]int
	for i := 0; i < H; i++ {
		g[i] = r.Bytes()
		for j := 0; j < W; j++ {
			if g[i][j] == 'S' {
				start = [2]int{i, j}
			} else if g[i][j] == 'G' {
				goal = [2]int{i, j}
			}
		}
	}
	
	visited := make([][][]bool, H)
	for i := range visited {
		visited[i] = make([][]bool, W)
		for j := range visited[i] {
			visited[i][j] = make([]bool, 2)
		}
	}
	
	type state struct {
		pos [2]int
		door int
		steps int
	}
	posEquals := func (a, b [2]int) bool {
		return a[0] == b[0] && a[1] && b[1]
	}
	posMove := func (a [2]int, d [2]int) {
		return [2]int{a[0]+dx, a[1]+dy}
	}

	queue := []state{{start, 0, 0}}	
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if posEquals(cur, goal) {
			w.Println(cur.steps)
			return
		}
		for _, d := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			pos := posMove(cur.pos, d)
			if pos[0] < 0 || pos[0] >= H || pos[1] < 0 || pos[1] >= W {
				continue
			}		
			
			door := cur.door
			switch g[pos[0]][pos[1]] {
				case '#':
					continue
				case 'x':
					if cur.door == 0 {
						continue
					}
				case 'o':
					if cur.door == 1 {
						continue
					}
				case '?':
					door = 1 - door
			}
			
			if !visited[pos[0]][pos[1]][door] {
				visited[pos[0]][pos[1]][door] = true
				queue = append(queue, state{pos, door, cur.steps+1})
			}
		}
	}
}
```

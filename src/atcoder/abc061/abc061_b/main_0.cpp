#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

int main() {
  int N, M;
  cin >> N >> M;

  vector<vector<int>> g(N, vector<int>(0, 0));
  for (int i = 0; i < M; i++) {
    int a, b;
    cin >> a >> b;
    a--;
    b--;
    g[a].push_back(b);
    g[b].push_back(a);
  }

  for (int i = 0; i < N; i++) {
    cout << g[i].size() << endl;
  }
}

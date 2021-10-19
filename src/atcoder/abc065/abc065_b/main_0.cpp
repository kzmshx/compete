#include <bits/stdc++.h>

using namespace std;

int main() {
  int N;
  cin >> N;

  vector<int> A(N);
  for (int i = 0; i < N; i++) {
    int a;
    cin >> a;
    A[i] = a - 1;
  }

  vector<bool> visited(N, false);
  int next = 0;
  int ans = 0;
  while (next != 1) {
    visited[next] = true;
    ans++;
    next = A[next];
    if (visited[next]) {
      ans = -1;
      break;
    }
  }

  cout << ans << endl;
}

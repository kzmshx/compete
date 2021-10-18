#define _GLIBCXX_DEBUG
#include <bits/stdc++.h>

using namespace std;

int main() {
  int N, M;
  cin >> N >> M;

  vector<int> R(N * N, 0);
  for (int i = 0; i < M; i++) {
    int A, B;
    cin >> A >> B;
    R[N * (A - 1) + B - 1] = 1;
    R[N * (B - 1) + A - 1] = -1;
  }

  for (int i = 0; i < N; i++) {
    for (int j = 0; j < N; j++) {
      switch (R[i * N + j]) {
      case -1:
        cout << 'x';
        break;
      case 0:
        cout << '-';
        break;
      case 1:
        cout << 'o';
        break;
      }
      if (j < N - 1) {
        cout << ' ';
      }
    }
    cout << endl;
  }
}

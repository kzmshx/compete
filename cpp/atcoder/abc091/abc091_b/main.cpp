#include <bits/stdc++.h>

using namespace std;

int main() {
  int N;
  cin >> N;
  vector<string> S(N);
  for (int i = 0; i < N; i++) {
    cin >> S[i];
  }

  int M;
  cin >> M;
  vector<string> T(M);
  for (int i = 0; i < M; i++) {
    cin >> T[i];
  }

  int ans = 0;
  while (0 < S.size()) {
    string s = S[0];
    int sc = 0;
    while (find(S.begin(), S.end(), s) != S.end()) {
      S.erase(find(S.begin(), S.end(), s));
      sc++;
    }
    int tc = 0;
    while (find(T.begin(), T.end(), s) != T.end()) {
      T.erase(find(T.begin(), T.end(), s));
      tc++;
    }
    ans = max(max(sc - tc, 0), ans);
  }

  cout << ans << endl;
}

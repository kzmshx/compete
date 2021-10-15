#define _GLIBCXX_DEBUG
#include <iostream>
#include <unordered_map>

using namespace std;

int main() {
  int N, M;
  unordered_map<string, int> m;

  cin >> N;
  for (int i = 0; i < N; i++) {
    string s;
    cin >> s;
    m[s]++;
  }
  cin >> M;
  for (int i = 0; i < M; i++) {
    string t;
    cin >> t;
    m[t]--;
  }

  int ans = 0;
  for (auto &&p : m) {
    ans = max(p.second, ans);
  }
  cout << ans << endl;
}

#define _GLIBCXX_DEBUG
#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

int main() {
  int N;
  cin >> N;

  vector<int> D(N);
  for (int i = 0; i < N; i++) {
    cin >> D[i];
  }

  sort(D.begin(), D.end());

  int ans = 1;
  int cur = D[0];
  for (int i = 1; i < N; i++) {
    if (D[i] != cur) {
      ans++;
      cur = D[i];
    }
  }

  cout << ans << endl;
}

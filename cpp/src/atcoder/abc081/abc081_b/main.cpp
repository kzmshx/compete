#include <bits/stdc++.h>

using namespace std;

int main() {
  int n;
  cin >> n;
  int ans = 30;
  for (int i = 0; i < n; i++) {
    int a, v = 0;
    cin >> a;
    while (a % 2 == 0) {
      a /= 2;
      v++;
    }
    ans = min(ans, v);
  }
  cout << ans << endl;
}

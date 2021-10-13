#include <bits/stdc++.h>

using namespace std;

int main() {
  int s;
  cin >> s;

  int ans = 0;
  while (0 < s) {
    ans += s % 10;
    s /= 10;
  }

  cout << ans << endl;
}

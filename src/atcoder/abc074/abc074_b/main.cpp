#include <bits/stdc++.h>

using namespace std;

int main() {
  int n, k;
  cin >> n >> k;
  int a = 0;
  for (int i = 0; i < n; i++) {
    int x;
    cin >> x;
    a += 2 * min(k - x, x - 0);
  }
  cout << a << endl;
}

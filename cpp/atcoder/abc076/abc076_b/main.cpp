#include <bits/stdc++.h>

using namespace std;

int main() {
  int n, k;
  cin >> n >> k;

  int a = 1;
  for (int i = 0; i < n; i++) {
    a = min(2 * a, a + k);
  }
  cout << a << endl;
}

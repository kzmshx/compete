#include <bits/stdc++.h>

using namespace std;

int main() {
  vector<int> A(5);
  cin >> A[0];
  for (int i = 0; i < 4; i++) {
    cin >> A[i + 1];
    if (A[i] == A[i + 1]) {
      cout << "YES" << endl;
      return 0;
    }
  }

  cout << "NO" << endl;
}

#include <bits/stdc++.h>

using namespace std;

int main() {
  int n;
  cin >> n;

  int v = n;
  int digit_sum = 0;
  while (0 < v) {
    digit_sum += v % 10;
    v /= 10;
  }

  cout << (n % digit_sum == 0 ? "Yes" : "No") << endl;
}

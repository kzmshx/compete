#include <bits/stdc++.h>

using namespace std;

/**
 * https://cpprefjp.github.io/reference/algorithm/sort.html
 * https://cpprefjp.github.io/reference/functional/greater.html
 */
int main() {
  int N;
  cin >> N;

  vector<int> A(N);
  for (int i = 0; i < N; i++) {
    cin >> A[i];
  }

  sort(A.begin(), A.end(), greater<>());

  int a = 0;
  int b = 0;
  for (int i = 0; i < N; i++) {
    if (i % 2 == 0) {
      a += A[i];
    } else {
      b += A[i];
    }
  }

  cout << a - b << endl;
}

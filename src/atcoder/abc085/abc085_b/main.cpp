#include <bitset>
#include <iostream>

using namespace std;

/**
 * https://cpprefjp.github.io/reference/bitset/bitset.html
 */
int main() {
  int N;
  cin >> N;
  bitset<101> D;
  for (int i = 0; i < N; i++) {
    int d;
    cin >> d;
    D.set(d);
  }
  cout << D.count() << endl;
}

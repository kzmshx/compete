#include <bits/stdc++.h>

using namespace std;

int main() {
  int H, W;
  cin >> H >> W;
  cout << string(W + 2, '#') << endl;
  for (int i = 0; i < H; i++) {
    string a;
    cin >> a;
    cout << '#' << a << '#' << endl;
  }
  cout << string(W + 2, '#') << endl;
}

#include <bits/stdc++.h>

using namespace std;

int main() {
  int a, b, c, d;
  cin >> a >> b >> c >> d;
  cout << (a + b == c + d ? "Balanced" : (a + b < c + d ? "Right" : "Left")) << endl;
}

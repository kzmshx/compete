#include <bits/stdc++.h>

using namespace std;

int main() {
  int h, w;
  cin >> h >> w;
  vector<string> f(h);
  for (int i = 0; i < h; i++) {
    cin >> f[i];
  }
  for (int i = 0; i < h; i++) {
    for (int j = 0; j < w; j++) {
      if (f[i][j] == '#') {
        cout << f[i][j];
        continue;
      }
      int c = 0;
      for (int a = i - 1; a <= i + 1; a++)
        for (int b = j - 1; b <= j + 1; b++)
          if (!(a == i && b == j) && 0 <= a && a < h && 0 <= b && b < w && f[a][b] == '#')
            c++;
      cout << c;
    }
    cout << endl;
  }
}

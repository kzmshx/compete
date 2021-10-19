#include <bits/stdc++.h>

using namespace std;

int main() {
  size_t H, W;
  cin >> H >> W;

  vector<char> f(H * W);
  for (size_t i = 0; i < H * W; ++i) {
    cin >> f[W * (i / W) + i % W];
  }

  for (size_t i = 0; i < H; ++i) {
    bool skipRow = true;
    for (size_t x = 0; x < W; ++x) {
      if (f[W * i + x] != '.') {
        skipRow = false;
      }
    }
    if (skipRow) {
      continue;
    }

    for (size_t j = 0; j < W; ++j) {
      bool skipCol = true;
      for (size_t y = 0; y < H; ++y) {
        if (f[W * y + j] != '.') {
          skipCol = false;
          break;
        }
      }
      if (skipCol) {
        continue;
      }
      cout << f[W * i + j];
    }
    cout << endl;
  }
}

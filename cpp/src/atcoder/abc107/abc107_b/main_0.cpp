#include <bits/stdc++.h>

using namespace std;

int main() {
  int H, W;
  cin >> H >> W;

  vector<vector<char>> field(H, vector<char>(W));
  for (int i = 0; i < H; ++i) {
    for (int j = 0; j < W; ++j) {
      cin >> field[i][j];
    }
  }

  for (int i = 0; i < H; ++i) {
    bool skipRow = true;
    for (int a = 0; a < W; ++a) {
      if (field[i][a] != '.') {
        skipRow = false;
        break;
      }
    }
    if (skipRow) {
      continue;
    }

    for (int j = 0; j < W; ++j) {
      bool skipCol = true;
      for (int a = 0; a < H; ++a) {
        if (field[a][j] != '.') {
          skipCol = false;
          break;
        }
      }
      if (skipCol) {
        continue;
      }
      cout << field[i][j];
    }
    cout << endl;
  }
}

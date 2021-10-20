#include <bits/stdc++.h>

using namespace std;

void visit_neighbours(vector<vector<char>> &field, int x, int y) {
    if (x < 0 || 10 <= x || y < 0 || 10 <= y || field[y][x] == 'x') {
        return;
    }

    field[y][x] = 'x';
    visit_neighbours(field, x - 1, y);
    visit_neighbours(field, x + 1, y);
    visit_neighbours(field, x, y - 1);
    visit_neighbours(field, x, y + 1);
}

int main() {
    vector<vector<char>> field(10, vector<char>(10));
    for (auto &r : field) {
        for (auto &c : r) {
            cin >> c;
        }
    }

    for (int i = 0; i < 10; ++i) {
        for (int j = 0; j < 10; ++j) {
            if (field[i][j] == 'o') {
                continue;
            }

            vector<vector<char>> f(10, vector<char>(10));
            for (auto k = 0; k < field.size(); ++k) {
                copy(field[k].begin(), field[k].end(), f[k].begin());
            }
            f[i][j] = 'o';

            visit_neighbours(f, j, i);

            bool ok = true;
            for (auto &r : f) {
                for (auto &c : r) {
                    if (c == 'o') {
                        ok = false;
                        break;
                    }
                }
                if (!ok) {
                    break;
                }
            }
            if (ok) {
                cout << "YES" << endl;
                return 0;
            }
        }
    }

    cout << "NO" << endl;
}

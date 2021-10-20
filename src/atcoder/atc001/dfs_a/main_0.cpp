#include <bits/stdc++.h>

using namespace std;

bool can_reach_goal(vector<vector<char>> &field, int x, int y, vector<vector<bool>> &visited) {
    if (x < 0 || field[0].size() <= x || y < 0 || field.size() <= y || field[y][x] == '#' || visited[y][x]) {
        return false;
    }
    if (field[y][x] == 'g') {
        return true;
    }
    visited[y][x] = true;
    return can_reach_goal(field, x - 1, y, visited)
        || can_reach_goal(field, x + 1, y, visited)
        || can_reach_goal(field, x, y - 1, visited)
        || can_reach_goal(field, x, y + 1, visited);
}

int main() {
    int height, width;
    cin >> height >> width;

    vector<vector<char>> field(height, vector<char>(width));
    int start_x, start_y;
    for (int h = 0; h < height; ++h) {
        for (int w = 0; w < width; ++w) {
            cin >> field[h][w];
            if (field[h][w] == 's') {
                start_x = w;
                start_y = h;
            }
        }
    }

    vector<vector<bool>> visited(height, vector<bool>(width, false));

    cout << (can_reach_goal(field, start_x, start_y, visited) ? "Yes" : "No") << endl;
}

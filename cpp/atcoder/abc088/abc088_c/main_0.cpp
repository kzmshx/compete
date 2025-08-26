#include <bits/stdc++.h>

using namespace std;

using i32 = int32_t;
using i64 = int64_t;
using u32 = uint32_t;
using u64 = uint64_t;
using f32 = float;
using f64 = double;

template<typename T>
bool choose_min(T &min, const T &value) {
    if (min > value) {
        min = value;
        return true;
    }
    return false;
}

template<typename T>
bool choose_max(T &max, const T &value) {
    if (max < value) {
        max = value;
        return true;
    }
    return false;
}

template<typename T, typename = enable_if_t<is_integral_v<T>>>
bool is_prime(const T &integer) {
    if (integer == 2) {
        return true;
    }
    if (integer <= 1 || integer % 2 == 0) {
        return false;
    }
    for (int v = 3; v <= sqrt(integer); v += 2) {
        if (integer % v == 0) {
            return false;
        }
    }
    return true;
}

void run();

int main() {
    cin.tie(0);
    ios::sync_with_stdio(false);
    run();
}

bool solve(array<array<i32, 3>, 3> &grid) {
    for (i32 y = 0; y < 2; y++) {
        if (!(grid[y][0] - grid[y + 1][0] == grid[y][1] - grid[y + 1][1] && grid[y][1] - grid[y + 1][1] == grid[y][2] - grid[y + 1][2])) {
            return false;
        }
    }
    for (i32 x = 0; x < 2; x++) {
        if (!(grid[0][x] - grid[0][x + 1] == grid[1][x] - grid[1][x + 1] && grid[2][x] - grid[2][x + 1] == grid[2][x] - grid[2][x + 1])) {
            return false;
        }
    }
    return true;
}

void run() {
    array<array<i32, 3>, 3> grid;
    for (i32 i = 0; i < 3; i++) {
        for (i32 j = 0; j < 3; j++) {
            cin >> grid[i][j];
        }
    }

    cout << (solve(grid) ? "Yes" : "No") << endl;
}

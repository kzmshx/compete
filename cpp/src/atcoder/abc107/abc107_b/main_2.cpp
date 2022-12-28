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

void run() {
    i32 H, W;
    cin >> H >> W;

    vector<vector<char>> f(H, vector<char>(W));
    for (i32 i = 0; i < H; i++) {
        for (i32 j = 0; j < W; j++) {
            cin >> f[i][j];
        }
    }

    map<i32, bool> skip_y, skip_x;
    for (i32 i = 0; i < H; i++) {
        bool skip = true;
        for (i32 j = 0; j < W; j++) {
            if (f[i][j] != '.') {
                skip = false;
            }
        }
        if (skip) {
            skip_y[i] = true;
        }
    }
    for (i32 i = 0; i < W; i++) {
        bool skip = true;
        for (i32 j = 0; j < H; j++) {
            if (f[j][i] != '.') {
                skip = false;
            }
        }
        if (skip) {
            skip_x[i] = true;
        }
    }

    for (i32 i = 0; i < H; i++) {
        if (skip_y[i]) {
            continue;
        }

        for (i32 j = 0; j < W; j++) {
            if (skip_x[j]) {
                continue;
            }

            cout << f[i][j];
        }

        cout << endl;
    }
}

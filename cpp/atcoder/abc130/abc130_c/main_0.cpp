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

pair<f64, i32> solve(i32 W, i32 H, i32 x, i32 y) {
    return make_pair<f64, i32>(1.0 * W * H / 2, x * 2 == W && y * 2 == H ? 1 : 0);
}

void run() {
    i32 W, H, x, y;
    cin >> W >> H >> x >> y;
    pair<f64, i32> answer = solve(W, H, x, y);
    printf("%.6f %d\n", answer.first, answer.second);
}

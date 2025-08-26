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

tuple<i32, i32, i32> solve(i32 N, i32 Y) {
    Y /= 1000;
    for (i32 x = 0; 10 * x <= Y && x <= N; x++) {
        for (i32 y = 0; 10 * x + 5 * y <= Y && x + y <= N; y++) {
            if (10 * x + 5 * y + N - (x + y) == Y) {
                return make_tuple(x, y, N - (x + y));
            }
        }
    }
    return make_tuple(-1, -1, -1);
}

void run() {
    i32 N, Y;
    cin >> N >> Y;

    auto answer = solve(N, Y);
    cout << get<0>(answer) << " " << get<1>(answer) << " " << get<2>(answer) << endl;
}

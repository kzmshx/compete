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

i64 solve(i64 L, i64 R) {
    i64 min_mod = INT64_MAX;
    for (i64 i = L; i < min(R, L + 2018); i++) {
        for (i64 j = i + 1; j <= min(R, L + 2018); j++) {
            choose_min(min_mod, ((i % 2019) * (j % 2019)) % 2019);
        }
    }
    return min_mod;
}

void run() {
    i64 L, R;
    cin >> L >> R;
    cout << solve(L, R) << endl;
}

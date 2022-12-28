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
    i64 N;
    cin >> N;

    vector<i64> acm(N, 0);
    i64 sum = 0;
    for (i64 i = 0; i < N; i++) {
        i64 A;
        cin >> A;
        sum += A;
        acm[i] = A;
        if (0 < i) {
            acm[i] += acm[i - 1];
        }
    }

    i64 min_side_diff = INT64_MAX;
    for (i64 i = 0; i < N; i++) {
        choose_min(min_side_diff, abs(sum - 2 * acm[i]));
    }

    cout << min_side_diff << endl;
}

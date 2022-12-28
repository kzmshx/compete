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

i64 solve(vector<i64> &timings, i64 T) {
    i64 ans = timings[timings.size() - 1] + T;
    for (i64 i = 0; i < (i64) timings.size() - 1; i++) {
        ans -= max(timings[i + 1] - timings[i] - T, (i64) 0);
    }
    return ans;
}

void run() {
    i64 N, T;
    cin >> N >> T;

    vector<i64> timings(N + 1);
    timings[0] = 0;
    for (i64 i = 1; i <= N; i++) {
        cin >> timings[i];
    }

    cout << solve(timings, T) << endl;
}

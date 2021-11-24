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
    return timings[timings.size() - 1] + T - accumulate(timings.begin(), timings.end(), make_pair(0, 0), [T](pair<i64, i64> acm, i64 &timing) {
                                                 return make_pair(acm.first + max(timing - acm.second - T, (i64) 0), timing);
                                             }).first;
}

void run() {
    i64 N, T;
    cin >> N >> T;

    vector<i64> timings(N);
    for (auto &timing : timings) {
        cin >> timing;
    }

    cout << solve(timings, T) << endl;
}

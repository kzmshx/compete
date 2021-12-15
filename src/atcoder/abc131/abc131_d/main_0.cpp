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

bool solve(vector<pair<i64, i64>> AB) {
    sort(AB.begin(), AB.end(), [](pair<i64, i64> &a, pair<i64, i64> &b) -> bool {
        return a.second < b.second || (a.second == b.second && a.first < b.first);
    });
    i64 current = 0;
    for (auto &ab : AB) {
        current += ab.first;
        if (current > ab.second) {
            return false;
        }
    }
    return true;
}

void run() {
    i64 N;
    cin >> N;

    vector<pair<i64, i64>> AB(N, make_pair(0, 0));
    for (auto &ab : AB) {
        cin >> ab.first >> ab.second;
    }

    cout << (solve(AB) ? "Yes" : "No") << endl;
}

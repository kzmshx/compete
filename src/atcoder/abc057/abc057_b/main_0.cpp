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
    i32 N, M;
    cin >> N >> M;

    vector<pair<i32, i32>> ab(N), cd(M);
    for (i32 i = 0; i < N; i++) {
        pair<i32, i32> p;
        cin >> p.first >> p.second;
        ab[i] = p;
    }
    for (i32 i = 0; i < M; i++) {
        pair<i32, i32> p;
        cin >> p.first >> p.second;
        cd[i] = p;
    }

    for (i32 i = 0; i < N; i++) {
        i32 d = INT32_MAX;
        i32 answer = 0;
        for (i32 j = 0; j < M; j++) {
            i32 v = abs(ab[i].first - cd[j].first) + abs(ab[i].second - cd[j].second);
            if (v < d) {
                d = v;
                answer = j + 1;
            }
        }
        cout << answer << endl;
    }
}

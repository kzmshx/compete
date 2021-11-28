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

i32 solve(i32 N, vector<pair<i32, i32>> SC) {
    for (i32 answer = 0; answer < 1000; answer++) {
        string s = to_string(answer);
        if (s.length() != N) {
            continue;
        }

        bool ok = true;
        for (i32 i = 0; i < (i32) SC.size(); i++) {
            for (i32 j = 0; j < N; j++) {
                if (SC[i].first == j && SC[i].second != s[j] - '0') {
                    ok = false;
                }
            }
        }
        if (ok) {
            return answer;
        }
    }

    return -1;
}

void run() {
    i32 N, M;
    cin >> N >> M;

    vector<pair<i32, i32>> SC(M);
    for (auto &sc : SC) {
        cin >> sc.first >> sc.second;
        sc.first--;
    }

    cout << solve(N, SC) << endl;
}

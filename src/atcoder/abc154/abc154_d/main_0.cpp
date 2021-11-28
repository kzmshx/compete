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

f64 solve(vector<i64> P, i64 K) {
    f64 max_sum_of_exps = 0;

    f64 current_sum_of_exps = 0;
    for (i64 j = 0; j < K; j++) {
        current_sum_of_exps += (f64) (P[j] + 1) / 2;
    }

    max_sum_of_exps = current_sum_of_exps;

    for (i64 i = 1; i <= (i64) P.size() - K; i++) {
        current_sum_of_exps -= (f64) (P[i - 1] + 1) / 2;
        current_sum_of_exps += (f64) (P[i - 1 + K] + 1) / 2;
        choose_max(max_sum_of_exps, current_sum_of_exps);
    }

    return max_sum_of_exps;
}

void run() {
    i64 N, K;
    cin >> N >> K;

    vector<i64> P(N);
    for (auto &p : P) {
        cin >> p;
    }

    printf("%.12f\n", solve(P, K));
}

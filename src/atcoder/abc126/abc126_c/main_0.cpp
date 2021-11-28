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

f64 solve(i32 N, i32 K) {
    f64 answer = 0.;
    for (i64 i = 1; i <= N; i++) {
        f64 current = i;
        i32 n = 0;
        while (current < K) {
            current *= 2;
            n++;
        }
        answer += pow(0.5, n) / N;
    }
    return answer;
}

void run() {
    i32 N, K;
    cin >> N >> K;
    printf("%.12f\n", solve(N, K));
}

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

i64 combinations(i64 n, i64 k) {
    if (n < k) {
        return -1;
    }
    if (n == k) {
        return 1;
    }
    if (n / k < 2) {
        k = n - k;
    }
    auto f = [](i64 a, i64 b) -> i64 {
        for (i64 i = a - 1; i >= b; i--) {
            a *= i;
        }
        return a;
    };
    return f(n, n - k + 1) / f(k, 1);
}

i64 solve(vector<i64> A, i64 P) {
    i64 even = 0, odd = 0;
    for (const auto &a : A) {
        if (a % 2 == 0) {
            even++;
        }
    }
    odd = (i64) A.size() - even;

    i64 odd_patterns = (P == 0 ? 1 : 0);
    for (i64 i = (P == 0 ? 2 : 1); i <= odd; i += 2) {
        odd_patterns += combinations(odd, i);
    }
    return odd_patterns * pow(2, even);
}

void run() {
    i64 N, P;
    cin >> N >> P;

    vector<i64> A(N);
    for (auto &a : A) {
        cin >> a;
    }

    cout << solve(A, P) << endl;
}

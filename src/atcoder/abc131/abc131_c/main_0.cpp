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

i64 f(i64 x, i64 C, i64 D) {
    if (x == 0) {
        return 0;
    }

    i64 LCM = lcm(C, D);
    return x - x / C - x / D + x / LCM;
}

i64 solve(i64 A, i64 B, i64 C, i64 D) {
    i64 LCM = lcm(C, D);
    auto f = [C, D, LCM](i64 x) { return x - x / C - x / D + x / LCM; };
    return f(B) - f(A - 1);
}

void run() {
    i64 A, B, C, D;
    cin >> A >> B >> C >> D;
    cout << solve(A, B, C, D) << endl;
}

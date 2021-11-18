#include <bits/stdc++.h>

using namespace std;
using i32 = int32_t;
using i64 = int64_t;

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
    string S;
    i64 K;
    cin >> S >> K;

    i64 last_one_index = 0;
    while (S[last_one_index] == '1') {
        last_one_index++;
    }

    cout << (K <= last_one_index ? 1 : S[last_one_index] - '0') << endl;
}

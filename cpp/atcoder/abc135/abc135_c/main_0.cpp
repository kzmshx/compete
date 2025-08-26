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
    i64 N;
    cin >> N;

    vector<i64> A(N + 1);
    for (i64 i = 0; i < N + 1; i++) {
        cin >> A[i];
    }

    i64 answer = 0;
    for (i64 i = 0; i < N; i++) {
        i64 B;
        cin >> B;

        i64 v1 = min(A[i], B);
        answer += v1;
        A[i] -= v1;
        B -= v1;

        i64 v2 = min(A[i + 1], B);
        answer += v2;
        A[i + 1] -= v2;
    }
    cout << answer << endl;
}

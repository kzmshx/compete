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

i32 solve(vector<i32> A) {
    i32 answer = 1;
    i32 rising = 0;
    for (i32 i = 0; i < (i32) A.size() - 1; i++) {
        if (A[i] < A[i + 1]) {
            if (rising == 0) {
                rising = 1;
            } else if (rising == -1) {
                rising = 0;
                answer++;
            }
        } else if (A[i] > A[i + 1]) {
            if (rising == 0) {
                rising = -1;
            } else if (rising == 1) {
                rising = 0;
                answer++;
            }
        }
    }
    return answer;
}

void run() {
    i32 N;
    cin >> N;

    vector<i32> A(N);
    for (auto &a : A) {
        cin >> a;
    }

    cout << solve(A) << endl;
}

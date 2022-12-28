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

bool solve(vector<i32> A) {
    u32 multiple_of_2_count = 0;
    u32 multiple_of_4_count = 0;
    for (const auto &a : A) {
        if (a % 4 == 0) {
            multiple_of_4_count++;
        } else if (a % 2 == 0) {
            multiple_of_2_count++;
        }
    }
    if (multiple_of_2_count % 2 == 1) {
        multiple_of_2_count--;
    }
    return (A.size() - multiple_of_2_count - multiple_of_4_count) <= multiple_of_4_count + 1;
}

void run() {
    i32 N;
    cin >> N;

    vector<i32> A(N);
    for (auto &a : A) {
        cin >> a;
    }

    cout << (solve(A) ? "Yes" : "No") << endl;
}

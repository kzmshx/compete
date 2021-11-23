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
    i32 N;
    cin >> N;

    vector<vector<int>> chars(N, vector(26, 0));
    for (i32 i = 0; i < N; i++) {
        char c;
        cin >> c;
        chars[i][c - 'a']++;
        if (0 < i) {
            for (i32 j = 0; j < (i32) chars[0].size(); j++) {
                chars[i][j] += chars[i - 1][j];
            }
        }
    }

    i32 max_common_chars = 0;
    for (i32 i = 0; i < N; i++) {
        i32 common_chars = 0;
        for (i32 j = 0; j < (i32) chars[i].size(); j++) {
            if (chars[i][j] > 0 && chars[N - 1][j] - chars[i][j] > 0) {
                common_chars++;
            }
        }
        choose_max(max_common_chars, common_chars);
    }

    cout << max_common_chars << endl;
}

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

i32 count_common_char_count(string a, string b) {
    i32 common_char_count = 0;
    for (char c = 'a'; c <= 'z'; c++) {
        if (count(a.begin(), a.end(), c) > 0 && count(b.begin(), b.end(), c) > 0) {
            common_char_count++;
        }
    }
    return common_char_count;
}

void run() {
    i32 N;
    string S;
    cin >> N >> S;

    int max_common_char_count = 0;
    for (i32 i = 0; i < (i32) S.size(); i++) {
        choose_max(max_common_char_count, count_common_char_count(S.substr(0, i), S.substr(i)));
    }

    cout << max_common_char_count << endl;
}

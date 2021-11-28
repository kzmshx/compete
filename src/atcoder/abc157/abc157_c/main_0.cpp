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

i32 solve(i32 N, vector<pair<i32, i32>> sc) {
    vector<i32> digits(N, -1);

    for (const auto &v : sc) {
        if (v.first == 0 && v.second == 0 && N > 1) {
            return -1;
        }
        if (digits[v.first] != -1 && digits[v.first] != v.second) {
            return -1;
        }
        digits[v.first] = v.second;
    }

    i32 answer = 0;
    for (i32 i = 0; i < (i32) digits.size(); i++) {
        answer += digits[i] != -1 ? digits[i] : (i > 0 ? 0 : (digits.size() == 1 ? 0 : 1));
        if (i < (i32) digits.size() - 1) {
            answer *= 10;
        }
    }

    return answer;
}

void run() {
    i32 N, M;
    cin >> N >> M;

    vector<pair<i32, i32>> sc(M, make_pair(0, 0));
    for (auto &v : sc) {
        i32 s, c;
        cin >> s >> c;
        v.first = s - 1;
        v.second = c;
    }

    cout << solve(N, sc) << endl;
}

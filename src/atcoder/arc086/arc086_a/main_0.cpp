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

i64 solve(vector<i64> A, i64 K) {
    map<i64, i64> count_of_numbers = {};
    for (const auto &a : A) {
        count_of_numbers[a]++;
    }

    vector<pair<i64, i64>> vectorized_count_of_numbers;
    copy(count_of_numbers.begin(), count_of_numbers.end(), back_inserter<vector<pair<i64, i64>>>(vectorized_count_of_numbers));

    sort(vectorized_count_of_numbers.begin(), vectorized_count_of_numbers.end(), [](const pair<i64, i64> &a, const pair<i64, i64> &b) {
        return a.second < b.second;
    });

    i64 answer = 0;
    for (i32 i = 0; i < (i32) vectorized_count_of_numbers.size() - K; i++) {
        answer += vectorized_count_of_numbers[i].second;
    }

    return answer;
}

void run() {
    i64 N, K;
    cin >> N >> K;

    vector<i64> A(N);
    for (auto &a : A) {
        cin >> a;
    }

    cout << solve(A, K) << endl;
}

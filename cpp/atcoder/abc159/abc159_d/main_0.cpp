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

vector<i64> solve(vector<i64> &A) {
    unordered_map<i64, i64> counts = {};
    for (i64 i = 0; i < (i64) A.size(); i++) {
        counts[A[i]]++;
    }
    i64 sum = 0;
    for (i64 i = 1; i <= (i64) A.size(); i++) {
        sum += counts[i] * (counts[i] - 1) / 2;
    }

    vector<i64> answers(A.size());
    for (i64 i = 0; i < (i64) A.size(); i++) {
        answers[i] = sum - counts[A[i]] * (counts[A[i]] - 1) / 2 + (counts[A[i]] - 1) * (counts[A[i]] - 2) / 2;
    }
    return answers;
}

void run() {
    i64 N;
    cin >> N;

    vector<i64> A(N);
    for (auto &a : A) {
        cin >> a;
    }

    vector<i64> answer = solve(A);
    copy(answer.begin(), answer.end(), ostream_iterator<i64>(cout, "\n"));
}

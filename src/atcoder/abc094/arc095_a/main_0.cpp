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

vector<i32> solve(i32 N, vector<i32> X) {
    vector<i32> c(N);
    copy(X.begin(), X.end(), c.begin());
    sort(c.begin(), c.end());

    map<i32, i32> indices;
    for (i32 i = N - 1; i >= 0; i--) {
        indices[c[i]] = i;
    }

    vector<i32> answer(N);
    for (i32 i = 0; i < N; i++) {
        i32 j = indices[X[i]];
        answer[i] = j < N / 2 ? c[N / 2] : c[N / 2 - 1];
    }

    return answer;
}

void run() {
    i32 N;
    cin >> N;

    vector<i32> X(N);
    for (auto &x : X) {
        cin >> x;
    }

    auto answer = solve(N, X);
    for (auto &a : answer) {
        cout << a << endl;
    }
}

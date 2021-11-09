#include <bits/stdc++.h>

using namespace std;

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

int64_t lucas_number_inner(size_t i, vector<int64_t> &memo) {
    if (i < memo.size()) {
        return memo[i];
    }

    memo.push_back(lucas_number_inner(i - 2, memo) + lucas_number_inner(i - 1, memo));

    return memo[i];
}

int64_t lucas_number(size_t N) {
    vector<int64_t> memo = {2, 1};
    return lucas_number_inner(N, memo);
}

void run() {
    size_t N;
    cin >> N;
    cout << lucas_number(N) << endl;
}

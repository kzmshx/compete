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

void run() {
    size_t N;
    int T, A;
    cin >> N >> T >> A;

    int min_diff = INT_MAX;
    size_t min_diff_index;
    for (size_t i = 1; i <= N; i++) {
        int H;
        cin >> H;

        int diff = abs(A * 1000 - (T * 1000 - H * 6));
        if (diff < min_diff) {
            min_diff = diff;
            min_diff_index = i;
        }
    }

    cout << min_diff_index << endl;
}

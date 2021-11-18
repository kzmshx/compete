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
    int N;
    cin >> N;

    vector<int> r1(N);
    for (int i = 0; i < N; i++) {
        cin >> r1[i];
    }
    vector<int> r2(N);
    for (int i = 0; i < N; i++) {
        cin >> r2[i];
    }

    for (int i = 1; i < N; i++) {
        int j = N - i - 1;
        r1[i] += r1[i - 1];
        r2[j] += r2[j + 1];
    }

    int max_sum = 0;
    for (int i = 0; i < N; i++) {
        choose_max(max_sum, r1[i] + r2[i]);
    }

    cout << max_sum << endl;
}

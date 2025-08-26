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
    string S, T;
    cin >> S >> T;

    for (size_t i = 0; i < S.size(); i++) {
        bool ok = true;
        for (size_t j = 0; j < S.size(); j++) {
            size_t index = S.size() <= j + i ? j + i - S.size() : j + i;
            if (S[index] != T[j]) {
                ok = false;
            }
        }
        if (ok) {
            cout << "Yes" << endl;
            return;
        }
    }

    cout << "No" << endl;
}

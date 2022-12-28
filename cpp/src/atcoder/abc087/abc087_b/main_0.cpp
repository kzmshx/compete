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
    int A, B, C, X;
    cin >> A >> B >> C >> X;

    int pattern_count = 0;
    for (int a = 0; a <= A; a++) {
        for (int b = 0; b <= B; b++) {
            int rest = X - 500 * a - 100 * b;
            if (rest < 0) {
                break;
            }
            if (rest / 50 <= C) {
                pattern_count++;
            }
        }
    }

    cout << pattern_count << endl;
}

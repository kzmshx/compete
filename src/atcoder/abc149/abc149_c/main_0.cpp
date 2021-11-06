#include <bits/stdc++.h>

using namespace std;

template<typename T> bool choose_min(T &min, const T &value);
template<typename T> bool choose_max(T &max, const T &value);

bool is_prime(int value) {
    if (value <= 1) {
        return false;
    }
    if (value == 2) {
        return true;
    }
    if (value % 2 == 0) {
        return false;
    }
    for (int v = 3; v <= sqrt(value); v += 2) {
        if (value % v == 0) {
            return false;
        }
    }
    return true;
}

int main() {
    int X;
    cin >> X;

    while (!is_prime(X)) {
        X++;
    }

    cout << X << endl;
}

template<typename T> bool choose_min(T &min, const T &value) {
    if (min > value) {
        min = value;
        return true;
    }
    return false;
}

template<typename T> bool choose_max(T &max, const T &value) {
    if (max < value) {
        max = value;
        return true;
    }
    return false;
}

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
    string S;
    cin >> S;

    int min_distance_to_753 = INT_MAX;
    for (size_t i = 0; i < S.size() - 2; i++) {
        choose_min(min_distance_to_753, abs(753 - (100 * (S[i] - '0') + 10 * (S[i + 1] - '0') + S[i + 2] - '0')));
    }

    cout << min_distance_to_753 << endl;
}

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

int main() {
    size_t N, x;
    cin >> N >> x;

    vector<size_t> A(N);
    for (size_t i = 0; i < N; i++) {
        cin >> A[i];
    }

    sort(A.begin(), A.end());

    for (size_t i = 0; i < N; i++) {
        if (A[i] <= x) {
            x -= A[i];
        } else {
            cout << i << endl;
            return 0;
        }
    }

    if (x == 0) {
        cout << N << endl;
    } else {
        cout << N - 1 << endl;
    }

    return 0;
}

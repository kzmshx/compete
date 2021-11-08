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
    int N, P[8], Q[8];
    cin >> N;
    for (int i = 0; i < N; i++) {
        cin >> P[i];
    }
    for (int i = 0; i < N; i++) {
        cin >> Q[i];
    }

    vector<int> v;
    for (int i = 0; i < N; i++) {
        v.push_back(i + 1);
    }

    int index = 0, a = -1, b = -1;
    do {
        bool ok_a = true;
        for (int i = 0; i < N; i++) {
            if (v[i] != P[i]) {
                ok_a = false;
                break;
            }
        }
        if (ok_a) {
            a = index;
        }

        bool ok_b = true;
        for (int i = 0; i < N; i++) {
            if (v[i] != Q[i]) {
                ok_b = false;
                break;
            }
        }
        if (ok_b) {
            b = index;
        }

        index++;
    } while (next_permutation(v.begin(), v.end()));

    cout << abs(a - b) << endl;
}

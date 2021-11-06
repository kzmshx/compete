#include <bits/stdc++.h>

using namespace std;

template<typename T> bool choose_min(T &min, const T &value);
template<typename T> bool choose_max(T &max, const T &value);

int main() {
    int N, M, X;
    cin >> N >> M >> X;

    for (int i = 0; i < M; i++) {
        int A;
        cin >> A;

        if (X < A) {
            cout << min(i, M - i) << endl;
            return 0;
        }
    }

    cout << 0 << endl;
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

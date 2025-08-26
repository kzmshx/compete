#include <bits/stdc++.h>

using namespace std;

template<typename T> bool choose_min(T &min, const T &value);
template<typename T> bool choose_max(T &max, const T &value);

int main() {
    vector<int> A(9);
    for (size_t i = 0; i < 9; i++) {
        cin >> A[i];
    }

    size_t N;
    cin >> N;
    for (size_t i = 0; i < N; i++) {
        int b;
        cin >> b;
        for (size_t j = 0; j < 9; j++) {
            if (A[j] == b) {
                A[j] = 0;
            }
        }
    }

    bool ok = false;
    vector<tuple<size_t, size_t, size_t>> bingos = vector<tuple<size_t, size_t, size_t>>{
        make_tuple(0, 1, 2),
        make_tuple(3, 4, 5),
        make_tuple(6, 7, 8),
        make_tuple(0, 3, 6),
        make_tuple(1, 4, 7),
        make_tuple(2, 5, 8),
        make_tuple(0, 4, 8),
        make_tuple(2, 4, 6),
    };
    for (auto bingo : bingos) {
        size_t first = get<0>(bingo), second = get<1>(bingo), third = get<2>(bingo);
        if (A[first] == 0 && A[second] == 0 && A[third] == 0) {
            ok = true;
            break;
        }
    }

    cout << (ok ? "Yes" : "No") << endl;
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

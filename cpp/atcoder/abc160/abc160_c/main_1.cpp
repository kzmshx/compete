#include <bits/stdc++.h>

using namespace std;

template<typename T> bool choose_min(T &min, const T &value);
template<typename T> bool choose_max(T &max, const T &value);

int main() {
    int K, N;
    cin >> K >> N;

    int max_interval = 0;

    int first;
    cin >> first;
    int current_location = first;

    for (size_t i = 0; i < N - 1; i++) {
        int A;
        cin >> A;
        choose_max(max_interval, A - current_location);
        current_location = A;
    }
    choose_max(max_interval, K - current_location + first);

    cout << K - max_interval << endl;
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

#include <bits/stdc++.h>

using namespace std;

template<typename T> bool choose_min(T &min, const T &value);
template<typename T> bool choose_max(T &max, const T &value);

int main() {
    int K, N;
    cin >> K >> N;

    vector<int> A(N);
    for (int i = 0; i < N; i++) {
        cin >> A[i];
    }

    vector<int> intervals(N);
    for (int i = 0; i < N - 1; i++) {
        intervals[i] = A[i + 1] - A[i];
    }
    intervals[N - 1] = K - A[N - 1] + A[0];

    cout << K - *max_element(intervals.begin(), intervals.end()) << endl;
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

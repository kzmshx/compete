#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

template<typename T> bool choose_min(T &min, const T &value);
template<typename T> bool choose_max(T &max, const T &value);

int main() {
    size_t N;
    cin >> N;

    vector<int> D(N);
    for (size_t i = 0; i < N; i++) {
        cin >> D[i];
    }

    sort(D.begin(), D.end());

    int mid_left = D[N / 2 - 1], mid_right = D[N / 2];
    cout << mid_right - mid_left << endl;
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

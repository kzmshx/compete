#include <algorithm>
#include <iostream>

using namespace std;

template<typename T> bool choose_min(T &min, const T &value);
template<typename T> bool choose_max(T &max, const T &value);

int main() {
    size_t N;
    cin >> N;

    vector<pair<int, int>> timings(N);
    for (size_t i = 0; i < N; i++) {
        size_t A;
        cin >> A;
        timings[i] = make_pair(i + 1, A);
    }

    sort(timings.begin(), timings.end(), [](pair<int, int> &a, pair<int, int> &b) -> bool { return a.second < b.second; });

    for (size_t i = 0; i < N; i++) {
        cout << timings[i].first;
        if (i < N - 1) {
            cout << " ";
        }
    }
    cout << endl;
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

#include <bits/stdc++.h>

using namespace std;

template<typename T> bool choose_min(T &min, const T &value);
template<typename T> bool choose_max(T &max, const T &value);

int main() {
    size_t N;
    cin >> N;
    vector<int> Xs(N);
    for (size_t i = 0; i < N; ++i) {
        cin >> Xs[i];
    }

    auto minmax = minmax_element(Xs.begin(), Xs.end());
    int min_cost = INT_MAX;
    for (int p = *minmax.first; p <= *minmax.second; ++p) {
        int cost = 0;
        for (auto X : Xs) {
            cost += (int) pow((double) (X - p), 2);
        }
        chmin(min_cost, cost);
    }

    cout << min_cost << endl;
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

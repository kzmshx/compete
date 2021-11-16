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
    size_t N, K;
    cin >> N >> K;

    vector<int> h(N);
    for (size_t i = 0; i < h.size(); i++) {
        cin >> h[i];
    }

    sort(h.begin(), h.end());

    int min_height_difference = INT_MAX;
    for (size_t i = 0; i <= h.size() - K; i++) {
        choose_min(min_height_difference, h[i + K - 1] - h[i]);
    }

    cout << min_height_difference << endl;
}

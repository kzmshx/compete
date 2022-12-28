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
    size_t N;
    cin >> N;

    vector<tuple<string, int, int>> restaurants(N);
    for (size_t i = 0; i < N; i++) {
        string S;
        int P;
        cin >> S >> P;
        restaurants[i] = make_tuple(S, P, i + 1);
    }

    sort(restaurants.begin(), restaurants.end(), [](const tuple<string, int, int> &a, const tuple<string, int, int> &b) {
        if (get<0>(a) != get<0>(b)) {
            return get<0>(a) < get<0>(b);
        }
        if (get<1>(a) != get<1>(b)) {
            return get<1>(a) > get<1>(b);
        }
        return get<2>(a) < get<2>(b);
    });

    for (const auto &restaurant : restaurants) {
        cout << get<2>(restaurant) << endl;
    }
}

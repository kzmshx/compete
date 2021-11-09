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
    vector<int> cooking_times(5);
    int max_to_10 = 0;
    for (size_t i = 0; i < 5; i++) {
        cin >> cooking_times[i];
        int to_10 = (10 - cooking_times[i] % 10) % 10;
        cooking_times[i] += to_10;
        choose_max(max_to_10, to_10);
    }

    cout << accumulate(cooking_times.begin(), cooking_times.end(), -max_to_10) << endl;
}

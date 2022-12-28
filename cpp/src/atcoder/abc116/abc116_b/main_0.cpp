#include <bits/stdc++.h>

using namespace std;

template<typename T> bool choose_min(T &min, const T &value);
template<typename T> bool choose_max(T &max, const T &value);

int main() {
    int a;
    cin >> a;

    int index = 1;
    unordered_map<int, int> previous_appearances = {};

    while (previous_appearances.find(a) == previous_appearances.end()) {
        previous_appearances.emplace(a, index);
        if (a % 2 == 0) {
            a /= 2;
        } else {
            a = 3 * a + 1;
        }
        index++;
    }

    cout << index << endl;
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

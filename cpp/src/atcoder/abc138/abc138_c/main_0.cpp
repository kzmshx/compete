#include <algorithm>
#include <iostream>

using namespace std;

template<typename T> bool choose_min(T &min, const T &value);
template<typename T> bool choose_max(T &max, const T &value);

int main() {
    size_t N;
    cin >> N;

    vector<double> values(N);
    for (size_t i = 0; i < N; i++) {
        cin >> values[i];
    }

    sort(values.begin(), values.end());

    double current_value = values[0];
    for (size_t i = 1; i < N; i++) {
        current_value = (current_value + values[i]) / 2;
    }

    cout << current_value << endl;
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

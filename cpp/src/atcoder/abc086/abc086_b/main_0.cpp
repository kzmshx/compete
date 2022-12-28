#include <bits/stdc++.h>

using namespace std;

template<typename T> bool choose_min(T &min, const T &value);
template<typename T> bool choose_max(T &max, const T &value);

int main() {
    string a, b;
    cin >> a >> b;

    int value = stoi(a + b);

    bool is_square_number = false;
    for (int i = 1; i <= (int) sqrt(100100); i++) {
        if (i * i == value) {
            is_square_number = true;
            break;
        }
    }

    cout << (is_square_number ? "Yes" : "No") << endl;
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

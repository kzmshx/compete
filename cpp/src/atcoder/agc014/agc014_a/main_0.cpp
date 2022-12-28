#include <bits/stdc++.h>

using namespace std;

template<typename T> bool choose_min(T &min, const T &value);
template<typename T> bool choose_max(T &max, const T &value);

int main() {
    vector<int> numbers(3);
    for (size_t i = 0; i < 3; i++) {
        cin >> numbers[i];
    }
    sort(numbers.begin(), numbers.end());

    vector<tuple<int, int, int>> combinations(0);
    tuple<int, int, int> current_combination = make_tuple(numbers[0], numbers[1], numbers[2]);
    combinations.push_back(current_combination);

    while (numbers[0] % 2 == 0 && numbers[1] % 2 == 0 && numbers[2] % 2 == 0) {
        int A = numbers[0], B = numbers[1], C = numbers[2];

        numbers[0] = (B + C) / 2;
        numbers[1] = (C + A) / 2;
        numbers[2] = (A + B) / 2;
        sort(numbers.begin(), numbers.end());

        current_combination = make_tuple(numbers[0], numbers[1], numbers[2]);
        if (find(combinations.begin(), combinations.end(), current_combination) == combinations.end()) {
            combinations.push_back(current_combination);
        } else {
            cout << -1 << endl;
            return 0;
        }
    }

    cout << combinations.size() - 1 << endl;

    return 0;
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

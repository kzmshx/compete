#include <bits/stdc++.h>

using namespace std;

template<typename T> bool choose_min(T &min, const T &value);
template<typename T> bool choose_max(T &max, const T &value);

int main() {
    string S;
    cin >> S;

    int max_acgt_number_length = 0;
    int current_length = 0;
    for (size_t i = 0; i < S.size(); i++) {
        if (S[i] == 'A' || S[i] == 'C' || S[i] == 'G' || S[i] == 'T') {
            current_length++;
        } else {
            choose_max(max_acgt_number_length, current_length);
            current_length = 0;
        }
    }
    choose_max(max_acgt_number_length, current_length);

    cout << max_acgt_number_length << endl;
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

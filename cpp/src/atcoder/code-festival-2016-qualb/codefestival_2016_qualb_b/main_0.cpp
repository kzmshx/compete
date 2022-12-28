#include <bits/stdc++.h>

using namespace std;

template<typename T> bool choose_min(T &min, const T &value);
template<typename T> bool choose_max(T &max, const T &value);

int main() {
    int N, A, B;
    cin >> N >> A >> B;
    string S;
    cin >> S;

    int domestic_students = 0;
    int overseas_students = 0;
    for (size_t i = 0; i < S.size(); i++) {
        switch (S[i]) {
        case 'a':
            if (domestic_students + overseas_students < A + B) {
                cout << "Yes" << endl;
                domestic_students++;
            } else {
                cout << "No" << endl;
            }
            break;
        case 'b':
            if (domestic_students + overseas_students < A + B && overseas_students < B) {
                cout << "Yes" << endl;
                overseas_students++;
            } else {
                cout << "No" << endl;
            }
            break;
        case 'c':
            cout << "No" << endl;
            break;
        }
    }
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

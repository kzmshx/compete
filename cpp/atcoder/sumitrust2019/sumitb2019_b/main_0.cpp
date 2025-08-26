#include <bits/stdc++.h>

using namespace std;

template<typename T> bool choose_min(T &min, const T &value);
template<typename T> bool choose_max(T &max, const T &value);

int main() {
    /**
     * N に対して X が存在しない場合を考慮しなければならないことに注意
     */
    int N;
    cin >> N;

    int derived = 100 * N / 108;

    if (derived * 108 / 100 == N) {
        cout << derived << endl;
    } else if ((derived + 1) * 108 / 100 == N) {
        cout << derived + 1 << endl;
    } else {
        cout << ":(" << endl;
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

#include <bits/stdc++.h>

using namespace std;

template<typename T> bool choose_min(T &min, const T &value);
template<typename T> bool choose_max(T &max, const T &value);

int main() {
    /**
     * https://blog.hamayanhamayan.com/entry/2020/03/15/001953
     *
     * H または W が 1 のときは、全く動かすことができないので、答えが 1 になる点に注意が必要。
     */
    int64_t H, W;
    cin >> H >> W;

    if (H == 1 || W == 1) {
        cout << 1 << endl;
    } else {
        cout << (H * W + 1) / 2 << endl;
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

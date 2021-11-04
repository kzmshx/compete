#include <bits/stdc++.h>

using namespace std;

template<typename T> bool choose_min(T &min, const T &value);
template<typename T> bool choose_max(T &max, const T &value);

int main() {
    size_t N, M;
    int C;
    cin >> N >> M >> C;

    vector<int> B(M);
    for (size_t i = 0; i < M; i++) {
        cin >> B[i];
    }

    int ac_count = 0;
    for (size_t i = 0; i < N; i++) {
        int sum = C;
        for (const auto &b : B) {
            int a;
            cin >> a;
            sum += a * b;
        }
        if (sum > 0) {
            ac_count++;
        }
    }

    cout << ac_count << endl;
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

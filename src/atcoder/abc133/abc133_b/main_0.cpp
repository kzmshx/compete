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
    size_t N, D;
    cin >> N >> D;

    vector<vector<int>> coordinates(N, vector<int>(D));
    for (size_t i = 0; i < N; i++) {
        for (size_t d = 0; d < D; d++) {
            cin >> coordinates[i][d];
        }
    }

    int answer = 0;
    for (size_t i = 0; i < N; i++) {
        for (size_t j = i + 1; j < N; j++) {
            double sum = 0;
            for (size_t d = 0; d < D; d++) {
                sum += pow(coordinates[i][d] - coordinates[j][d], 2);
            }
            int distance = (int) sqrt(sum);
            if (distance * distance == (int) sum) {
                answer++;
            }
        }
    }

    cout << answer << endl;
}

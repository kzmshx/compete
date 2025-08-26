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

int main() {
    size_t N_A, N_B, M;
    cin >> N_A >> N_B >> M;

    vector<int> A(N_A), B(N_B);
    int A_min = INT_MAX, B_min = INT_MAX;
    for (size_t i = 0; i < N_A; i++) {
        cin >> A[i];
        choose_min(A_min, A[i]);
    }
    for (size_t i = 0; i < N_B; i++) {
        cin >> B[i];
        choose_min(B_min, B[i]);
    }

    int min_cost = A_min + B_min;
    for (size_t i = 0; i < M; i++) {
        size_t x, y;
        int c;
        cin >> x >> y >> c;
        choose_min(min_cost, A[x - 1] + B[y - 1] - c);
    }

    cout << min_cost << endl;
}

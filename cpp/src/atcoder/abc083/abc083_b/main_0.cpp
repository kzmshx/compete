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
    int N, A, B;
    cin >> N >> A >> B;

    int answer = 0;
    for (int i = 1; i <= N; i++) {
        int x = i;
        int digit_sum = 0;
        while (0 < x) {
            digit_sum += x % 10;
            x /= 10;
        }

        if (A <= digit_sum && digit_sum <= B) {
            answer += i;
        }
    }

    cout << answer << endl;
}

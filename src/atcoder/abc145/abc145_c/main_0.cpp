#include <bits/stdc++.h>

using namespace std;

using i32 = int32_t;
using i64 = int64_t;
using u32 = uint32_t;
using u64 = uint64_t;
using f32 = float;
using f64 = double;

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
    i64 N;
    cin >> N;

    vector<pair<i64, i64>> points(N);
    for (i64 i = 0; i < N; i++) {
        pair<i64, i64> point;
        cin >> point.first >> point.second;
        points[i] = point;
    }

    vector<i64> orders(N);
    for (i64 i = 0; i < N; i++) {
        orders[i] = i;
    }

    f64 sum = 0;
    do {
        for (i64 i = 0; i < N - 1; i++) {
            f64 dx = points[orders[i]].first - points[orders[i + 1]].first, dy = points[orders[i]].second - points[orders[i + 1]].second;
            sum += sqrt(dx * dx + dy * dy);
        }
    } while (next_permutation(orders.begin(), orders.end()));

    for (i64 i = 1; i <= N; i++) {
        sum /= i;
    }

    printf("%.10lf\n", sum);
}

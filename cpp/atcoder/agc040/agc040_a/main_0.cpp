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
    string S;
    cin >> S;

    vector<i64> llt(S.size() + 1, 0), rgt(S.size() + 1, 0);
    llt[0] = 0;
    rgt[S.size()] = 0;
    for (i64 i = 0, j = S.size() - 1; i < i64(S.size()); i++, j--) {
        if (S[i] == '<') {
            llt[i + 1] = llt[i] + 1;
        } else if (S[i] == '>') {
            llt[i + 1] = 0;
        }
        if (S[j] == '<') {
            rgt[j] = 0;
        } else if (S[j] == '>') {
            rgt[j] = rgt[j + 1] + 1;
        }
    }

    i64 sum = 0;
    for (i64 i = 0; i <= i64(S.size()); i++) {
        sum += max(llt[i], rgt[i]);
    }

    cout << sum << endl;
}

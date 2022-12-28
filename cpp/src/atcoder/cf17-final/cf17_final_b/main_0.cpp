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

/**
 * 2文字以上の回文が含まれないためには、2文字または3文字の回文が含まれなければよい。
 *     偶数文字の回分の中心には2文字の回文が出現する。
 *     奇数文字の回文の中心には3文字の回文が出現する。
 * つまり、同じ文字が2文字以上離れていればよい。
 * 各文字の出現回数の差が1以下であれば可能である。
 */
bool solve(string S) {
    i32 a = 0, b = 0, c = 0;
    for (const auto &ch : S) {
        if (ch == 'a') {
            a++;
        } else if (ch == 'b') {
            b++;
        } else if (ch == 'c') {
            c++;
        }
    }
    return abs(a - b) < 2 && abs(b - c) < 2 && abs(c - a) < 2;
}

void run() {
    string S;
    cin >> S;
    cout << (solve(S) ? "YES" : "NO") << endl;
}

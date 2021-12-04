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

i32 solve(string s) {
    /**
     * 文字列sに含まれる全ての文字種について、
     * 複数出現しうるその文字と両端及び文字同士の間に含まれる文字数の最大値の最小値が、その文字で文字列の文字を一種類にできるまでの最小回数である
     * だから、問題の答えはそれを文字列に含まれる全ての文字種に関して最小値を取ることである
     */
    i32 answer = s.size();
    for (char c = 'a'; c <= 'z'; c++) {
        i32 current_max = 0;
        i32 last_appearance = -1;
        for (i32 i = 0; i < (i32) s.size(); i++) {
            if (s[i] == c) {
                choose_max(current_max, i - last_appearance - 1);
                last_appearance = i;
            }
        }
        if (last_appearance == -1) {
            continue;
        } else {
            choose_max(current_max, (i32) s.size() - last_appearance - 1);
        }
        choose_min(answer, current_max);
    }
    return answer;
}

void run() {
    string s;
    cin >> s;
    cout << solve(s) << endl;
}

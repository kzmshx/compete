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
    string w;
    cin >> w;

    int alphabet_counts[26] = {0};
    for (const char &c : w) {
        alphabet_counts[c - 'a']++;
    }

    int is_beautiful_word = true;
    for (const int &alphabet_count : alphabet_counts) {
        if (alphabet_count % 2 == 1) {
            is_beautiful_word = false;
        }
    }

    cout << (is_beautiful_word ? "Yes" : "No") << endl;
}

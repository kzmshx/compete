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
    size_t N;
    cin >> N;

    string W;
    cin >> W;

    unordered_map<string, bool> word_appearances = {{W, true}};
    char last_char = W[W.size() - 1];

    for (size_t i = 1; i < N; i++) {
        cin >> W;
        if (word_appearances.find(W) != word_appearances.end() || last_char != W[0]) {
            cout << "No" << endl;
            return;
        }
        word_appearances.emplace(W, true);
        last_char = W[W.size() - 1];
    }

    cout << "Yes" << endl;
}

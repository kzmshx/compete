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
    size_t N, M;
    cin >> N >> M;

    vector<pair<bool, int>> record(N, make_pair(false, 0));
    for (size_t i = 0; i < M; i++) {
        size_t p;
        string S;
        cin >> p >> S;

        if (record[p - 1].first) {
            continue;
        }

        if (S == "AC") {
            record[p - 1].first = true;
        } else if (S == "WA") {
            record[p - 1].second++;
        }
    }

    int ac = 0, wa = 0;
    for (const auto &r : record) {
        ac += r.first ? 1 : 0;
        wa += r.first ? r.second : 0;
    }

    cout << ac << " " << wa << endl;
}

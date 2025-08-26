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
    /**
     * 回転行列
     * https://w3e.kanazawa-it.ac.jp/math/category/gyouretu/henkan-tex.cgi?target=/math/category/gyouretu/kaitengyouretu.html
     */
    int x1, y1, x2, y2;
    cin >> x1 >> y1 >> x2 >> y2;

    int dx = x2 - x1, dy = y2 - y1;
    int x = x2, y = y2;

    for (int i = 0; i < 2; i++) {
        int _dx = -dy;
        int _dy = dx;
        dx = _dx;
        dy = _dy;

        x = x + dx;
        y = y + dy;

        printf("%d %d", x, y);
        if (i == 0) {
            printf(" ");
        } else {
            printf("\n");
        }
    }
}

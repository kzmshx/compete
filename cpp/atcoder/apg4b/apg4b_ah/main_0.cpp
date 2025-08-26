#include <bits/stdc++.h>

using namespace std;

template<typename T>
T square(T x) {
    return x * x;
}

template<typename T>
struct Point {
    T x;
    T y;

    void print() {
        cout << "(" << x << ", " << y << ")" << endl;
    }
};

template<int INDEX_1, int INDEX_2>
void tuple_swap(tuple<int, int, int> &x) {
    swap(get<INDEX_1>(x), get<INDEX_2>(x));
}

template<int MOD>
struct Modulo {
    int value;

    Modulo() {}

    Modulo(int x) {
        value = x % MOD;
    }

    Modulo operator+(const Modulo<MOD> &modulo) {
        Modulo<MOD> result;
        result.value = (value + modulo.value) % MOD;
        return result;
    }

    Modulo operator-(const Modulo<MOD> &modulo) {
        Modulo<MOD> result;
        result.value = (value - modulo.value + MOD) % MOD;
        return result;
    }
};

template<typename T>
bool chmax(T &a, const T &b) {
    if (a < b) {
        a = b;
        return true;
    }
    return false;
}

template<typename T>
bool chmin(T &a, const T &b) {
    if (a > b) {
        a = b;
        return true;
    }
    return false;
}

int f(int n) {
    return n * n - 8 * n + 3;
}

int main() {
    {
        int a = 3;
        double b = 1.2;

        cout << square(a) << endl;
        cout << square<int>(a) << endl;
        cout << square(b) << endl;
        cout << square<double>(b) << endl;
    }

    {
        Point<int> p1 = {0, 1};
        p1.print();

        Point<double> p2 = {2.3, 4.5};
        p2.print();
    }

    {
        tuple<int, int, int> x = make_tuple(1, 2, 3);
        tuple_swap<0, 2>(x);
        cout << get<0>(x) << ", " << get<1>(x) << ", " << get<2>(x) << endl;

        tuple_swap<0, 1>(x);
        cout << get<0>(x) << ", " << get<1>(x) << ", " << get<2>(x) << endl;
    }

    {
        Modulo<10> a(7), b(5), c(4);
        auto d = a + b;
        cout << d.value << endl;

        auto e = d - c;
        cout << e.value << endl;
    }

    {
        int ans_min = 1000000000;
        int ans_max = 0;
        for (int i = 1; i <= 10; ++i) {
            chmin(ans_min, f(i));
            chmax(ans_max, f(i));
        }
        cout << ans_min << " " << ans_max << endl;
    }
}

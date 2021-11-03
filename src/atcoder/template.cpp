#include <bits/stdc++.h>

using namespace std;

template<typename T> bool chmin(T &min, const T &value);
template<typename T> bool chmax(T &max, const T &value);

int main() {
}

template<typename T> bool chmin(T &min, const T &value) {
    if (min > value) {
        min = value;
        return true;
    }
    return false;
}

template<typename T> bool chmax(T &max, const T &value) {
    if (max < value) {
        max = value;
        return true;
    }
    return false;
}

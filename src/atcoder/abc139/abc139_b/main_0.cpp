#include <bits/stdc++.h>

using namespace std;

int main() {
    /**
     * 計算で求められる
     * 余りを見逃さないように注意
     */
    int A, B;
    cin >> A >> B;
    cout << (B - 1) / (A - 1) + (int) ((B - 1) % (A - 1) > 0) << endl;
}

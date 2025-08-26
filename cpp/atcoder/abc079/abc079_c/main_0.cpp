#include <bits/stdc++.h>

using namespace std;

int main() {
    vector<int> numbers(4);
    for (size_t i = 0; i < 4; ++i) {
        char c;
        cin >> c;
        numbers[i] = c - '0';
    }

    for (int i = 0; i < 1 << numbers.size() - 1; ++i) {
        int sum = numbers[0];
        for (size_t j = 0; j < numbers.size() - 1; ++j) {
            sum += ((i & 1 << j) != 0 ? 1 : -1) * numbers[j + 1];
        }
        if (sum == 7) {
            cout << numbers[0];
            for (size_t j = 0; j < numbers.size() - 1; ++j) {
                cout << ((i & 1 << j) != 0 ? '+' : '-') << numbers[j + 1];
            }
            cout << "=7" << endl;
            return 0;
        }
    }
}

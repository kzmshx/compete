#include <bits/stdc++.h>

using namespace std;

long long sum_of_patterns(vector<long long> &numbers, size_t from, vector<long long> &memo) {
    if (from == numbers.size()) {
        return 0;
    }
    if (memo[from] != -1) {
        return memo[from];
    }

    long long sum = 0;
    for (size_t i = from; i < numbers.size(); ++i) {
        auto current = numbers[from];
        for (size_t j = from + 1; j <= i; ++j) {
            current = 10 * current + numbers[j];
        }

        if (i < numbers.size() - 1) {
            sum += (long long) pow(2, numbers.size() - i - 2) * current + sum_of_patterns(numbers, i + 1, memo);
        } else {
            sum += current;
        }
    }

    memo[from] = sum;

    return sum;
}

int main() {
    string numbersString;
    cin >> numbersString;

    vector<long long> numbers(numbersString.size());
    for (size_t i = 0; i < numbersString.size(); ++i) {
        numbers[i] = numbersString[i] - '0';
    }

    vector<long long> memo(numbers.size(), -1);

    cout << sum_of_patterns(numbers, 0, memo) << endl;
}

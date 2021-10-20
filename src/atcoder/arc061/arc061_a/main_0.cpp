#include <bits/stdc++.h>

using namespace std;

long long sum_of_patterns(vector<long long> &numbers, size_t from) {
    if (from == numbers.size()) {
        return 0;
    }

    long long sum = 0;
    for (size_t i = from; i < numbers.size(); ++i) {
        auto current = numbers[from];
        for (size_t j = from + 1; j <= i; ++j) {
            current = 10 * current + numbers[j];
        }

        long long cur_sum = (long long) pow(2, (double) max(0, (int) (numbers.size() - i - 2))) * current + sum_of_patterns(numbers, i + 1);
        sum += cur_sum;
    }

    return sum;
}

int main() {
    string numbersString;
    cin >> numbersString;

    vector<long long> numbers(numbersString.size());
    for (size_t i = 0; i < numbersString.size(); ++i) {
        numbers[i] = numbersString[i] - '0';
    }

    cout << sum_of_patterns(numbers, 0) << endl;
}

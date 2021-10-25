#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

int main() {
    {
        vector<int> a = {3, 1, 5, 6, 7, 2, 4};

        auto iter_1 = a.begin();
        iter_1 = iter_1 + 2;
        auto iter_2 = iter_1 + 4;

        cout << *iter_1 << endl;
        cout << *iter_2 << endl;
    }

    {
        vector<string> a = {"kz", "to", "yk"};
        for (auto iter = a.begin(); iter != a.end(); ++iter) {
            cout << *iter << endl;
        }
    }

    {
        vector<int> vec = {1, 2, 3, 4, 5};

        auto iter_1 = vec.begin();
        cout << *iter_1 << endl;

        advance(iter_1, 3);
        cout << *iter_1 << endl;
    }

    {
        vector<int> vec = {1, 2, 3, 4, 5};
        auto iter_1 = vec.begin();
        auto iter_2 = iter_1 + 2;

        cout << distance(iter_1, iter_2) << endl;
        cout << distance(iter_2, iter_1) << endl;
    }

    {
        vector<int> vec = {1, 2, 3, 4, 5};

        auto iter_1 = vec.begin();
        cout << *iter_1 << " = 1" << endl;

        auto iter_2 = next(iter_1);
        cout << *iter_2 << " = 2" << endl;

        auto iter_3 = next(iter_2, 2);
        cout << *iter_3 << " = 4" << endl;

        auto iter_4 = prev(iter_3);
        cout << *iter_4 << " = 3" << endl;

        auto iter_5 = prev(iter_4, 2);
        cout << *iter_5 << " = 1" << endl;

        // runtime error
        //        auto iter_6 = prev(iter_5, 4);
        //        cout << *iter_6 << " = 5" << endl;
    }

    {
        vector<int> vec = {3, 1, 5, 6, 77, 91, 13, 45, 362};
        sort(vec.begin(), vec.end());
        for (auto i = vec.begin(); i != vec.end(); ++i) {
            cout << *i;
            if (i < vec.end() - 1) {
                cout << " ";
            }
        }
        cout << endl;
    }

    {
        vector<int> a = {1, 3, 4, 5, 9, 10};

        auto iter = find_if(a.begin(), a.end(), [](int x) { return (x % 2 == 0); });
        if (iter == a.end()) {
            cout << "not found" << endl;
        } else {
            cout << *iter << endl;
        }
    }

    {
        vector<int> vec = {53, 13, 62, 91, 45, 33, 12, 1, 97, 76, 48};
        sort(vec.begin(), vec.end());

        auto iter = lower_bound(vec.begin(), vec.end(), 5);
        if (iter == vec.end()) {
            cout << "not found" << endl;
        } else {
            cout << *iter << endl;
        }

        iter = lower_bound(vec.begin(), vec.end(), 34);
        if (iter == vec.end()) {
            cout << "not found" << endl;
        } else {
            cout << *iter << endl;
        }
    }
}

#include <bits/stdc++.h>

using namespace std;

int main() {
  int p, price, n;
  string text;
  cin >> p;

  switch (p) {
  case 1:
    cin >> price >> n;
    cout << n * price << endl;
    break;
  case 2:
    cin >> text >> price >> n;
    cout << text << "!" << endl;
    cout << n * price << endl;
    break;
  }
}

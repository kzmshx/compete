#include <bits/stdc++.h>

using namespace std;

int main() {
  string s;
  cin >> s;
  cout << 1 + count(s.begin(), s.end(), '+') - count(s.begin(), s.end(), '-') << endl;
}

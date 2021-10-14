#include <iostream>

using namespace std;

int main() {
  int n, a;
  cin >> n >> a;

  for (int i = 1; i <= n; i++) {
    char op;
    int b;
    cin >> op >> b;

    switch (op) {
    case '+':
      a += b;
      break;
    case '-':
      a -= b;
      break;
    case '*':
      a *= b;
      break;
    case '/':
      if (b == 0) {
        cout << "error" << endl;
        return 0;
      }
      a /= b;
      break;
    }

    cout << i << ":" << a << endl;
  }
}

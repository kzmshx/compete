#include <bits/stdc++.h>

using namespace std;

namespace A {
    void f() {
        cout << "namespace A" << endl;
    }
} // namespace A

namespace B {
    void f() {
        cout << "namespace B" << endl;
    }
} // namespace B

namespace C {
    namespace D {
        void f() {
            cout << "namespace C::D" << endl;
        }
    } // namespace D

    void f() {
        cout << "namespace C" << endl;
    }
} // namespace C

int main() {
    A::f();
    B::f();
    C::f();
    C::D::f();
}

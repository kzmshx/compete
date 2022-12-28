#include <iostream>
#include <memory>
#include <vector>

using namespace std;

struct A {
    int data;

    void print() {
        cout << data << endl;
    }
};

void f(int &ref) {
    ref = 2;
}

void g(int *ptr) {
    *ptr = 2;
}

int main() {
    {
        cout << "====================" << endl;
        cout << "sizeof types" << endl;

        cout << sizeof(int8_t) << endl;
        cout << sizeof(int16_t) << endl;
        cout << sizeof(int32_t) << endl;
        cout << sizeof(int64_t) << endl;
    }

    {
        cout << "====================" << endl;
        cout << "try using pointer" << endl;

        int x = 1;
        cout << x << endl;

        int *p;
        p = &x;
        *p = 2;
        cout << x << endl;

        int y;
        y = *p;
        cout << y << endl;
    }

    {
        cout << "====================" << endl;
        cout << "access members via pointer" << endl;

        A a = A{1};

        A *p = &a;
        p->print();

        p->data = 2;
        p->print();
    }

    {
        cout << "====================" << endl;
        cout << "value of pointer" << endl;

        uint8_t x = 1;
        uint8_t *p;
        p = &x;
        cout << &p << endl;
    }

    {
        cout << "====================" << endl;
        cout << "compare addresses" << endl;

        int x = 1;

        int *p = &x;
        int *q = &x;

        if (p == q) {
            cout << "p == q" << endl;
        } else {
            cout << "p != q" << endl;
        }
    }

    {
        cout << "====================" << endl;
        cout << "allocate and free memory" << endl;

        uint32_t *p;

        p = new uint32_t;
        *p = 1;
        cout << *p << endl;

        delete p;
    }

    {
        cout << "====================" << endl;
        cout << "allocate and free a sequence of memory areas" << endl;

        // uint32_t 型10個分、ヒープ領域からメモリを確保する
        uint32_t *p;
        p = new uint32_t[10];

        // ポインタを介して p を使うために、p のアドレスをコピーする
        uint32_t *tmp = p;
        for (uint32_t i = 0; i < 10; ++i) {
            // i 番目に i を書き込む
            *tmp = i;
            // 次の要素へ
            ++tmp;
        }

        tmp = p;
        for (int i = 0; i < 10; ++i) {
            cout << *tmp << endl;
            tmp++;
        }

        delete[] p;
    }

    {
        cout << "====================" << endl;
        cout << "nullptr test" << endl;

        uint8_t x = 1;
        uint8_t *p = nullptr;
        p = &x;
        *p = 2;
        p = nullptr;
        cout << (int) x << endl;

        if (p) {
            cout << "not nullptr" << endl;
        } else {
            cout << "nullptr" << endl;
        }
    }

    {
        cout << "====================" << endl;
        cout << "calculate addresses" << endl;

        int *p = new int[10];
        int *q = nullptr;

        for (int i = 0; i < 10; ++i) {
            q = p + i;
            *q = i;
        }

        q = p;
        for (int i = 0; i < 10; ++i) {
            cout << *q << endl;
            ++q;
        }

        delete[] p;
    }

    {
        cout << "====================" << endl;
        cout << "use pointer like reference" << endl;

        int x = 1;
        cout << x << endl;
        f(x);
        cout << x << endl;

        int y = 1;
        cout << y << endl;
        g(&y);
        cout << y << endl;
    }

    {
        cout << "====================" << endl;
        cout << "use pointer like iterator" << endl;

        vector<int> a = {1, 2, 3};
        for (auto iter = a.begin(); iter != a.end(); ++iter) {
            cout << *iter << endl;
        }

        vector<int> b = {1, 2, 3};
        int *begin_address = b.data();
        for (int *ptr = begin_address; ptr < begin_address + 3; ptr = ptr + 1) {
            cout << *ptr << endl;
        }
    }

    {
        cout << "====================" << endl;
        cout << "void pointer" << endl;

        int x = 1234;
        string y = "hello";
        double z = 1.234;

        cout << "任意のポインターを扱える" << endl;
        vector<void *> pointers = {&x, &y, &z};

        cout << "void pointer へのアクセスには元のポインタ型にキャストしなければならない" << endl;
        int *pointer_of_x = (int *) pointers[0];
        string *pointer_of_y = (string *) pointers[1];
        double *pointer_of_z = (double *) pointers[2];

        cout << *pointer_of_x << endl;
        cout << *pointer_of_y << endl;
        cout << *pointer_of_z << endl;
    }

    {
        cout << "====================" << endl;
        cout << "unique_ptr" << endl;

        // #include <memory>
        unique_ptr<int> p1 = make_unique<int>(123);
        *p1 += 1;
        cout << *p1 << endl;

        unique_ptr<int> p2;
        p2 = move(p1); // p1 のメモリ所有権を p2 へ移動
        *p2 += 1;
        cout << *p2 << endl;
    }

    {
        cout << "====================" << endl;
        cout << "shared_ptr" << endl;

        shared_ptr<int> p3;
        {
            shared_ptr<int> p1 = make_shared<int>(123);
            {
                shared_ptr<int> p2 = p1;
                *p2 += 1;
                p3 = p2;
            }
            *p1 += 1;
        }
        *p3 += 1;
        cout << *p3 << endl;
    }
}

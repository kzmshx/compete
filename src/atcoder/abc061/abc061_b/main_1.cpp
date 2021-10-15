#define _GLIBCXX_DEBUG
#include <iostream>
#include <vector>

using namespace std;

int main() {
  int N, M;
  cin >> N >> M;
  vector<int> numOfRoads(N, 0);
  for (int i = 0; i < M; i++) {
    int a, b;
    cin >> a >> b;
    numOfRoads[a - 1]++;
    numOfRoads[b - 1]++;
  }
  for (int i = 0; i < N; i++) {
    cout << numOfRoads[i] << endl;
  }
}

#include <iostream>

using namespace std;

int main(int argc, char *argv[]) {
  if (argc < 3) {
    cout << "need 2 arguments." << endl;
    return 1;
  }

  auto x = atoi(argv[1]);
  auto y = atoi(argv[2]);
  auto sum = x + y;
  cout << "sum:" << sum << endl;

  return 0;
}

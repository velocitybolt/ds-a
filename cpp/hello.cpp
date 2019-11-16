#include <iostream>
#include <vector>

using namespace std;

int calculate_pairs(vector<int> vec) {
  int result{0};

  for (int i = 0; i < vec.size(); i++) {
    for (int j = i + 1; j < vec.size(); j++) {
      cout << vec[i] << vec[j] << '\n';
      result += vec[i] * vec[j];
    }
  }

  return result;
}

int main() {
  cout << calculate_pairs(vector<int>{1, 2, 3}) << '\n';
  vector<int> vec{3324, 34, 23, 42, 34, 234, 4};
}
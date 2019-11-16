#include <iostream>
#include <vector>

using namespace std;

int vector_shop(vector<int> vec) {
  cout << "Please Select an Option" << '\n';
  char input;
  cin >> input;
  if (tolower(input) == 'a') {
    cout << "Please enter an Integer to add to the vector" << '\n';
    int aInput;
    cin >> aInput;
    // Dont care anymore...
  }
}

int main() { cout << vector_shop(vector<int>{1, 2, 3, 4, 5}) << '\n'; }
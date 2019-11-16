#include <iostream>
#include <vector>

using namespace std;

void print(int);
void print(double);
void print(string);
void print(string, string);
void print(vector<string>);

void print(int num) { cout << "Printing int: " << num << '\n'; }

void print(double num) { cout << "Printing double: " << num << '\n'; }

void print(string s) { cout << "Printing string: " << s << '\n'; }

void print(string s, string t) {
  cout << "Printing string: " << s << '\n' << t << '\n';
}

void print(vector<string> vec) { cout << 1; }

int main() {}

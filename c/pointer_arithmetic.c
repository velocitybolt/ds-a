#include <stdio.h>

/*
 * This section contains stuff about pointer arithmetic
 */

char newTestFunction(char* input) {
  printf("The element within strPtr is currently %c\n", *input);
  input++;
  printf("The element within strPtr is currently %c\n", *input);
  input += 2;
  printf("The element within strPtr is currently %c\n", *input);
}

int testFunction(int* input) {
  printf("The element within intPtr is currently %d\n", *input);
  input++;
  printf("The element within intPtr is currently %d\n", *input);
  input += 2;
  printf("The element within intPtr is currently %d\n", *input);
}

int main() {
  int testArr[4] = {1, 3, 4, 5};
  int* testPtr = testArr;
  char* strPtr = "Hello World!";
  testFunction(testPtr);
  newTestFunction(strPtr);
}

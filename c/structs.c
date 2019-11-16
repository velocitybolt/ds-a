#include <stdbool.h>
#include <stdio.h>

struct someData {
  int value;
  char ch;
  bool is_Valid;
};

int main() {
  struct someData some_Data;
  some_Data.value = 2;
  some_Data.ch = 'a';
  some_Data.is_Valid = true;

  struct someData *struct_ptr = &some_Data;

  struct_ptr->value = 5;
  // "->" is dereferencing for struct related values ...
  printf("The value is %d\n", struct_ptr->value);
  // add one to the number stored inside the struct ...
  struct_ptr->value += 1;
  printf("The value is %d\n", struct_ptr->value);
}
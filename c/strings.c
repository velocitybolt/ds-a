#include <stdio.h>
#include <string.h>

size_t my_Strlen(const char* str) {
  /*
  Set variable to count length
  Set char variable pointer to point to the first element of the string
  Increment the char pointer to move to next character
  Increment length pointer as we move through characters
  */
  if (str == NULL) {
    return 0;
  }
  int length = 0;
  const char* ch = str;
  while (*ch != '\0') {
    length++;
    ch++;
  }
  return length;
}

/*
strchr function to check whether a character is the string or not
*/
char* my_Strchr(const char* str, int c) {
  if (str == NULL) {
    return NULL;
  }
  while (*str != '\0') {
    if (*str == c) return (char*)str;
    printf("This is the char %c\n", *str);
    str++;
  }
  return NULL;
}
/*
strstr is a function that checks for substrings within the string
*/
char* my_Strstr(const char* haystack, const char* needle) {
  if (haystack == NULL || needle == NULL) {
    return NULL;
  }
  while (*haystack != '\0') {
    if (*needle == *haystack) {
      const char* h = haystack;
      const char* n = needle;
      while (*n != '\0' && *h == *n) {
        h++;
        n++;
      }
      if (*n == '\0') {
        return (char*)haystack;
      }
    }
    haystack++;
  }
  return NULL;
}
/*
strconcat appends to the end of a string
*/
char* my_Strconcat(const char* dest, const char* src) {
  if (dest == NULL || src == NULL) {
    return dest;
  }
  char* d = dest;
  while (dest != '\0') {
    d++;  // get to the end of the string so you can append starting then ...
  }
  while (*src != '\0') {
    *d = *src;  // point the first char to be appended to the back of dest...
    d++;
    src++;
  }
  *d = '\0';
  return dest;
}

int main() {
  char* helloStr = "Hello";
  char* strPtr = "Hello World!";
  printf("The length of the string is %ld\n", my_Strlen(strPtr));
  printf("Answer %s\n", my_Strchr(strPtr, 111));
  printf("The answer %s\n", my_Strstr(helloStr, strPtr));
}
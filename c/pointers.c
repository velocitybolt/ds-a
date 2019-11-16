#include <stdio.h>
int main()
{
    int circumference = 49;
    int *cPtr = &circumference;

    printf("The value at stored at c is %d\n", *cPtr);
    printf("This mem addr of c is %u\n", cPtr);
}

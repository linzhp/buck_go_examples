#include <stdio.h>
#include "_cgo_export.h"

void printSomethingFromGo(int a, int b)
{
    extern GoInt GoFunction(GoInt, GoInt);
    printf("From C: %d\n", (int)GoFunction(a, b));
}
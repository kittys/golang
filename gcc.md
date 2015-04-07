### using gcc to compile c program
sample c program. write and save as hello.c
```c
#include <stdio.h>
 
int main()
{
  printf("Hello world\n");
}
```
* [sudo apt-get install libc6-dev]()
* [sudo apt-get install g++](http://pages.cs.wisc.edu/~beechung/ref/gcc-intro.html)
* g++ hello.c -o hell0
* ./hello
* note <stdio.h> search standard lib, "stdio.h" search local file

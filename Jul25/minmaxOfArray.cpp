/******************************************************************************

Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
#include <iostream>
using namespace std;

int main()
{
    int a[10] = {10,4,5,12,29,45,6,19,1105,49};
    int l = sizeof(a) / sizeof(a[0]);
    int min=a[0] , max=a[0];
    for(int i=0 ; i<l ; i++)
    {
        if(a[i]>max)
            max=a[i];
        if(a[i]<min)
            min=a[i];
    }
    cout << "Smallest element is : " << min << endl;
    cout << "Largest element is : " << max;
    
    
}
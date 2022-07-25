/******************************************************************************

Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
#include <iostream>
using namespace std;

bool isPalin(string str)
{
    int l = str.length();
    string newstr;
    for(int i=l-1 ; i>=0 ; i--)
    {
        newstr = newstr+str[i];
    }
    if(str==newstr)
        return true;
    else
        return false;    
    
}

int main()
{
    string s = "aabaa";
    string x = "abababc";
    cout << isPalin(s);
    cout<< "   ";
    cout << isPalin(x);
}
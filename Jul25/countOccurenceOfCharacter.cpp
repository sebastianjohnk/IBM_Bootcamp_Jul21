/******************************************************************************

Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
#include <iostream>
using namespace std;

int countTheLetter(string str,char ch)
{
    int l = str.length();
    int count = 0;
    for(int i=0 ; i<l ; i++)
    {
        if(str[i]==ch)
        {
            count++;
        }
    }
    return count;
}

int main()
{
    string s = "ababcababc";
    cout << countTheLetter(s,'c');
    
}
/******************************************************************************

Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
#include <iostream>
using namespace std;

string remRepLet(string str)
{
    int l = str.length();
    string newstr;
    for(int i=0 ; i<l ; i++)
    {
        int len = newstr.length();
        int flag=0;
        for(int j=0 ; j<len ; j++)
        {
            if(str[i]==newstr[j])
                flag=1;
        }
        if(flag==0)
            newstr = newstr+str[i];
        else
            continue;
    }
    return newstr;
}

int main()
{
    string s = "abcdabcefgdhijkabdcfgihf";
    cout << "Original string : " << s << endl;
    cout << "New string : " << remRepLet(s);
    
}
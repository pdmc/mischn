#include <iostream>
#include "base64.hpp"

using namespace std;
int main()
{
    unsigned char str[] = "I'm a test string.";
    string normal,encoded;
    int i,len = sizeof(str)-1;
    Base64 *base = new Base64();
    encoded = base->Encode(str,len);
    cout << "base64 encode : " << encoded << endl;
    len = encoded.length();
    const char * str2 = encoded.c_str();
    normal = base->Decode(str2,len);
    cout << "base64 decode : " << normal <<endl;
    return 0;
}

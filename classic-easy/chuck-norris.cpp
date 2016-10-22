#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
int main()
{
    string MESSAGE;
    getline(cin, MESSAGE);
    unsigned char chars [MESSAGE.length()];
    string answer = "";
    bool ones = false;
    bool first = true;
    
    for(int i=0; i<MESSAGE.length(); i++){
        chars[i] = MESSAGE.at(i);
        char c [8];
        for (int j = 1; j<8; j++){
            c[j] = chars[i]>>(7-j)&1;
            if(c[j] == 1){
                if(ones){
                    cout << 0;
                }else{
                    ones = true;
                    if (first){
                        first = false;
                        cout << 0 << ' ' << 0;
                    }else{
                        cout << ' ' << 0 << ' ' << 0;
                    }
                } 
            }else{
                if(!ones){
                    if(first){
                        first = false;
                        cout << 0 << 0 << ' ' << 0;
                    }else{
                        cout << 0;
                    }
                }else{
                    ones = false;
                    cout << ' ' << 0 << 0 << ' ' << 0;
                }
            }
        }
    }
    cout << endl;
    
    // Write an action using cout. DON'T FORGET THE "<< endl"
    // To debug: cerr << "Debug messages..." << endl;

    //cout << "answer" << endl;
}
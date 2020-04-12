#include<iostream>
using namespace std;

class Singleton{
private:
    static Singleton* instance;
    Singleton(){}
public:
    static Singleton* getInstance(){
        if(instance == NULL){
            instance = new Singleton();
        }
        return instance;
    }
};

Singleton* Singleton::instance = NULL;


int main(){
    Singleton* sig1 =  Singleton::getInstance();
    Singleton* sig2 =  Singleton::getInstance();
    Singleton* sig3 =  sig1->getInstance();
    
    // Printed value: 003F5320 003F5320 003F5320
    cout << sig1 << " " << sig2 << " " << sig3 << endl;
    
    return 0;
}
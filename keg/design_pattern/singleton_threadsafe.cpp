#include<iostream>
#include<mutex>
using namespace std;

class Singleton{
private:
    static Singleton* instance;
    static std::mutex mtx;
public:
    Singleton(){}

    static Singleton* getInstance(){
        // Double checked locking pattern
        if(instance == NULL){
            std::lock_guard<std::mutex> loker(mtx);
            if(instance == NULL){
                instance = new Singleton();
            }
        }
        return instance;
    }
};

Singleton* Singleton::instance = NULL;
std::mutex Singleton::mtx;

int main(){
    Singleton* sig1 =  Singleton::getInstance();
    Singleton* sig2 =  Singleton::getInstance();
    Singleton* sig3 =  sig1->getInstance();
    
    // Printed value: 003F5320 003F5320 003F5320
    cout << sig1 << " " << sig2 << " " << sig3 << endl;
    
    return 0;
}
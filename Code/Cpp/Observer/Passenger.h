#include <iostream>
using namespace std;

class Train;

class Passenger {
    protected: Train* train;
    public: void enter(Train* train){
        this->train = train;
    }
    public: virtual void arrive(string station) = 0;
};
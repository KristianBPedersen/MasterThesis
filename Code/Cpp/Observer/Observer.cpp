#include <string>
#include <iostream>
#include <vector>
#include <algorithm>

#include "Train.h"

using namespace std;

class CountingPassenger : public Passenger {
    private: int count;

    public: CountingPassenger(int count) {
        this->count = count;
    }
    
    public: void arrive(string station) {
        count--;
        if (count == 0) train->discharge(this);
    } 
};

class CheckingPassenger : public Passenger {
    private: string station;

    public: CheckingPassenger(string station) {
        this->station = station;
    }
    
    public: void arrive(string station) {
        if (this->station == station) train->discharge(this);
    } 
};

int main(){
        string stations[4] = {"Station A", "Station B", "Station C", "Station D"};

        Train* train = new Train();
        train->station = stations[0];

        for (int i = 1; i < 6; i++) {
            train->board(new CountingPassenger(i));
        }

        for (int i = 0; i < 4; i++) {
            train->board(new CheckingPassenger(stations[i]));
        }

        int i = 0;
        while(train->passengers.size() != 0) {
            i = (i+1)%4;
            train->arrive(stations[i]);           
        }

        printf("Final station: ");
        printf(train->station.c_str());
}
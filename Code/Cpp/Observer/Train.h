#include <string>
#include <iostream>
#include <vector>
#include <algorithm>

#include "Passenger.h"

class Train {
    public: vector<Passenger*> passengers = {};
    public: string station;

    public: void board(Passenger* passenger) {
        passenger->enter(this);   
        passengers.push_back(passenger); 
    }

    public: void discharge(Passenger* passenger) {
        passengers.erase(find(passengers.begin(), passengers.end(), passenger));
    }

    public: void arrive(string station) {
        this->station = station;
        announceStation();
    }

    public: void announceStation(){
        for (int i = passengers.size() - 1; i >= 0; i--) {
            passengers[i]->arrive(station);
        }
    }
};
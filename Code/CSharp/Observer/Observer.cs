using System;
using System.Collections.Generic;

class Train {

    public List<Passenger> passengers = new List<Passenger>();
    public string station;

    public void board(Passenger passenger) {
        passenger.enter(this);
        passengers.Add(passenger);
    }

    public void discharge(Passenger passenger) {
        passengers.Remove(passenger);
    }

    public void arrive(string station) {
        this.station = station;
        announceStation();
    }

    public void announceStation() {
        for(int i = passengers.Count - 1; i >= 0 ; i--) {
            passengers[i].arrive(station);
        }
    }
}

abstract class Passenger {
    protected Train train;
    
    public void enter(Train train) {
        this.train = train;
    }

    public abstract void arrive(string station);
}

class CountingPassenger : Passenger {
    private int count;

    public CountingPassenger(int count) {
        this.count = count;
    }
    
    public override void arrive(string station) {
        count--;
        if (count == 0) train.discharge(this);
    }
}

class CheckingPassenger : Passenger {
    public string station;
    public CheckingPassenger(string station) {
        this.station = station;
    }
    public override void arrive(string station) {
        if (this.station == station) train.discharge(this);
    }
}

class TrainLoop {
    public static void Main() {
        string[] stations = new string[]{"Station A", "Station B", "Station C", "Station D"};
        
        Train train = new Train();
        train.station = stations[0];

        for (int i = 1; i < 6; i++) {
            train.board(new CountingPassenger(i));
        }

        for (int i = 0; i < stations.Length; i++) {
            train.board(new CheckingPassenger(stations[i]));
        }

        int j = 0;
        while (train.passengers.Count != 0) {
            j = (j+1)%stations.Length;
            train.arrive(stations[j]);
        }
        Console.WriteLine("Final Station: " + train.station);
    }
}
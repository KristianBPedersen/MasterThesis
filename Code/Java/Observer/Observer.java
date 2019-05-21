import java.util.ArrayList;
import java.util.List;

class Train {
    public List<Passenger> passengers = new ArrayList<Passenger>();
    public String station;

    public void board(Passenger passenger) {
        passenger.enter(this);
        passengers.add(passenger);
    }

    public void discharge(Passenger passenger) {
        passengers.remove(passenger);
    }

    public void arrive(String station) {
        this.station = station;
        announceStation();
    }

    public void announceStation() {
        for(int i = passengers.size() - 1; i >= 0 ; i--) {
            passengers.get(i).arrive(station);
        }
    } 
}

abstract class Passenger {
    protected Train train;
    
    public void enter(Train train) {
        this.train = train;
    }

    public abstract void arrive(String station);
}

class CountingPassenger extends Passenger {
    private int count;

    public CountingPassenger(int count) {
        this.count = count;
    }
    
    public void arrive(String station) {
        count--;
        if (count == 0) train.discharge(this);
    }
}

class CheckingPassenger extends Passenger {
    public String station;
    public CheckingPassenger(String station) {
        this.station = station;
    }
    public void arrive(String station) {
        if (this.station == station) train.discharge(this);
    }
}

class Observer {
    public static void main(String[] args) {
        String[] stations = new String[]{"Station A", "Station B", "Station C", "Station D"};
        
        Train train = new Train();
        train.station = stations[0];

        for (int i = 1; i < 6; i++) {
            train.board(new CountingPassenger(i));
        }

        for (int i = 0; i < stations.length; i++) {
            train.board(new CheckingPassenger(stations[i]));
        }

        int j = 0;
        while (train.passengers.size() != 0) {
            j = (j+1)%stations.length;
            train.arrive(stations[j]);
        }
        System.out.println("Final Station: " + train.station);
    }
}


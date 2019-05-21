class Train {
    constructor() {    
        this.passengers = [];
    }

    board(passenger) {
        passenger.enter(this);
        this.passengers.push(passenger)
    }

    discharge(passenger) {
        this.passengers.splice(this.passengers.indexOf(passenger), 1)
    }

    arrive(station) {
        this.station = station;
        this.announceStation();
    }

    announceStation() {
        for(i = this.passengers.length - 1; i >= 0; i--) {
            this.passengers[i].arrive(this.station);
        }
    }
}

class Passenger {
    constructor() {
        if (this.constructor === Passenger) {
            throw new TypeError('Abstract class "Passenger" cannot be instantiated directly.'); 
        }
    }
    enter(train) {
        this.train = train;
    }
}

class CountingPassenger extends Passenger {
    constructor(count){
        super()
        this.count = count;
    }

    arrive(station) {
        this.count--;
        if (this.count == 0) this.train.discharge(this);
    }
}

class CheckingPassenger extends Passenger {
    constructor(station) {
        super()
        this.station = station;
    }

    arrive(station) {
        if (this.station == station) this.train.discharge(this);
    }
}

stations = ["Station A", "Station B", "Station C", "Station D"];
train = new Train();
train.station = stations[0];

for (i = 1; i < 6; i++) {
    train.board(new CountingPassenger(i))
}

for (i = 0; i < 4; i++) {
    train.board(new CheckingPassenger(stations[i]))
}

j = 0
while(train.passengers.length != 0) {
    j = (j+1)%stations.length;
    train.arrive(stations[j]);
}
console.log("Final Station" + train.station);
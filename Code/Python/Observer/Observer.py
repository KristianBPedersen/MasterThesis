class Train:
    passengers = []

    def board(self, passenger):
        passenger.enter(self)
        self.passengers += [passenger]
    
    def discharge(self, passenger):
        self.passengers.remove(passenger)
    
    def arrive(self, station):
        self.station = station
        self.announceStation()

    def announceStation(self):
        for i in range(len(self.passengers)-1, -1, -1):
            self.passengers[i].arrive(self.station)

class Passenger():
    def enter(self, train):
        self.train = train

class CountingPassenger(Passenger):
    def __init__(self, count):
        self.count = count

    def arrive(self, station):
        self.count -= 1
        if self.count == 0: self.train.discharge(self)

class CheckingPassenger(Passenger):
    def __init__(self, station):
        self.station = station

    def arrive(self, station):
        if self.station == station: self.train.discharge(self)

def trainLoop():
    stations = ["Station A", "Station B", "Station C", "Station D"]

    train = Train()
    train.station = stations[0]
    
    for passenger in [CountingPassenger(i) for i in range(1,6)]:
        train.board(passenger)

    for passenger in [CheckingPassenger(station) for station in stations]:
        train.board(passenger)

    i = 0
    while len(train.passengers) != 0:
        i = (i+1) % len(stations)
        train.arrive(stations[i])
    print(f"Final station: {train.station}")

trainLoop()
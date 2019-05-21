class Train:

    def __init__(self, wagonCount):
        self.wagonCount = wagonCount
    
    def setLine(self, lineNumber):
        self.line = "T" + str(lineNumber)

    def acceptPassengers(self):
        self.passengerCount = self.wagonCount*50

    def stop(self):
        self.passengerCount -= self.wagonCount*50
    
    def checkTrainTicket(self, code):
        if (code != "train"):
            print("Error")
    
    def drive(self):
        print("Train " + self.line + " driving with " + str(self.passengerCount) + " passengers")

class Bus:

    def __init__(self, capacity):
        self.capacity = capacity
    
    def setLine(self, lineNumber):
        self.line = "B" + str(lineNumber)

    def acceptPassengers(self):
        self.passengerCount = self.capacity

    def releasePassengers(self):
        self.passengerCount = 0
    
    def checkBusTicket(self, code):
        if (code != "bus"):
            print("Error")
    
    def drive(self):
        print("Bus " + self.line + " driving with " + str(self.passengerCount) + " passengers")

class TrainReplacement(Bus, Train):
    def __init__(self, capacity):
        super().__init__(capacity)
        Train.__init__(self, 1)
        
    def setLine(self, lineNumber):
        Train.setLine(self, lineNumber)

    def stop(self):
        self.releasePassengers()

class RailroadController:
    @staticmethod
    def runLine(train, lineNumber):
        train.setLine(lineNumber)
        train.acceptPassengers()
        train.checkTrainTicket("train")
        train.drive()
        train.stop()

replacement = TrainReplacement(40)
RailroadController.runLine(replacement, 2)
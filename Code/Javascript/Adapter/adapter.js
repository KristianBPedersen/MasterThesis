function Train(wagonCount) {
    this.wagonCount = wagonCount

    this.setLine = function (lineNumber) {
        this.line = "T" + lineNumber
    }

    this.acceptPassengers = function () {
        this.passengerCount = wagonCount*50
    }

    this.stop = function () {
        this.passengerCount -= wagonCount*50
    }

    this.checkTrainTicket = function (code) {
        if (code != "train") {
            console.log("Error")
        }
    }

    this.drive = function () {
        console.log("Train " + this.line + " driving with " + this.passengerCount + " passengers")
    }
}

function Bus(capacity) {
    this.capacity = capacity

    this.setLine = function (lineNumber) {
        this.line = "B" + lineNumber
    }

    this.acceptPassengers = function () {
        this.passengerCount = capacity
    }

    this.releasePassengers = function () {
        this.passengerCount = 0
    }

    this.checkBusTicket = function (code) {
        if (code != "bus") {
            console.log("Error")
        }
    }

    this.drive = function () {
        console.log("Bus " + this.line + " driving with " + this.passengerCount + " passengers")
    }
}

function TrainReplacement(capacity) {
    Train.call(this, 1)
    var trainSetLine = this.setLine
    Bus.call(this, capacity)
    this.stop = this.releasePassengers
    this.setLine = trainSetLine

}

class RailroadController {
    static runLine(train, lineNumber) {
        train.setLine(lineNumber)
        train.acceptPassengers()
        train.checkTrainTicket("train")
        train.drive()
        train.stop()
    }
}

replacement = new TrainReplacement(40)
RailroadController.runLine(replacement, 2)
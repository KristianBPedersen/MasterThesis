public interface ITrain {
    void setLine(int lineNumber);
    void acceptPassengers();
    void stop();
    void checkTrainTicket(string code);
    void drive();
}

public class Train : ITrain {
    private int wagonCount;
    protected string line;
    protected int passengerCount;

    public Train(int wagonCount) {
        this.wagonCount = wagonCount;
    }

    public virtual void setLine(int lineNumber) {
        line = "T" + lineNumber;
    }

    public virtual void acceptPassengers(){
        passengerCount = wagonCount*50;
    }

    public virtual void stop() {
        passengerCount -= wagonCount*50;
    }

    public void checkTrainTicket(string code) {
        if (code != "train") {
            System.Console.WriteLine("Error");
        } 
    }
    public virtual void drive() {
        System.Console.WriteLine("Train " + line + " driving with " + passengerCount + " passengers");
    }
}

public class Bus {
    protected int capacity;
    protected string line;
    protected int passengerCount;

    public Bus(int capacity) {
        this.capacity = capacity;
    }

    public virtual void setLine(int lineNumber) {
        line = "B" + lineNumber;
    }

    public void acceptPassengers() {
        passengerCount = capacity;
    }

    public void releasePassengers() {
        passengerCount = 0;
    }

    public void checkBusTicket(string code) {
        if (code != "bus") {
            System.Console.WriteLine("Error");
        }
    }

    public void drive() {
        System.Console.WriteLine("Bus " + line + " driving with " + passengerCount + " passengers");
    }
}

public class RailroadController {
    public static void runLine(ITrain train, int n) {
        train.setLine(n);
        train.acceptPassengers();
        train.checkTrainTicket("train");
        train.drive();
        train.stop();
    }
}
public class AccessibleTrain : Train {
    public AccessibleTrain(int wagonCount) : base(wagonCount){
    }

    public string Line {
        get {
            return line;
        }
    }
}

public class TrainReplacement : Bus, ITrain {
    AccessibleTrain train;
    public TrainReplacement(int capacity) : base(capacity) {
        train = new AccessibleTrain(1);
    }
    public void stop() {
        this.releasePassengers();
    }

    public void checkTrainTicket(string code) {
        train.checkTrainTicket(code);
    }

    public override void setLine(int lineNumber) {
        train.setLine(lineNumber);
        line = train.Line;
    }
}

public class Tester {
    public static void Main(string[] args) {
        ITrain replacementTrain = new TrainReplacement(40);
        RailroadController.runLine(replacementTrain, 2);
    }
}
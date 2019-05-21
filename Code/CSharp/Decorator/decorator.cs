public interface Train {
    string describe();
}

public class CargoTrain : Train {
    private int id;
    public CargoTrain(int id) {
        this.id = id;
    }

    public string describe() {
        return id + " is a cargo train. ";
    }
}

public class Modification : Train {
    private Train train;

    public Modification(Train train) {
        this.train = train;
    }

    public virtual string describe() {
        return train.describe();
    }
}

public class ColorModification : Modification {
    private string color;
    public ColorModification(Train train, string color) : base(train) {
        this.color = color;
    }

    public override string describe() {
        return base.describe() + "It is painted " + color + ". ";
    }
}

public class SilencerModification : Modification {
    public SilencerModification(Train train) : base(train) {
    }

    public override string describe() {
        return base.describe() + " When it drives it makes no sound. ";
    }
}

public class Tester {
    public static void Main(string[] args) {
        ColorModification classicTrain = new ColorModification(new CargoTrain(6), "red");
        SilencerModification stealthTrain = new SilencerModification(new ColorModification(new CargoTrain(7), "black"));
        System.Console.WriteLine(classicTrain.describe());
        System.Console.WriteLine(stealthTrain.describe());
    }
}
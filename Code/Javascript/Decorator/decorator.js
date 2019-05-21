function CargoTrain(id) {
    this.id = id;
    this.describe = function () {
        return id + " is a cargo train. ";
    }
}

class Modification {
    constructor(train) {
        this.train = train
    }

    describe() {
        return this.train.describe()
    }
}

class ColorModification extends Modification {
    constructor(train, color) {
        super(train)
        this.color = color
    }
    describe() {
        return super.describe() + "It is painted " + this.color + ". "
    }
}

class SilencerModification extends Modification {
    describe() {
        return super.describe() + "When it drives it makes no sound. "
    }
}

classicTrain = new ColorModification(new CargoTrain(6),"red")
stealthTrain = new SilencerModification(new ColorModification(new CargoTrain(7),"black"))

console.log(classicTrain.describe());
console.log(stealthTrain.describe());
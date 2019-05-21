class CargoTrain:
    def __init__(self, id):
        self.id = id

    def describe(self):
        return f"{self.id} is a cargo train. "

class Modification:
    def __init__(self, train):
        self.train = train

    def describe(self):
        return self.train.describe()

class ColorModification(Modification):
    def __init__(self, train, color):
        super().__init__(train)
        self.color = color
    def describe(self):
        return super().describe() + f"It is painted {self.color}. "

class SilencerModification(Modification):
    def describe(self):
        return super().describe() + f"When it drives it makes no sound. "

classicTrain = ColorModification(CargoTrain(6), "red")
stealthTrain = SilencerModification(ColorModification(CargoTrain(7), "black"))
print(classicTrain.describe())
print(stealthTrain.describe())
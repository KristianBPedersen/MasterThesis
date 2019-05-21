from abc import ABC, abstractmethod

class Train(ABC):
    @abstractmethod
    def __init__(self, id, color):
        self.__trainID = id
        self.__color = color

    def describe(self):
        print("I am a train with id = %s and color = %s" % (self.__trainID, self.__color))

class ClassicTrain(Train):
    def __init__(self, id):
        super().__init__(id, "red")

class ModernTrain(Train):
    def __init__(self, id):
        super().__init__(id, "white")

class TrainCreator(ABC):
    @abstractmethod
    def createTrain(self):
        pass

class ModernTrainCreator(TrainCreator):
    def createTrain(self):
        return ModernTrain(1)

class ClassicTrainCreator(TrainCreator):
    def createTrain(self):
        return ClassicTrain(2)        

modernTrain = ModernTrainCreator().createTrain().describe()
classicTrain = ClassicTrainCreator().createTrain().describe()
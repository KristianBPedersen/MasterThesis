import copy

class Product:
    def clone(self):
        return copy.deepcopy(self)

class Locomotive(Product):
    def __init__(self, max_speed):
        self.max_speed = max_speed
        self.color = "red"
        self.price = 15
    
    def set_color(self, color):
        self.color = color

    def set_price(self, price):
        self.price = price

    def get_cost(self):
        return self.price

class Train(Product):
    def __init__(self, locomotive):
        self.locomotive = locomotive
        self.wagons = []
    
    def add_wagon(self, wagon):
        self.wagons.append(wagon)

    def get_cost(self):
        return self.locomotive.get_cost() + 5*len(self.wagons)

oldLocomotive = Locomotive(10)
newLocomotive = oldLocomotive.clone()

longTrain = Train(oldLocomotive)
longTrain.add_wagon("Passenger")
shortTrain = longTrain.clone()
longTrain.add_wagon("Freight")

oldLocomotive.set_price(7)

print(newLocomotive.get_cost() - oldLocomotive.get_cost())
print(shortTrain.get_cost() - longTrain.get_cost())
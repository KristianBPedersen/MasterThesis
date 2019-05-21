class Structure:
    def getChild(self, i):
        raise NotImplementedError
    def add(self, child):
        raise NotImplementedError
    def remove(self, child):
        raise NotImplementedError

class Railroad(Structure):
    def __init__(self, length):
        self.length = length
    def getMaintananceCost(self):
        return 2*self.length

class Tunnel(Structure):
    def __init__(self, length):
        self.length = length
    def getMaintananceCost(self):
        return 5*self.length

class Station(Structure):
    def getMaintananceCost(self):
        return 20

class Network:
    def __init__(self):
        self.children = []

    def getChild(self, i):
        return self.children[i]
    def add(self, child):
        self.children.append(child)
    def remove(self, child):
        self.children.remove(child)
    def getMaintananceCost(self):
        return sum(child.getMaintananceCost() for child in self.children)

tunnelNetwork = Network()
tunnelNetwork.add(Tunnel(5))
tunnelNetwork.add(Railroad(5))

stationNetwork = Network()
stationNetwork.add(Station())
stationNetwork.add(Railroad(10))

bigNetwork = Network()
bigNetwork.add(tunnelNetwork)
bigNetwork.add(stationNetwork)
print(bigNetwork.getMaintananceCost())
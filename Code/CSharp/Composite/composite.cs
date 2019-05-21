using System.Collections.Generic;
using System.Linq;

interface NetworkElement {
    NetworkElement GetChild(int i);
    void Add(NetworkElement child);
    void Remove(NetworkElement child);
    int GetMaintananceCost();
}

abstract class Structure : NetworkElement {
    public NetworkElement GetChild(int i){
        throw new System.NotImplementedException();
    }
    public void Add(NetworkElement child) {
        throw new System.NotImplementedException();
    }
    public void Remove(NetworkElement child) {
        throw new System.NotImplementedException();
    }

    public abstract int GetMaintananceCost();
}


class Network : NetworkElement {
    private List<NetworkElement> children = new List<NetworkElement>();

    public NetworkElement GetChild(int i) {
        return children[i];
    }
    public void Add(NetworkElement child) {
        children.Add(child);
    }
    public void Remove(NetworkElement child) {
        children.Remove(child);
    }
    public int GetMaintananceCost() {
        return children.Sum(child => child.GetMaintananceCost());
    }
}

class Railroad : Structure {
    private int length;
    public Railroad(int length) {
        this.length = length;
    }

    public override int GetMaintananceCost() {
        return length*2;
    }
}

class Tunnel : Structure {
    private int length;
    public Tunnel(int length) {
        this.length = length;
    }
    public override int GetMaintananceCost() {
        return this.length*5;
    }
}

class Station : Structure {
    public override int GetMaintananceCost() {
        return 20;
    }
}

class Tester {
    static void Main(string[] args) {
        NetworkElement tunnelNetwork = new Network();
        tunnelNetwork.Add(new Tunnel(5));
        tunnelNetwork.Add(new Railroad(5));

        NetworkElement stationNetwork = new Network();
        stationNetwork.Add(new Station());
        stationNetwork.Add(new Railroad(10));

        NetworkElement bigNetwork = new Network();
        bigNetwork.Add(tunnelNetwork);
        bigNetwork.Add(stationNetwork);

        System.Console.WriteLine(bigNetwork.GetMaintananceCost());
    }
}

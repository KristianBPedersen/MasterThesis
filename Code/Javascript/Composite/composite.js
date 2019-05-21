function Structure() {
    this.remove = this.add = this.getChild = function(child) {
        throw Error("Not implemented")
    }
};

function Tunnel(length) {
    Structure.call(this)
    this.length = length;
    this.getMaintananceCost = function() {
        return 5*this.length;
    }
}

function Railroad(length) {
    Structure.call(this)
    this.length = length;
    this.getMaintananceCost = function() {
        return 2*this.length;
    }
}

function Station() {
    Structure.call(this)
    this.getMaintananceCost = function() {
        return 20;
    }
}

function Network() {
    this.children = [];

    this.add = function(child) {
        this.children.push(child)
    }

    this.getChild = function(i) {
        return this.children[i]
    }

    this.remove = function(child) {
        this.children.remove(child)
    }

    this.getMaintananceCost = function() {
        return this.children.reduce((r, child) => r + child.getMaintananceCost(),0)
    }
}

var tunnelNetwork = new Network();
tunnelNetwork.add(new Tunnel(5));
tunnelNetwork.add(new Railroad(5));

var stationNetwork = new Network();
stationNetwork.add(new Station());
stationNetwork.add(new Railroad(10));

var bigNetwork = new Network();
bigNetwork.add(stationNetwork);
bigNetwork.add(tunnelNetwork);

console.log(bigNetwork.getMaintananceCost())

function Locomotive(max_speed) {
    this.max_speed = max_speed
    this.price = 15
    this.color = "red"

    this.clone = function() {
        return Object.assign(Object.create(this), JSON.parse(JSON.stringify(this)))
    }

    this.set_color = function (color) {
        this.color = color
    }

    this.set_price = function (price) {
        this.price = price
    }

    this.get_cost = function () {
        return this.price
    }
}

function Train(locomotive) {
    this.locomotive = locomotive

    this.wagons = []

    this.add_wagon = function (wagon) {
        this.wagons.push(wagon)
    }

    this.get_cost = function () {
        return this.locomotive.get_cost() + 5*this.wagons.length
    }

    this.clone = function () {
        var clone = Object.assign(Object.create(this), JSON.parse(JSON.stringify(this)))
        clone.locomotive = this.locomotive.clone()
        return clone
    }
}

oldLocomotive = new Locomotive(10)
newLocomotive = oldLocomotive.clone()

longTrain = new Train(oldLocomotive)
longTrain.add_wagon("Passenger")
shortTrain = longTrain.clone()
longTrain.add_wagon("Freight")

oldLocomotive.set_price(7)

console.log(newLocomotive.get_cost() - oldLocomotive.get_cost())
console.log(shortTrain.get_cost() - longTrain.get_cost())
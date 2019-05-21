function Train(id, color) {
    this.id = id;
    this.color = color;
}

function ModernTrain(id, color){
    Train.call(this, id, color)
}

ModernTrain.prototype = Object.create(Train.prototype);

function ClassicTrain(id, color){
    Train.call(this, id, color)
}
ClassicTrain.prototype = Object.create(Train);

function TrainCreator() { }

TrainCreator.prototype.createTrain = function() { throw Error('Not implemented') }

function ClassicTrainCreator(){}
ClassicTrainCreator.prototype = Object.create(TrainCreator.prototype);
ClassicTrainCreator.prototype.createTrain = () => new ClassicTrain(2, 'red');


function ModernTrainCreator(){}
ModernTrainCreator.prototype = Object.create(TrainCreator.prototype);
ModernTrainCreator.prototype.createTrain = () => new ModernTrain(1, 'white');

console.log(new ClassicTrainCreator().createTrain())
console.log(new ModernTrainCreator().createTrain())
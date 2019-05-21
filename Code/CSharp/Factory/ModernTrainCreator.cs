class ModernTrainCreator : TrainCreator {
    public override Train createTrain(){
        return new ModernTrain(1);
    }
}
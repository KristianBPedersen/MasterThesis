class ClassicTrainCreator : TrainCreator {
    public override Train createTrain(){
        return new ClassicTrain(2);
    }
}
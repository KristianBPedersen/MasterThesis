using System;

class FactoryRunner {
    public static void Main(){
        ClassicTrainCreator oldCreator = new ClassicTrainCreator();
        ModernTrainCreator newCreator = new ModernTrainCreator();

        oldCreator.createTrain().describe();
        newCreator.createTrain().describe();
    }
}
class FactoryRunner {
    public static void main(String[] args) {
        ClassicTrainCreator oldCreator = new ClassicTrainCreator();
        ModernTrainCreator newCreator = new ModernTrainCreator();

        oldCreator.createTrain().describe();
        newCreator.createTrain().describe();
    } 
}
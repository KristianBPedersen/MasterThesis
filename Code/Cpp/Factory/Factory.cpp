#include <string>
#include <iostream>

using namespace std;

class Train {
    private: 
    int id;
    string color;

    public:
    Train(int id, string color){
        this->id = id;
        this->color = color;
    }
     void describe() {
        printf("I am a train with id = %d and color = %s\n",id, color.c_str());
    }
};

class ClassicTrain : public Train {
    public: ClassicTrain(int id) : Train(id, "red") {}
};

class ModernTrain : public Train {
    public: ModernTrain(int id) : Train(id, "white") {}
};

class TrainCreator {
    public: virtual Train createTrain();
};

class ModernTrainCreator {
    public: Train createTrain(){
        return *(new ModernTrain(1));
    }
};

class ClassicTrainCreator {
    public: Train createTrain(){
        return *(new ClassicTrain(2));
    }
};

int main(){
    ModernTrainCreator* modernTrainCreator = new ModernTrainCreator();
    ClassicTrainCreator* classicTrainCreator = new ClassicTrainCreator();

    modernTrainCreator->createTrain().describe();
    classicTrainCreator->createTrain().describe();
}
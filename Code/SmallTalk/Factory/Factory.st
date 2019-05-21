Object subclass: Train [
    | id color |
    
    withId: idNumber withColor: paintColor [
	id := idNumber .
	color := paintColor .
    ]

    describe [
	|s|
	('I am a train with id = ', (id printString), ' and color ', color) printNl .
    ]
]

Train subclass: ClassicTrain [
    ClassicTrain class >> withId: id [
	| r |
	r := super new .
	^(r withId: id withColor: 'red')
    ]
]

Train subclass: ModernTrain [
    ModernTrain class >> withId: id [
	| r |
	r := super new .
	^(r withId: id withColor: 'white')
    ]
]

Object subclass: TrainCreator [
    createTrain [
	self announceStation .
    ]
]

TrainCreator subclass: ClassicTrainCreator [
    createTrain [
	^(ClassicTrain withId: 2 ) .
    ]
]

TrainCreator subclass: ModernTrainCreator [
    createTrain [
	^(ModernTrain withId: 1 ) .
    ]
]

(ModernTrainCreator new createTrain) describe .
(ClassicTrainCreator new createTrain) describe .

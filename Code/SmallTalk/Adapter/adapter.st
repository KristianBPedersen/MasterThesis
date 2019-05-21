Object subclass: Train [
    | wagonCount line passengerCount |

    Train class >> new: wagonCount [
	^(self new init: wagonCount) .
    ]

    init: wc [
	wagonCount := wc .
    ]

    setLine: lineNumber [
	line := 'T' , (lineNumber printString) .
    ]

    acceptPassengers [
	passengerCount := 50 * wagonCount .
    ]

    stop [
	passengerCount := passengerCount - (50 * wagonCount) .
    ]

    checkTrainTicket: code [
	'train' = code ifFalse: [
	    'Error' printNl .
	]
    ]

    drive [
	('Train '  , line , ' driving with ', (passengerCount printString) , ' passengers') printNl .
    ]
]

Object subclass: Bus [
    | capacity line passengerCount |

    Bus class >> new: capacity [
	^(self new init: capacity) .
    ]

    init: c [
	capacity := c .
    ]

    setLine: lineNumber [
	line := 'B' , (lineNumber printString) .
    ]

    acceptPassengers [
	passengerCount := capacity .
    ]

    releasePassengers [
	passengerCount := 0 .
    ]

    checkBusTicket: code [
	'bus' = code ifFalse: [
	    'Error' printNl .
	]
    ]

    drive [
	('Bus '  , line , ' driving with ', (passengerCount printString) , ' passengers') printNl .
    ]
]

Train subclass: SharingTrain [
    getLine [
	^line
    ]
]

Bus subclass: TrainReplacement [
    | train |
    
    init: capacity [
	super init: capacity .
	train := SharingTrain new: 1 .
    ]
	
    checkTrainTicket: code [
	train checkTrainTicket: code .
    ]

    setLine: lineNumber [
	train setLine: lineNumber .
	line := train getLine .
    ]

    stop [
        self releasePassengers .
    ]
]


Object subclass: TrainController [

    TrainController class >> runLine: lineNumber with: train [
	train setLine: lineNumber .
	train acceptPassengers .
	train checkTrainTicket: 'train' .
	train drive .
	train stop .
    ]
]


train := TrainReplacement new: 40 .
TrainController runLine: 2 with: train .

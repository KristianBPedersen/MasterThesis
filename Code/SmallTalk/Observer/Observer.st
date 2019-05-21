Object subclass: Train [
    | passengers currentStation |
    Train class >> new [
	| r |
	r := super new.
	r init.
	^r .
    ]
    
    init [
	passengers := OrderedCollection new .
    ]

    board: passenger [
	passenger enter: self.
	passengers add: passenger.
    ]

    discharge: passenger [
	passengers identityRemove: passenger.
    ]

    arrive: station [
	currentStation := station .
	self announceStation .
    ]

    announceStation [
	passengers size to: 1 by: -1 do: [:i | (passengers at: i) arrive: currentStation]
    ]

    isEmpty [
	^(passengers size = 0)
    ]

    currentStation [
	^currentStation .
    ]
]

Object subclass: Passenger [
    | train |

    enter: enteredTrain [
	train := enteredTrain
    ]

    arrive: station [self subclassResponsibility]
]

Passenger subclass: CountingPassenger [
    | counter |
    CountingPassenger class >> countingTo: count [
	| r |
	r := super new.
	r init: count .
	^r .
    ]

    init: count [
	counter := count.
    ]

    arrive: station [
	counter := counter - 1 .
	(counter = 0) ifTrue: [ train discharge: self ].
    ]
]

Passenger subclass: CheckingPassenger [
    | destination |
    CheckingPassenger class >> goingTo: station [
	| r |
	r := super new.
	r init: station .
	^r .
    ]

    init: station [
	destination := station.
    ]

    arrive: station [
	(destination = station) ifTrue: [ train discharge: self ].
    ]
]

train := Train new .

stations := {'station1' . 'station2' . 'station3' . 'station4' } .

1 to: 5 do: [:i | train board: (CountingPassenger countingTo: i) ] .

1 to: 4 do: [:i | train board: (CheckingPassenger goingTo: (stations at: i)) ] .

p := CheckingPassenger goingTo: 'home' .
p2 := CountingPassenger countingTo: 1 .

j := 1 .
[train isEmpty not] whileTrue: [
j := 1 + (j rem: (stations size)) .
train arrive: (stations at: j)
] .

train currentStation printNl


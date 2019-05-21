Object subclass: Product [
    get_cost [
	self subclassResponsibility
    ]
    clone [
	^(self deepCopy)
    ]
]

Product subclass: Locomotive [
    | price color max_speed |

    Locomotive class >> new: max_speed [
	^(self new init: max_speed) .
    ]

    init: speed [
	max_speed := speed .
	color := 'red' .
	price := 15 .
    ]

    set_price: p [
	price := p .
    ]

    get_cost [
	^(price) .
    ]

    get_color [
        ^color .
    ]

    set_color: c [
        color := c .
    ]
]

Product subclass: Train [
    | locomotive wagons |
    
    Train class >> new: locomotive [
	^(self new init: locomotive) .
    ]

    init: l [
	locomotive := l .
	wagons := OrderedCollection new .
    ]

    add_wagon: wagon [
	wagons add: wagon .
    ]

    get_cost [
	^(locomotive get_cost + (5 * wagons size))
    ]
]

oldLocomotive := Locomotive new: 15 .
newLocomotive := oldLocomotive clone .

longTrain := Train new: oldLocomotive .
longTrain add_wagon: 'Passenger' .
shortTrain := longTrain clone .
longTrain add_wagon: 'Freight' .

oldLocomotive set_price: 7 .

(newLocomotive get_cost - oldLocomotive get_cost) printNl .
(shortTrain get_cost - longTrain get_cost) printNl .

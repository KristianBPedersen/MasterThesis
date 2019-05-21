Object subclass: CargoTrain [
    | id |
    CargoTrain class >> new: idNr [
	^(self new init: idNr) .
    ]

    init: idNr [
	id := idNr .
    ]
    
    describe [
	^((id printString), ' is a cargo train. ' ) .
    ]
]

Object subclass: Decorator [
    | train |
    Decorator class >> new: t [
	^(self new init: t) .
    ]

    init: t [
	train := t .
    ]
    
    describe [
	^(train describe) .
    ]
]

Decorator subclass: ColorDecorator [
    | color |

    Decorator class >> new: t withColor: c [
	^(self new init: t withColor: c) .
    ]
    
    withColor: c [
	color := c .
    ]
    
    describe [
	^(super describe, 'It is painted ', color, '. ' ).
    ]
]

Decorator subclass: SilencerDecorator [
    describe [
	^(super describe, 'When it drives it makes no sound. ' ).
    ]
]

classicTrain := (ColorDecorator new: (CargoTrain new: 6)) withColor: 'red' .
stealthTrain := SilencerDecorator new: ((ColorDecorator new: (CargoTrain new: 7)) withColor: 'black') .

classicTrain describe printNl .
stealthTrain describe printNl .

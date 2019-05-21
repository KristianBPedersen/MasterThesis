Object subclass: Structure [
    getChild: i [
	^SystemExceptions.NotImplemented signal .
    ]
    add: child [
	^SystemExceptions.NotImplemented signal .
    ]
    remove: child [
	^SystemExceptions.NotImplemented signal .
    ]
    getMaintananceCost [
	self subclassResponsibility .
    ]
]

Structure subclass: Tunnel [
    | length |
    Tunnel class >> new: length [
	^(self new init: length) .
    ]

    init: l [
	length := l .
    ]

    getMaintananceCost [
	^(length*5) .
    ]
]

Structure subclass: Railroad [
    | length |
    Railroad class >> new: length [
	^(self new init: length) .
    ]

    init: l [
	length := l .
    ]

    getMaintananceCost [
	^(length*2) .
    ]
]

Structure subclass: Station [
    getMaintananceCost [
	^20 .
    ]
]

Object subclass: Network [
    | children |
    Network class >> new [
	^(super new init) .
    ]
    
    init [
	children := OrderedCollection new .
    ]

    getChild: i [
	^(children at: i) .
    ]

    add: child [
	children add: child .
    ]

    remove: child [
	children identityRemove: child .
    ]

    getMaintananceCost [
	| s |
	s := 0 .
	children do: [:child |
        s := s + child getMaintananceCost 
    ] .
	^s .
    ]
]
tunnelNetwork := Network new .
tunnelNetwork add: (Tunnel new: 5) .
tunnelNetwork add: (Railroad new: 5) .

stationNetwork := Network new .
stationNetwork add: (Railroad new: 10) .
stationNetwork add: Station new .

bigNetwork := Network new .
bigNetwork add: tunnelNetwork .
bigNetwork add: stationNetwork .

(bigNetwork getMaintananceCost) printNl .

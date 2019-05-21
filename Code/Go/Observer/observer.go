package main

import (
	"container/list"
	"fmt"
)

type Train struct {
	passengers *list.List
	station    string
}

type Passenger interface {
	enter(train *Train)
	arrive(station string)
}

type PassengerStruct struct {
	train *Train
	Passenger
}

type CountingPassenger struct {
	counter int
	PassengerStruct
}

type CheckingPassenger struct {
	destination string
	PassengerStruct
}

func NewTrain() *Train {
	return &Train{list.New(), ""}
}

func NewCountingPassenger(count int) *CountingPassenger {
	return &CountingPassenger{count, PassengerStruct{}}
}

func NewCheckingPassenger(destination string) *CheckingPassenger {
	return &CheckingPassenger{destination, PassengerStruct{}}
}

func (train *Train) board(passenger Passenger) {
	passenger.enter(train)
	train.passengers.PushFront(passenger)
}

func (train Train) discharge(passenger Passenger) {
	for e := train.passengers.Front(); e != nil; e = e.Next() {
		if e.Value.(Passenger) == passenger {
			train.passengers.Remove(e)
			return
		}
	}
}

func (train *Train) arrive(station string) {
	train.station = station
	train.announceStation()
}

func (train Train) announceStation() {
	next := train.passengers.Back()
	for e := train.passengers.Back(); e != nil; e = next {
		next = e.Prev()
		e.Value.(Passenger).arrive(train.station)
	}
}

func (passenger *PassengerStruct) enter(train *Train) {
	passenger.train = train
}

func (passenger *CountingPassenger) arrive(station string) {
	passenger.counter--
	if passenger.counter == 0 {
		passenger.train.discharge(passenger)
	}
}

func (passenger *CheckingPassenger) arrive(station string) {
	if passenger.destination == station {
		passenger.train.discharge(passenger)
	}
}

func main() {
	var train *Train = NewTrain()

	stations := [4]string{"Station1", "Station2", "Station3", "Station4"}

	for i := 1; i < 6; i++ {
		train.board(NewCountingPassenger(i))
	}

	for i := 0; i < 4; i++ {
		train.board(NewCheckingPassenger(stations[i]))
	}

	j := 0
	for train.passengers.Len() != 0 {
		j = (j + 1) % 4
		train.arrive(stations[j])
	}

	fmt.Printf("Final Station: %s", train.station)
}

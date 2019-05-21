package main

import (
	"fmt"
	"strconv"
)

type railroadRunner interface {
	setLine(int)
	acceptPassengers()
	stop()
	drive()
	checkTrainTicket(string)
}

type Train struct {
	wagonCount     int
	line           string
	passengerCount int
}

type Bus struct {
	capacity       int
	line           string
	passengerCount int
}

type TrainReplacement struct {
	*Bus
	train *Train
}

type RailroadController struct {
}

func NewTrain(wagonCount int) *Train {
	return &Train{wagonCount, "", 0}
}

func (train *Train) setLine(lineNumber int) {
	train.line = "T" + strconv.Itoa(lineNumber)
}

func (train *Train) acceptPassengers() {
	train.passengerCount = 50 * train.wagonCount
}

func (train *Train) stop() {
	train.passengerCount -= 50 * train.wagonCount
}

func (train *Train) drive() {
	fmt.Println("Train " + train.line + " drives with " + strconv.Itoa(train.passengerCount) + " passengers")
}

func (train *Train) checkTrainTicket(code string) {
	if code != "train" {
		fmt.Println("Error")
	}
}

func NewBus(capacity int) *Bus {
	return &Bus{capacity, "", 0}
}

func (bus *Bus) setLine(lineNumber int) {
	bus.line = "B" + strconv.Itoa(lineNumber)
}

func (bus *Bus) acceptPassengers() {
	bus.passengerCount = bus.capacity
}

func (bus *Bus) releasePassengers() {
	bus.passengerCount = 0
}

func (bus *Bus) drive() {
	fmt.Println("Bus " + bus.line + " drives with " + strconv.Itoa(bus.passengerCount) + " passengers")
}

func (bus *Bus) checkBusTicket(code string) {
	if code != "bus" {
		fmt.Println("Error")
	}
}

func NewTrainReplacement(capacity int) TrainReplacement {
	return TrainReplacement{NewBus(capacity), NewTrain(1)}
}

func (rTrain TrainReplacement) checkTrainTicket(code string) {
	rTrain.train.checkTrainTicket(code)
}

func (rTrain TrainReplacement) stop() {
	rTrain.releasePassengers()
}

func (rTrain TrainReplacement) setLine(lineNumber int) {
	rTrain.train.setLine(lineNumber)
	rTrain.line = rTrain.train.line
}

func (rc RailroadController) runLine(train railroadRunner, lineNumber int) {
	train.setLine(lineNumber)
	train.acceptPassengers()
	train.checkTrainTicket("train")
	train.drive()
	train.stop()
}

func main() {
	train := NewTrainReplacement(40)
	RailroadController{}.runLine(train, 2)
}

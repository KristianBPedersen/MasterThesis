package main

import (
	"fmt"
)

type Train interface {
	getId() int
	getColor() string
}

type TrainStruct struct {
	id    int
	color string
}

func (t TrainStruct) getId() int {
	return t.id
}

func (t TrainStruct) getColor() string {
	return t.color
}

type ModernTrain struct {
	TrainStruct
}

type ClassicTrain struct {
	TrainStruct
}

type TrainCreator interface {
	createTrain() Train
}

type ModernTrainCreator struct {
	TrainCreator
}

type ClassicTrainCreator struct {
	TrainCreator
}

func (c ModernTrainCreator) createTrain() Train {
	return ModernTrain{TrainStruct{1, "white"}}
}

func (c ClassicTrainCreator) createTrain() Train {
	return ClassicTrain{TrainStruct{2, "red"}}
}

func describe(t Train) {
	fmt.Println("I am a train with id =", t.getId(), "and color = "+t.getColor())
}

func main() {
	describe(ModernTrainCreator{}.createTrain())
	describe(ClassicTrainCreator{}.createTrain())
}

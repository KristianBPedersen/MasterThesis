package main

import (
	"fmt"
	"strconv"
)

type Train interface {
	describe() string
}

type CargoTrain struct {
	id int
}

type Modification struct {
	train Train
}

type ColorModification struct {
	*Modification
	super *Modification
	color string
}

type SilencerModification struct {
	*Modification
	super *Modification
}

func (train *CargoTrain) describe() string {
	return strconv.Itoa(train.id) + " is a cargo train. "
}

func (modification *Modification) describe() string {
	return modification.train.describe()
}

func (modification *ColorModification) describe() string {
	return modification.super.describe() + "It is painted " + modification.color + ". "
}

func (modification *SilencerModification) describe() string {
	return modification.super.describe() + "When it drives it makes no sound. "
}

func NewModification(train Train) *Modification {
	return &Modification{train}
}

func NewColorModification(train Train, color string) *ColorModification {
	mod := NewModification(train)
	return &ColorModification{mod, mod, color}
}

func NewSilencerModification(train Train) *SilencerModification {
	mod := NewModification(train)
	return &SilencerModification{mod, mod}
}

func main() {
	classicTrain := NewColorModification(&CargoTrain{6}, "red")
	stealthTrain := NewSilencerModification(NewColorModification(&CargoTrain{7}, "black"))
	fmt.Println(classicTrain.describe())
	fmt.Println(stealthTrain.describe())
}

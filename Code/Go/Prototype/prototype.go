package main

import "fmt"

type Product interface {
	get_cost() int
	clone() Product
}

type Locomotive struct {
	max_speed int
	color     string
	price     int
}

type Train struct {
	locomotive *Locomotive
	wagons     []string
}

func (l *Locomotive) get_cost() int {
	return l.price
}

func (l *Locomotive) set_price(price int) {
	l.price = price
}

func (l *Locomotive) set_color(color string) {
	l.color = color
}

func (l *Locomotive) clone() Product {
	newLocomotive := *l
	return &newLocomotive
}

func NewLocomotive(max_price int) *Locomotive {
	return &Locomotive{max_price, "red", 15}
}

func NewTrain(locomotive *Locomotive) *Train {
	return &Train{locomotive, []string{}}
}

func (train *Train) add_wagon(wagon string) {
	train.wagons = append(train.wagons, wagon)
}

func (train *Train) clone() Product {
	newTrain := *train
	copy(newTrain.wagons, train.wagons)
	newTrain.locomotive = train.locomotive.clone().(*Locomotive)
	return &newTrain
}

func (train *Train) get_cost() int {
	return train.locomotive.get_cost() + 5*len(train.wagons)
}

func main() {
	oldLocomotive := NewLocomotive(15)

	newLocomotive := oldLocomotive.clone()

	longTrain := NewTrain(oldLocomotive)
	longTrain.add_wagon("Passenger")
	shortTrain := longTrain.clone()
	longTrain.add_wagon("Freight")

	oldLocomotive.set_price(7)

	fmt.Println(newLocomotive.get_cost() - oldLocomotive.get_cost())
	fmt.Println(shortTrain.get_cost() - longTrain.get_cost())
}

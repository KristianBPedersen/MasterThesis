package main

import (
	"fmt"
)

type NetworkElement interface {
	getChild(int) NetworkElement
	add(NetworkElement)
	remove(NetworkElement)
	getMaintananceCost() int
}

type Structure struct {
}

type Railroad struct {
	*Structure
	length int
}

type Tunnel struct {
	*Structure
	length int
}

type Station struct {
	*Structure
}

type Network struct {
	children []NetworkElement
}

func (s *Structure) getChild(i int) NetworkElement {
	panic("NotImplemented")
}

func (s *Structure) add(child NetworkElement) {
	panic("NotImplemented")
}

func (s *Structure) remove(child NetworkElement) {
	panic("NotImplemented")
}

func (railroad *Railroad) getMaintananceCost() int {
	return railroad.length * 2
}

func (railroad *Tunnel) getMaintananceCost() int {
	return railroad.length * 5
}

func (station *Station) getMaintananceCost() int {
	return 20
}

func NewNetwork() *Network {
	return &Network{[]NetworkElement{}}
}

func (network *Network) add(child NetworkElement) {
	network.children = append(network.children, child)
}

func (network *Network) getChild(i int) NetworkElement {
	return network.children[i]
}
func (network *Network) remove(child NetworkElement) {
	for i := 0; i < len(network.children); i++ {
		if network.children[i] == child {
			network.children = append(network.children[:i], network.children[i+1:]...)
		}
	}
}
func (network *Network) getMaintananceCost() int {
	s := 0
	for _, child := range network.children {
		s += child.getMaintananceCost()
	}
	return s
}

func main() {
	tunnelNetwork := NewNetwork()
	tunnelNetwork.add(&Tunnel{&Structure{}, 5})
	tunnelNetwork.add(&Railroad{&Structure{}, 5})

	stationNetwork := NewNetwork()
	stationNetwork.add(&Station{&Structure{}})
	stationNetwork.add(&Railroad{&Structure{}, 10})

	bigNetwork := NewNetwork()
	bigNetwork.add(stationNetwork)
	bigNetwork.add(tunnelNetwork)

	fmt.Println(bigNetwork.getMaintananceCost())
}

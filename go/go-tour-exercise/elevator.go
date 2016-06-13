package main

import (
	"container/list"
	"fmt"
)

type Elevator struct {
	Ip string
}

type Elevators struct {
	List *list.List
}

func (e *Elevators) AddToList(elevator *Elevator) {
	e.List.PushBack(elevator)
}

func (e *Elevators) IPIsInList(ip string) *Elevator {
	for c := e.List.Front(); c != nil; c = c.Next() {
		if c.Value.(*Elevator).Ip == ip {
			return c.Value.(*Elevator)
		}
	}
	return nil
}

func main() {
	elevators := Elevators{list.New()}
	ip := "192.168.0.10"
	elevator := Elevator{Ip: ip}
	elevators.AddToList(&elevator)
	eip := elevators.IPIsInList(ip)
	if eip != nil {
		fmt.Println(*eip)
	}
}

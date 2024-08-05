package main

import (
	"math/rand"
)

type Machine struct {
	Data []uint
	Size int

	Head int
}

func NewMachine(size int) *Machine {
	m := Machine{Size: size}
	m.Data = []uint{}
	for range m.Size {
		m.Data = append(m.Data, 0)
	}
	return &m
}

func (m *Machine) Randomize() {
	m.Data = []uint{}
	for range m.Size {
		m.Data = append(m.Data, uint(rand.Intn(2)))
	}
}

func (m *Machine) Execute(inst *Instruction) int {
	if m.Data[m.Head] == inst.Expected {
		m.Data[m.Head] = inst.WriteYes
		m.Head += int(inst.DirYes)
		return inst.NextInstYes
	} else {
		m.Data[m.Head] = inst.WriteNo
		m.Head += int(inst.DirNo)
		return inst.NextInstNo
	}
}

type Direction int

const (
	LEFT  Direction = -1
	STAY  Direction = 0
	RIGHT Direction = 1
)

type Instruction struct {
	Expected uint

	WriteYes    uint
	DirYes      Direction
	NextInstYes int

	WriteNo    uint
	DirNo      Direction
	NextInstNo int

	Null bool
}
